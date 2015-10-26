package cf_deployment_tracker

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"os"
	"time"
)

var deploymentTrackerURL = "https://deployment-tracker.mybluemix.net/api/v1/track"

type Repository struct {
	Url string
}

type Package struct {
	Name       string
	Version    string
	Repository Repository
}

type Event struct {
	DateSent           string   `json:"date_sent"`
	CodeVersion        string   `json:"code_version"`
	RepositoryURL      string   `json:"repository_url"`
	ApplicationName    string   `json:"application_name"`
	SpaceID            string   `json:"space_id"`
	ApplicationVersion string   `json:"application_version"`
	ApplicatonURIs     []string `json:"application_uris"`
}

func Track() (errs []error) {
	content, err := ioutil.ReadFile("package.json")
	//exit early if we cant read the file
	if err != nil {
		return
	}

	var info Package
	err = json.Unmarshal(content, &info)
	//exit early if we can't parse the file
	if err != nil {
		return
	}

	vcapApplication := os.Getenv("VCAP_APPLICATION")
	var event Event
	err = json.Unmarshal([]byte(vcapApplication), &event)
	//exit early if we are not running in CF
	if err != nil {
		return
	}

	if info.Repository.Url != "" {
		event.RepositoryURL = info.Repository.Url
	}
	if info.Name != "" {
		event.ApplicationName = info.Name
	}
	if info.Version != "" {
		event.CodeVersion = info.Version
	}

	now := time.Now()
	event.DateSent = now.UTC().Format("2006-01-02T15:04:05.999Z")

	request := gorequest.New()
	resp, _, errs := request.Post(deploymentTrackerURL).
		Send(event).
		End()

	if errs != nil {
		return errs
	}

	fmt.Println(resp.Status)

	return nil

}
