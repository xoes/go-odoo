package main

import (
	"fmt"

	"github.com/xoes/go-odoo/api"
)

func connect() {
	c, err := api.NewClient("http://localhost:8069", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = c.Login("dbName", "admin", "password")
	if err != nil {
		fmt.Prinln(err.Error())
	}
}
