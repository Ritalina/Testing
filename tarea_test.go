package main

import "testing"
        

func TestSuma(t *testing.T) {
        total := suma(5, 5)
    if total != 10 {
       t.Errorf("Suma incorrecta")
    }
}
