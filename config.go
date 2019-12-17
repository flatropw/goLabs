package main

type configList struct {
	address    string
	port       string
	delimiter  byte
	exitPhrase string
}

var config = configList{
	address:    "127.0.0.1",
	port:       "8081",
}

const delimiter = '\n'
const exitPhrase = "exit"
