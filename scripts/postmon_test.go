// Copyright 2014 The Postmon API Team. All rights reserved.
// Use of this source code is governed by an Apache-style
// license tha can be found in the README.md file

//START_1 OMIT
package postmongo

import (
	"reflect"
	"testing" // HLtest
)

func TestPostmon(t *testing.T) { // HLtest
	//END_1 OMIT
	info := map[string]interface{}{
		"complemento": "de 2161 ao fim - lado ímpar",
		"bairro":      "Cerqueira César",
		"cidade":      "São Paulo",
		"cep":         "01419101",
		"logradouro":  "Alameda Santos",
		"estado_info": map[string]interface{}{
			"area_km2":    "248.222,801",
			"codigo_ibge": "35",
			"nome":        "São Paulo",
		},
		"cidade_info": map[string]interface{}{
			"area_km2":    "1521,101",
			"codigo_ibge": "3550308",
		},
		"estado": "SP",
	}
	//START_2 OMIT
	res, err := BuscarCep("01419101")
	if err != nil {
		t.Errorf("%v", err)
	}

	if equal := reflect.DeepEqual(res, info); equal == false {
		t.Errorf("CEP 01419101 = %v, want %v", res, info) // HLtest
	}
	//END_2 OMIT
}
