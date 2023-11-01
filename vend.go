package vend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiURI = "vendhq.com/api"

type Client struct {
	DomainPrefix string
	APIVersion   string
	Token        string
	Product      ProductService
	// Shipment  ShipmentService
	// Store     StoreService
	// Tag       TagService
}

func NewClient(domainPrefix, apiVersion, token string) *Client {

	c := &Client{
		DomainPrefix: domainPrefix,
		APIVersion:   apiVersion,
		Token:        token,
	}

	c.Product = &ProductServiceOp{client: c}
	// c.Shipment = &ShipmentServiceOp{client: c}
	// c.Store = &StoreServiceOp{client: c}
	// c.Tag = &TagServiceOp{client: c}

	return c

}

func (c *Client) Request(method string, path string, body interface{}, v interface{}) error {

	var bodyReader io.Reader
	if body != nil {
		requestJson, errMarshal := json.Marshal(body)
		if errMarshal != nil {
			return errMarshal
		}

		bodyReader = bytes.NewBuffer(requestJson)
	}

	//HTTP
	url := "https://" + c.DomainPrefix + "." + apiURI + "/" + c.APIVersion + "/" + path
	fmt.Println(url)
	httpReq, errNewRequest := http.NewRequest(method, url, bodyReader)
	if errNewRequest != nil {
		return errNewRequest
	}

	// Basic Auth
	httpReq.Header.Set("Authorization", "Bearer "+c.Token)

	// Content Type
	httpReq.Header.Set("Content-Type", "application/json")

	//Client
	client := &http.Client{}
	resp, errDo := client.Do(httpReq)

	if resp != nil {
		defer resp.Body.Close()
	}
	if errDo != nil {
		return errDo
	}

	if v != nil {
		decoder := json.NewDecoder(resp.Body)
		errDecode := decoder.Decode(&v)
		if errDecode != nil {
			return errDecode
		}
	}
	return nil
}
