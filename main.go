package main

import (
	"fmt"
	"gcp-quota-control/entity/services"
)

func main() {
	fmt.Println("Start Quota monitoring...")
	services.Get()
}
