package vpa

import (
	"net/http"
	"io/ioutil"
	"bytes"
	"k8s.io/kubernetes/staging/src/k8s.io/apimachinery/pkg/util/json"
)

type RecommenderClient struct {
	url string
}

func CreateRecommenderClient (url string) (*RecommenderClient){
	return &RecommenderClient{url: url}
}

func (c *RecommenderClient) SendJSON(object interface{}) ([]byte, error ){
	data, err:= json.Marshal(object)
	if err !=nil {return nil, err}

	response, err:= c.sendData(data, "application/json")
	if err !=nil {return nil, err}

	return response, nil
}

func (c *RecommenderClient) sendData(data []byte, dataType string) ([]byte, error) {

	resp, err := http.Post(c.url, dataType, bytes.NewBuffer(data))
	if err != nil { return nil, err }
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err !=nil { return nil, err }

	return body, nil
}