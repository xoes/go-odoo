package main

import (
	"fmt"

	"github.com/xoes/go-odoo/api"
)

func connect() {
	c := api.Config{"http://localhost:8069", nil, nil}
	client, err := c.NewClient()
	if err != nil {
		fmt.Println(err.Error())
	}
	s := api.Session{"dbName", "user", "password"}
	err = c.Login(s)
	if err != nil {
		fmt.Prinln(err.Error())
	}
}
