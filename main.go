package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"git.garena.com/matheus.setiawan/go-jira/controllers"
	"git.garena.com/matheus.setiawan/go-jira/models"
	"git.garena.com/matheus.setiawan/go-jira/services"
	"github.com/joho/godotenv"
)

func main() {
	// request get jira
	err := godotenv.Load()
    if err != nil {
        fmt.Printf("err load env: %v\n", err)
    }

	jql := os.Getenv("JQL")
	fmt.Println(jql)
	body := controllers.GetIssue(jql)

	// convert json to struct Issue
	var data models.Issue
	json.Unmarshal(body, &data)

	// convert json to array struct Key
	var keyModel []models.Key
	for i := range data.Issues {
		var fixVersions []string
		for j := range data.Issues[i].Fields.FixVersions {
			fixVersions = append(fixVersions, data.Issues[i].Fields.FixVersions[j].Name)
		}

		jira := models.Key{
			Key:        data.Issues[i].Key,
			Qa:         data.Issues[i].Fields.Customfield10308.EmailAddress,
			Dev:        data.Issues[i].Fields.Customfield10307.EmailAddress,
			Status:     data.Issues[i].Fields.Status.Name,
			Summary:    data.Issues[i].Fields.Summary,
			FixVersion: fixVersions,
			Labels:     data.Issues[i].Fields.Labels,
		}

		keyModel = append(keyModel, jira)
	}

	// create one line string
	var text string
	for i := range keyModel {
		text += "__https://jira.shopee.io/browse/" + keyModel[i].Key + "__ | *" + keyModel[i].Status + "*" +
			"\n\nSummary: " + keyModel[i].Summary +
			"\n\nQA: " + keyModel[i].Qa +
			"\n\nDev: " + keyModel[i].Dev +
			"\n\nFixVersions: " + strings.Join(keyModel[i].FixVersion, ", ") +
			"\n\nLabels: " + strings.Join(keyModel[i].Labels, ", ") +
			"\n\n .\n\n"
	}
	text += "__Total: " + strconv.Itoa(len(keyModel)) + "__"

	services.SendMessage(text)
}
