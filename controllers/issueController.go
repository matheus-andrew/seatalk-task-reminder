package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetIssue(jql string) []byte{
	url := "https://jira.shopee.io/rest/api/2/search?jql="
	bearer := "Bearer " + os.Getenv("BEARER_TOKEN")
	fmt.Println(bearer)
	req, err := http.NewRequest("GET", url+jql, nil)
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	return body
}