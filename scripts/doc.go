// Copyright 2014 The Postmon API Team. All rights reserved.
// Use of this source code is governed by an Apache-style
// license that can be found in the README.md file.

//START_1 OMIT
/*
Package postmongo implements an easy way to use the Postmon API. // HLdoc

//END_1 OMIT
Use

The most simplest way to use:

        package main

        import (
            "fmt"
            "github.com/PostmonAPI/postmongo"
        )

        func main() {
            res, err := postmongo.BuscarCep("01419101")
            fmt.Println(res)
            fmt.Println(err)
        }


ToDo

List of things we need to improve:

    * Add support for Postmon's method rastreio
    * Improve unit tests

//START_2 OMIT
Useful links

See more about the "Postmon API" at http://postmon.com.br

See all Wrappers and API Source Code at http://github.com/PostmonAPI
*/
package postmongo // HLdoc

//END_2 OMIT
