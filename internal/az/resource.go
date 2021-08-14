package az

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
)

func (rc *Resource) Download(dst string, azurePAT string) error {
	// build request
	reqAuth := fmt.Sprintf("username:%s", azurePAT)
	reqBearer := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(reqAuth)))
	req, err := http.NewRequest("GET", rc.DownloadUrl, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", reqBearer)

	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, rsp.Body)
	return err
}
