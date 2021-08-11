package az

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var azurePAT = ""
var azureOrg = ""
var azurePrj = ""

type BuildList struct {
	Count int          `json:"count"`
	Value []BuildValue `json:"value"`
}

type BuildValue struct {
	Id                  int        `json:"id"`
	BuildNumber         string     `json:"buildNumber"`
	BuildNumberRevision int        `json:"buildNumberRevision"`
	Status              string     `json:"status"`
	Result              string     `json:"result"`
	QueueTime           string     `json:"queueTime"`
	StartTime           string     `json:"startTime"`
	FinishTime          string     `json:"finishTime"`
	Url                 string     `json:"url"`
	Uri                 string     `json:"uri"`
	SourceBranch        string     `json:"sourceBranch"`
	SourceVersion       string     `json:"sourceVersion"`
	Definition          Definition `json:"definition"`
	// Queue struct `json:"queue"`

	Project Project `json:"project"`
	Links   Links   `json:"_links"`
	// Properties        map[string]string  `json:"properties"`
	// Tags              map[string]string  `json:"tags"`
	ValidationResults []ValidationResult `json:"validationResults"`
}

type Definition struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Url         string  `json:"url"`
	Uri         string  `json:"uri"`
	Path        string  `json:"path"`
	Type        string  `json:"type"`
	QueueStatus string  `json:"queueStatus"`
	Revision    int     `json:"revision"`
	Project     Project `json:"project"`
}
type Project struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Url            string `json:"url"`
	State          string `json:"state"`
	Revision       int    `json:"revision"`
	Visibility     string `json:"visibility"`
	LastUpdateTime string `json:"lastUpdateTime"`
}

type Links struct {
	Self                    Link `json:"self"`
	Web                     Link `json:"web"`
	SourceVersionDisplayUri Link `json:"sourceVersionDisplayUri"`
	Timeline                Link `json:"timeline"`
	Badge                   Link `json:"badge"`
}

type Link struct {
	Href string `json:"href"`
}

type ValidationResult struct {
	Result  string `json:"result"`
	Message string `json:"message"`
}

func GetBuildsDefinitions() {
	// build request
	reqAuth := fmt.Sprintf("username:%s", azurePAT)
	reqBearer := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(reqAuth)))
	url := fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/build/builds?api-version=4.1", azureOrg, azurePrj)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", reqBearer)

	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}
	bl := BuildList{}
	err = json.Unmarshal(body, &bl)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Builds count %d\n", bl.Count)
	for i, v := range bl.Value {
		fmt.Printf("  [%d] def:%v %s id:%d status:%s result:%s\n", i, v.Definition.Id, v.Definition.Name, v.Id, v.Status, v.Result)
	}
}
