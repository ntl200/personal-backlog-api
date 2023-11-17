package main

import (
	"fmt"
	"ntl200/personal-backlog-api/models"
	"ntl200/personal-backlog-api/models/enums"
	"github.com/gin-gonic/gin"
)

func main() {
	var d enums.Status = enums.Completed
	fmt.Println(d)
	fmt.Println(d.EnumIndex())
}