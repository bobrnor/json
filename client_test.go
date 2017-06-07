package json

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestPost0(t *testing.T) {
	client := Client{}
	if err := client.Post("", nil, nil); err == nil {
		t.Errorf("bad url => should produce error")
	}
}

func TestPost1(t *testing.T) {
	client := Client{}
	err := client.Post("", map[string]interface{}{
		"bad": make(chan int),
	}, nil)
	if err == nil {
		t.Errorf("bad data => should produce error")
	}
	t.Log(err.Error())
}

func TestPost2(t *testing.T) {
	client := Client{}
	err := client.Post("", map[string]interface{}{}, nil)
	if err == nil {
		t.Errorf("bad url => should produce error")
	}
	t.Log(err.Error())
}

func TestPost3(t *testing.T) {
	resp := map[string]interface{}{
		"test_key": "test_value",
	}
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsonResp, _ := json.Marshal(resp)
		w.Write(jsonResp)
	}))

	client := Client{}
	response := map[string]interface{}{}
	err := client.Post(testServer.URL, map[string]interface{}{}, &response)
	if err != nil {
		t.Errorf("error %+v", err.Error())
	}

	if !reflect.DeepEqual(resp, response) {
		t.Errorf("Responses are not equals %+v %+v", resp, response)
	}
}

func TestPost4(t *testing.T) {
	resp := map[string]interface{}{
		"test_key_0": "test_value",
		"test_key_1": float64(-123),
		"test_key_2": float64(123),
		"test_key_4": float64(123.321),
	}
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsonResp, _ := json.Marshal(resp)
		w.Write(jsonResp)
	}))

	client := Client{}
	response := map[string]interface{}{}
	err := client.Post(testServer.URL, map[string]interface{}{}, &response)
	if err != nil {
		t.Errorf("error %+v", err.Error())
	}

	if !reflect.DeepEqual(resp, response) {
		t.Errorf("Responses are not equals %+v %+v", resp, response)
		for k, v := range resp {
			t.Errorf("%s %T %v", k, v, v)
		}

		for k, v := range response {
			t.Errorf("%s %T %v", k, v, v)
		}
	}
}

func TestPost5(t *testing.T) {
	type St struct {
		Testkey0 string
		// Testkey1 int64
		// Testkey2 uint64
		// Testkey3 time.Time
		// Testkey4 float64
	}
	resp := St{
		"test_value",
		// int64(-123),
		// uint64(123),
		// time.Now(),
		// float64(123.321),
	}
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsonResp, _ := json.Marshal(resp)
		w.Write(jsonResp)
	}))

	client := Client{}
	response := St{}
	err := client.Post(testServer.URL, map[string]interface{}{}, &response)
	if err != nil {
		t.Errorf("error %+v", err.Error())
	}

	if !reflect.DeepEqual(resp, response) {
		t.Errorf("Responses are not equals %+v %+v", resp, response)
	}
}
