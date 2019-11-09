package getEx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

type user struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Company  company `json:"company"`
}

func RetrieveUsers() {
	resp, err := http.Get("http://jsonplaceholder.typicode.com/users")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var users []user
	if err = json.Unmarshal(body, &users); err != nil {
		fmt.Println(err)
		return
	}

	for _, u := range users {
		fmt.Println(u.Name, u.Username, u.Email)
		fmt.Println("\t", u.Company.Name)
	}
}
