package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// License contains our licenses
type License []struct {
	ActivationEmail        string `json:"ActivationEmail"`
	ActiveInstanceCount    int64  `json:"ActiveInstanceCount"`
	AirgapDownloadEnabled  bool   `json:"AirgapDownloadEnabled"`
	AirgapDownloadPassword string `json:"AirgapDownloadPassword"`
	Anonymous              bool   `json:"Anonymous"`
	AppID                  string `json:"AppId"`
	AppStatus              string `json:"AppStatus"`
	Archived               bool   `json:"Archived"`
	Assignee               string `json:"Assignee"`
	AssistedSetupEnabled   bool   `json:"AssistedSetupEnabled"`
	Billing                struct {
		Begin     string `json:"begin"`
		End       string `json:"end"`
		Frequency string `json:"frequency"`
		Revenue   string `json:"revenue"`
	} `json:"Billing"`
	BillingEvents          []interface{} `json:"BillingEvents"`
	ChannelID              string        `json:"ChannelId"`
	ChannelName            string        `json:"ChannelName"`
	Clouds                 string        `json:"Clouds"`
	ExpirationPolicy       string        `json:"ExpirationPolicy"`
	ExpireDateString       string        `json:"ExpireDate"`
	ExpireDate             time.Time
	ExternalSupportBundle  bool        `json:"ExternalSupportBundle"`
	FieldValues            interface{} `json:"FieldValues"`
	GrantDate              string      `json:"GrantDate"`
	ID                     string      `json:"Id"`
	InactiveInstanceCount  int64       `json:"InactiveInstanceCount"`
	IsAppVersionLocked     bool        `json:"IsAppVersionLocked"`
	IsInstanceTracked      bool        `json:"IsInstanceTracked"`
	LastSync               string      `json:"LastSync"`
	LicenseType            string      `json:"LicenseType"`
	LicenseVersions        interface{} `json:"LicenseVersions"`
	LockedAppVersion       int64       `json:"LockedAppVersion"`
	RequireActivation      bool        `json:"RequireActivation"`
	RevokationDate         interface{} `json:"RevokationDate"`
	UntrackedInstanceCount int64       `json:"UntrackedInstanceCount"`
	UpdatePolicy           string      `json:"UpdatePolicy"`
	UseConsoleSupportSpec  bool        `json:"UseConsoleSupportSpec"`
}

func addLicenseExpiry(d License) License {
	for i := len(d) - 1; i >= 0; i-- {
		if d[i].ExpireDateString != "" {
			t, err := time.Parse(time.RFC3339, d[i].ExpireDateString)
			if err != nil {
				log.Fatal(err)
			}
			d[i].ExpireDate = t
		}
	}
	return d
}

func filter(d License) License {
	var l = License{}
	for i := len(d) - 1; i >= 0; i-- {
		if d[i].Archived == false {
			if time.Now().After(d[i].ExpireDate) {
				if (d[i].LicenseType == "prod") || (d[i].LicenseType == "paid") || (d[i].LicenseType == "trial") {
					l = append(l, d[i])
				}
			}
		}
	}
	return l
}

func getData() []byte {
	token := "REPLICATED_AUTHORIZATION_TOKEN"
	authKey, exists := os.LookupEnv(token)
	if !exists {
		log.Fatalf("%s is missing\n", token)
	}
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.replicated.com/vendor/v1/licenses", nil)
	req.Header.Set("Authorization", authKey)
	res, err := client.Do(req)
	if err != nil || res.StatusCode != 200 {
		log.Fatalf("Error:\nStatus Code: %v\n%v\n", res.Status, err)
	}
	defer res.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	return bodyBytes
}

func main() {
	licenses := getData()
	collection := License{}
	err := json.Unmarshal(licenses, &collection)
	if err != nil {
		log.Fatalf("Unable to unmarshal: %v\n", err)
	}
	collection = addLicenseExpiry(collection)
	collection = filter(collection)
	fmt.Printf("Final Count: %d\n", len(collection))

	//for i := len(collection) - 1; i >= 0; i-- {
	//	URL := fmt.Sprintf("https://vendor.replicated.com/apps/waffleio/customers/%s/manage", collection[i].ID)
	//	fmt.Printf("%s - %s\n", collection[i].Assignee, URL)
	//}
}
