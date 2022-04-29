package services

import (
	"fmt"
	"gcp-quota-control/pkg/webtool"
	"os"
)

func Get(){

	project := os.Getenv("GCP_PROJECT")
	url := fmt.Sprintf("https://monitoring.googleapis.com/v3/projects/%s/services", project)
	body, err := webtool.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", body)
}
