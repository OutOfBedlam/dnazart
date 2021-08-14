package az

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetArtifactList(azureOrg, azurePrj string, buildId int, azurePAT string) (*ArtifactList, error) {
	// build request
	reqAuth := fmt.Sprintf("username:%s", azurePAT)
	reqBearer := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(reqAuth)))
	url := fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/build/builds/%d/artifacts?api-version=4.1", azureOrg, azurePrj, buildId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", reqBearer)

	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	al := ArtifactList{}
	err = json.Unmarshal(body, &al)
	if err != nil {
		return nil, err
	}

	return &al, nil
}
