package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func scrapeUFC() {
	var fighter string
	fmt.Print("Which fighter are you looking for? ")
	fmt.Scan(&fighter)
	fighterName := strings.Split(fighter, "-")

	res, err := http.Get("https://ufc.com/athlete/" + fighter)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".hero-profile").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".hero-profile__nickname").Text()
		record := s.Find(".hero-profile__division-body").Text()
		fmt.Printf("It's time for %s. %s %s has a record of %s", title, fighterName[0], fighterName[1], record)
		fmt.Println(len(fighterName))
	})
}

func main() {
	scrapeUFC()
}
