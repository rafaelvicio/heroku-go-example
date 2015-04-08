package controllers

import (
	"fmt"
	"net/http"
)

func GetHome(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello World!")
}
