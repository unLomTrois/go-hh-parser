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
			vacancy, err := Core.Requests.GetFullVacancy(url)
			if err != nil {
				log.Fatal(err)
			}
			defer wg.Done()

			// get keyskills
			for _, skill := range vacancy.Keyskills {
				log.Println(skill["name"])
			}

			jsonData, err := json.MarshalIndent(&vacancy, "", "  ")
			if err != nil {
				log.Fatal(err)
			}

			log.Println(string(jsonData))
		}(url)
	}
	wg.Wait()
}
