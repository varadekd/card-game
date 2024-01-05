// The file is responsible for generating the default list of cards required for the
// application to run smoothly

package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/varadekd/card-game/model"
)

var DEFAULT_SUIT_SEQUENCE = []string{"SPADES", "DIAMONDS", "CLUBS", "HEARTS"}
var CARD_SEQUENCE = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

func GenerateDefaultDeck() error {
	// fetching the file location from env variable
	filePath, err := GetEnvVariable("DEFAULT_CARDS_FILE_STORAGE")

	if err != nil {
		return err
	}

	cards := []model.Card{}

	for _, suit := range DEFAULT_SUIT_SEQUENCE {
		for _, value := range CARD_SEQUENCE {
			card := model.Card{
				Value: value,
				Suit:  suit,
				Code:  fmt.Sprintf("%s%s", value, string(suit[0])),
			}
			cards = append(cards, card)
		}
	}

	// Marshalling the generated cards to store it in a file
	cardsToStrings, err := json.Marshal(cards)

	if err != nil {
		fmt.Errorf("we encountered an error '%s' while marshalling the generated cards\n", err.Error())
		return err
	}

	// Writing the generate string to the file.
	err = ioutil.WriteFile(filePath, cardsToStrings, 0)

	if err != nil {
		fmt.Errorf("we encountered an error '%s' while writing cards to the file\n", err.Error())
		return err
	}

	fmt.Println("Default set of cards generated successfully.")

	return nil
}

func GetEnvVariable(variable string) (string, error) {
	val, variableFound := os.LookupEnv(variable)

	if !variableFound {
		return val, fmt.Errorf("we were unable to find variable %s in the application environment", variable)
	}

	if val == "" {
		return val, fmt.Errorf("no value found for variable %s in the application environment", variable)
	}

	return val, nil
}
