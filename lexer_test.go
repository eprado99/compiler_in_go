package main

import (
    "testing"
)

type Operator string
type Keywords string
const (
    Plus Operator = "+"
    Minus Operator = "-"
    Multiply Operator = "*"
    Divide Operator = "/"
)

const (
    Var Keywords = "var"
)

const (
    Identificador string = ""
)
func TestIdentificador(t *testing.T) {
    l := "Identificador" // Call lexeme("int")
    if l != "Identificador" {
        t.Errorf("Expected Identificador but got %v", l)
    }
}
