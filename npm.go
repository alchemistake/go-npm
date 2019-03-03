package npm

import (
  "net/http"
  "fmt"
  "encoding/json"
  "io/ioutil"
  "bytes"
)

const baseURL string = "https://registry.npmjs.org/-"

//Client struct for npm
type Client struct {
	Username string
	Password string
}

//Membership npm org membership
type Membership struct{
  User string `json:"user"`
  Role string `json:"role"`
}

//NewBasicAuthClient using username:pass
func NewBasicAuthClient(username, password string) *Client {
	return &Client{
		Username: username,
		Password: password,
	}
}

func (s *Client) doRequest(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(s.Username, s.Password)
	client := &http.Client{}
  req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

//GetUsers Get all users from org
func (s *Client) GetUsers(org, user string) (interface{}, error) {
	url := fmt.Sprintf(baseURL+"/org/%s/user", org)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data  map[string]interface{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}


	return data, nil
}

//AddUser to npm org
func (s *Client) AddUser(org, user, role string) ( error) {
	url := fmt.Sprintf(baseURL+"/org/%s/user", org)
  fmt.Println(url)
  member := Membership{User: user, Role: role }
  b, err := json.Marshal(member)
    if err != nil {
        fmt.Println(err)
        return nil
    }
    fmt.Println(string(b))
  body := bytes.NewBuffer(b)
	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return  err
	}
  fmt.Println(req)
	_, err = s.doRequest(req)
	if err != nil {
		return  err
	}
	return  nil
}

//DeleteUser from npm org
func (s *Client) DeleteUser(org, user string) (error) {
	url := fmt.Sprintf(baseURL+"/org/%s/user", org)
  fmt.Println(url)
  member := Membership{User: user}
  b, err := json.Marshal(member)
    if err != nil {
        return  err
    }
    fmt.Println(string(b))
  body := bytes.NewBuffer(b)
	req, err := http.NewRequest(http.MethodDelete, url, body)
	if err != nil {
		return  err
	}
	_, err = s.doRequest(req)
	if err != nil {
		return  err
	}
	return  nil
}
