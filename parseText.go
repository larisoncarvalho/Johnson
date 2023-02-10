package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkStringForKeywords(s string) (int, error) {
	sampleText := `Spice up your Valentine's date night with an Oriental dinner date and live music. Head over to Fortune Miramar on 14th February and enjoy a delicious meal at Ramens and More, Chris will surely keep you entertained with his popular beats. And lastly, end the perfect evening with a complimentary dessert platter with your loved one. For table reservations call 0832-6637373.

	#FortuneMiramar #FortuneHotels #Goa #ValentinesOffer #ValentinesDay #ValentinesDinner #Ramen #NoodleBowl #RiceBowl #Noodles #Food #Foodie #Meal #MealOfTheDay #Noodles`

	keywordsFile, err := os.Open("keywords.txt")
	if err != nil {
		return 0, err
	}
	defer keywordsFile.Close()
	reader := bufio.NewReader(keywordsFile)
	keyword, err := reader.ReadString('\n')
	var keywords []string
	for err == nil {
		keywords = append(keywords, strings.TrimSpace(keyword))
		keyword, err = reader.ReadString('\n')
	}

	hitCount := 0
	for _, key := range keywords {
		if strings.Contains(strings.ToLower(sampleText), strings.ToLower(key)) {
			hitCount++
			fmt.Println(key)
		}
	}
	return hitCount, nil
}
