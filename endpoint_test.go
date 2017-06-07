package json

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestDecorate0(t *testing.T) {
	type P struct {
		TestField string
	}
	resp := map[string]interface{}{
		"test": "float64(123)",
	}
	handlerFunc := func(i interface{}) interface{} {
		_, ok := i.(*P)
		if !ok {
			t.Errorf("Can't assrt type %+v %T", i, i)
		}
		return resp
	}

	fn := Decorate(handlerFunc, (*P)(nil))
	jsonBody, _ := json.Marshal(map[string]interface{}{
		"TestField": "test_value",
	})
	req := httptest.NewRequest("POST", "http://"+httptest.DefaultRemoteAddr, bytes.NewReader(jsonBody))
	w := httptest.NewRecorder()
	fn(w, req)

	result := w.Result()
	body, _ := ioutil.ReadAll(result.Body)
	response := map[string]interface{}{}
	json.Unmarshal(body, &response)
	if !reflect.DeepEqual(resp, response) {
		t.Errorf("Responses not equals %+v %+v", resp, response)
	}
}
