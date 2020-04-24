package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	logger "github.com/cihub/seelog"
)

var gClient *http.Client

func init() {
	gClient = &http.Client{
		Timeout: 20 * time.Second,
	}
}

func Post(url string, request interface{}, response interface{}, meta ...map[string]string) error {
	return doRequest("POST", url, request, response, meta...)
}

func Get(url string, response interface{}, meta ...map[string]string) error {
	return doRequest("GET", url, nil, response, meta...)
}

func doRequest(method string, url string, request interface{}, response interface{}, meta ...map[string]string) error {
	var requestBody []byte
	if request != nil {
		d, err := json.Marshal(request)
		if err != nil {
			logger.Error(err)
			return err
		}
		requestBody = d
	}

	logger.Infof("Post data: %s,\n to: %s\n", requestBody, url)

	//resp, err := gClient.Post(url, "application/json", bytes.NewBuffer(requestBody))

	httpReq, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	httpReq.Header.Set("Content-type", "application/json")

	if len(meta) > 0 && meta[0] != nil {
		for k, v := range meta[0] {
			httpReq.Header.Set(k, v)
		}
	}
	if len(meta) > 1 && meta[1] != nil {
		q := httpReq.URL.Query()
		for k, v := range meta[1] {
			q.Add(k, v)
		}
		httpReq.URL.RawQuery = q.Encode()
	}

	resp, err := gClient.Do(httpReq)
	if err != nil {
		logger.Error(err)
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf(fmt.Sprintf("HttpStatus:%d\n", resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Infof("post readall error: %+v\n", err)
		return err
	}

	logger.Infof("Post url: %s, response:%s\n", url, string(body))

	if response != nil {
		return json.Unmarshal(body, response)
	} else {
		return nil
	}
}
