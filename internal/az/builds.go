package az

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetBuildList(azureOrg, azurePrj, azurePAT string) *BuildList {
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

	return &bl
}
