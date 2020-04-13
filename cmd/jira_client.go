package cmd

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
)

func NewJiraClient(cnx ConnectionSettings) *jira.Client {
	tp := jira.BasicAuthTransport{Username: cnx.UserName, Password: cnx.Token}

	jiraClient, err := jira.NewClient(tp.Client(), cnx.BaseUrl)

	if err != nil {
		panic(err)
	}

	return jiraClient
}

func AddComment(issueKey, body string) (*jira.Comment, error) {
	cnx := LoadConnection()
	jiraClient := NewJiraClient(cnx)

	comment, _, err := jiraClient.Issue.AddComment(
		issueKey,
		&jira.Comment{Body: body},
	)

	return comment, err
}

func GetIssueList(key string) ([]jira.Issue, error) {
	cnx := LoadConnection()
	jiraClient := NewJiraClient(cnx)
	jql := fmt.Sprintf("project=%s", key)

	issues, _, err := jiraClient.Issue.Search(jql, &jira.SearchOptions{})
	return issues, err
}

func GetProjectList() (*jira.ProjectList, error) {
	cnx := LoadConnection()
	jiraClient := NewJiraClient(cnx)

	projects, _, err := jiraClient.Project.GetList()

	return projects, err
}

func GetProject(key string) (*jira.Project, error) {
	cnx := LoadConnection()
	jiraClient := NewJiraClient(cnx)

	project, _, err := jiraClient.Project.Get(key)
	return project, err
}

func GetIssue(key string) (*jira.Issue, error) {
	cnx := LoadConnection()
	jiraClient := NewJiraClient(cnx)

	issue, _, err := jiraClient.Issue.Get(key, nil)
	return issue, err
}
