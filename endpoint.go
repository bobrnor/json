package jsonep

import (
	"encoding/json"
	"net/http"
)

func Decorate(fn func(interface{}) interface{}, param interface{}) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		if err := read(req, &param); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		result := fn(param)

		if err := write(rw, result); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	}
}

func read(req *http.Request, param interface{}) error {
	defer req.Body.Close()

	decoder := json.NewDecoder(req.Body)
	return decoder.Decode(&param)
}

func write(rw http.ResponseWriter, data interface{}) error {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	rw.Write(dataJSON)
	return nil
}
