package main

import (
	"fmt"
	"github.com/selvamshan/bookstore_oauth-go/oauth"
)


func main() {
	//var at oauth.AccessToken
	at, err := oauth.GetAccessToken("b5da53189dc8a03fa0ea07e87e421c78")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*at)
}