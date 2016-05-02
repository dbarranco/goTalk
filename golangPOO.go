package main

import (
	"fmt"
	"log"
	"errors"
)

type Nodo struct{
	parametro1 int
	parametro2 string
	parametro3 byte
}

func(n *Nodo)setParametro1(valor int)error{
	n.parametro1 = valor
}

func(n *Nodo)getParametro3()byte{
	return parametro3
}