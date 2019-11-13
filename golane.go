package golane

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/azurlane-api/golane/structs"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	// Version the package version
	Version = "1.0.2"
	baseURL = "https://azurlane-api.appspot.com/v1"
)

var (
	userAgent = fmt.Sprintf("golane/%s (https://github.com/azurlane-api/golane)", Version)
	// Order enum of order
	Order = &orderBy{
		RARITY:      "rarity",
		TYPE:        "type",
		AFFILIATION: "affiliation",
	}
)

type order string

type orderBy struct {
	RARITY      order
	TYPE        order
	AFFILIATION order
}

// AzurLane set default settings
type AzurLane struct {
	UserAgent string
}

// Init initialize struct, first param can be user-agent to set a custom one, leave empty to use the package ua
func (al *AzurLane) Init(params ...string) {
	if len(params) > 0 {
		al.UserAgent = params[0]
	} else {
		al.UserAgent = userAgent
	}
}

func get(apiURL string, ua string) ([]byte, error) {
	request, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", ua)
	request.Header.Set("Accept", "application/json")

	client := http.Client{
		Timeout: time.Second * 10,
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	buffer := bytes.NewBuffer(nil)
	_, err = io.Copy(buffer, response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Expected status %d; Got %d\nResponse: %#v", 200, response.StatusCode, buffer.String())
	}

	return buffer.Bytes(), nil
}

// GetShipByName get ship info by name
func (al AzurLane) GetShipByName(name string) (*structs.Ship, error) {
	url := fmt.Sprintf("%s/ship?name=%s", baseURL, url.PathEscape(name))
	bytes, err := get(url, al.UserAgent)
	if err != nil {
		return nil, err
	}

	var response = new(structs.ShipResponse)
	json.Unmarshal(bytes, &response)
	return &response.Ship, nil
}

// GetShipByID get ship info by id
func (al AzurLane) GetShipByID(id string) (*structs.Ship, error) {
	url := fmt.Sprintf("%s/ship?id=%s", baseURL, url.PathEscape(id))
	bytes, err := get(url, al.UserAgent)
	if err != nil {
		return nil, err
	}

	var response = new(structs.ShipResponse)
	json.Unmarshal(bytes, &response)
	return &response.Ship, nil
}

// GetShips returns a list of ships from rarity, type or affiliation
// order should be set using the "enum" Order, for example Order.TYPE
// value depends on what order is set too, for example if `Order.RARITY` is used value can be `Super Rare`
func (al AzurLane) GetShips(order order, value string) ([]structs.SmallShip, error) {
	url := fmt.Sprintf("%s/ships?orderBy=%s&%s=%s", baseURL, order, order, url.PathEscape(value))
	bytes, err := get(url, al.UserAgent)
	if err != nil {
		return nil, err
	}

	var response = new(structs.ShipsResponse)
	json.Unmarshal(bytes, &response)
	return response.Ships, nil
}

// GetBuildInfo retuns info about a certain construction time
func (al AzurLane) GetBuildInfo(time string) (*structs.Construction, error) {
	url := fmt.Sprintf("%s/build?time=%s", baseURL, url.PathEscape(time))
	bytes, err := get(url, al.UserAgent)
	if err != nil {
		return nil, err
	}

	var response = new(structs.ConstructionResponse)
	json.Unmarshal(bytes, &response)
	return &response.Construction, nil
}
