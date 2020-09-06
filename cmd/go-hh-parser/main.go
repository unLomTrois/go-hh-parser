package main

import (
	"encoding/json"
	"go-hh-parser/internal/app/core"
	"log"
	"sync"
)

func main() {
	Core := core.New()

	urls := []string{
		"https://api.hh.ru/vacancies/38840984",
	}

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			vacancy, err := Core.Requests.GetVacancy(url)
			if err != nil {
				log.Fatal(err)
			}
			defer wg.Done()

			// data := *vacancy.Data

			// get keyskills
			for _, skill := range vacancy.Keyskills {
				log.Println(skill["name"])
			}

			// log.Println(string(*vacancy.Data))
			jsonData, err := json.MarshalIndent(&vacancy, "", "  ")
			if err != nil {
				log.Fatal(err)
			}

			log.Println(string(jsonData))
		}(url)
	}
	wg.Wait()
}
