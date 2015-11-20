package main

import (
	"fmt"
	"github.com/upamune/go-esa/esa"
	"log"
	"os"
)

func main() {
	apikey := os.Getenv("ESA_API_KEY")
	client := esa.NewClient(apikey)

	res, err := client.Team.GetTeam("moomin")
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(res)

}
