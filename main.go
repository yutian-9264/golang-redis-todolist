/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/yutian-9264/golang-redis-todolist/cmd"
	"github.com/yutian-9264/golang-redis-todolist/db"
)

func main() {
	cmd.Execute()
	db.ExampleClient()
}
