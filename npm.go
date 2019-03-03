package npm

import (
  "net/http"
  "fmt"
  "encoding/json"
  "io/ioutil"
)

const baseURL string = "https://registry.npmjs.org/-"

//Client struct for npm
type Client struct {
	Username string
	Password string
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
  // fmt.Println(string(body))
	return body, nil
}

//GetUsers Get all users from org
func (s *Client) GetUsers(org string) (interface{}, error) {
	url := fmt.Sprintf(baseURL+"/org/%s/user", org)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data interface{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
