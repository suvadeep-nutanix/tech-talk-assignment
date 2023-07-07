package main

import (
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func PublicMethod() {
	fmt.Println("vim-go")
}
