package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func prep() License {
	data, err := ioutil.ReadFile("test.json")
	if err != nil {
		log.Fatal(err)
	}

	collection := License{}
	err = json.Unmarshal(data, &collection)
	if err != nil {
		fmt.Printf("Unable to unmarshal: %v\n", err)
	}

	return collection
}

func TestAddLicenseExpiry(t *testing.T) {
	testData := prep()
	testData = addLicenseExpiry(testData)
	sD := testData[0].ExpireDateString
	tD := testData[0].ExpireDate.Format(time.RFC3339)
	assert.Equal(t, tD, sD, "should be equal")
}

func TestFilter(t *testing.T) {
	collection := prep()
	collection = addLicenseExpiry(collection)
	collection = filter(collection)
	filterCount := len(collection)
	desiredCount := 1

	assert.Equal(t, desiredCount, filterCount, "should be equal")
}
