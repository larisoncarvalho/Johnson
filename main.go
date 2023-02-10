package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Davincible/goinsta"
)

type Configuration struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	// Load username and password from config.json
	config, err := loadConfig()
	if err != nil {
		panic(err)
	}
	insta := goinsta.New(config.Username, config.Password)

	// First try to import if there is an existing config present
	insta, err = goinsta.Import(".goinsta")
	if err != nil {
		insta = goinsta.New(config.Username, config.Password)
		// Only call Login the first time you login. Next time import your config
		if err := insta.Login(); err != nil {
			panic(err)
		}
	} else {
		// method to perform a post-login sequence, such as fetching the timeline posts, and unread DM's etc
		insta.OpenApp()
	}

	// Export your configuration
	// it's useful when you want use goinsta repeatedly.
	// Export is deffered because every run insta should be exported at the end of the run
	//   as the header cookies change constantly.
	defer insta.Export(".goinsta")

	tl := insta.Timeline
	fmt.Printf("%d posts have been fetched already", len(tl.Items))
	for _, item := range tl.Items {
		// fmt.Println(item.)
		hitCount, err := checkStringForKeywords(item.Caption.Text)
		if err != nil {
			fmt.Println("Error while checking post caption for keywords " + err.Error())
		}
		if hitCount >= 2 {
			// Maybe add discord integration here
			// print for now
			fmt.Println(item.Caption)
		}
	}

}

func loadConfig() (Configuration, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return Configuration{}, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		return Configuration{}, err
	}
	return config, nil
}
