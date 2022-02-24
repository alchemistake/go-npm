package npm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL string = "https://registry.npmjs.org/-"

//Client struct for npm
type Client struct {
	Token 	string
}

//Membership npm org membership
type Membership struct {
	User string `json:"user"`
	Role string `json:"role"`
}

func NewTokenClient(token string) *Client {
	return &Client{
		Token: token
	}
}

func (s *Client) doRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	var bearer = "Bearer " + s.Token

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", bearer)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

//GetUsers Get all users from org
func (s *Client) GetUsers(org string) (map[string]string, error) {
	url := fmt.Sprintf(baseURL+"/org/%s/user", org)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data map[string]string
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

//AddUser to npm org
func (s *Client) AddUser(org, user, role string) error {
	url := fmt.Sprintf(baseURL+"/org/%s/user", org)
	fmt.Println(url)
	member := Membership{User: user, Role: role}
	b, err := json.Marshal(member)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(b)
	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return err
	}
	_, err = s.doRequest(req)
	if err != nil {
		return err
	}
	return nil
}

//DeleteUser from npm org
func (s *Client) DeleteUser(org, user string) error {
	url := fmt.Sprintf(baseURL+"/org/%s/user", org)
	fmt.Println(url)
	member := Membership{User: user}
	b, err := json.Marshal(member)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	body := bytes.NewBuffer(b)
	req, err := http.NewRequest(http.MethodDelete, url, body)
	if err != nil {
		return err
	}
	_, err = s.doRequest(req)
	if err != nil {
		return err
	}
	return nil
}
