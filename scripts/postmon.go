// Copyright 2014 The Postmon API Team. All rights reserved.
// Use of this source code is governed by an Apache-style
// license that can be found in the README.md file

//START_1 OMIT
package postmongo // HLpackage

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//END_1 OMIT

// BuscarCep returns a map with all results about the CEP
// An error will be launched if any problems occur

//START_2 OMIT
func BuscarCep(cep string) (map[string]interface{}, error) { // HLpack

	res, err := http.Get("http://api.postmon.com.br/v1/cep/" + cep)

	if err != nil {
		return nil, err
	}

	var cepResult map[string]interface{}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	jsonBody := []byte(body)

	decoding := json.Unmarshal(jsonBody, &cepResult)

	return cepResult, decoding // HLpack
}

//END_2 OMIT
