package main

import (
	"github.com/zaynkorai/eventus/pkg/api"
)

func main() {
	checkErr(api.Start())
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
