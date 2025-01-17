package game

import (
	"fmt"
	"github.com/ghadd/candy-wars/api"
	"github.com/ghadd/candy-wars/config"
	"github.com/ghadd/candy-wars/database"
	"log"
)

func handleUpdateMessage(client *api.Client, update api.Update) {
	message := update.Message

	if message.Text == "/start" {
		handleStartMessage(client, message)
	} else if message.Text == "/help" {
		handleHelpMessage(client, message)
	} else if message.Text == "/rules" {
		handleRulesMessage(client, message)
	} else if fitsState(message.FromUser, config.StateChangingName) {
		handleChangeNickNameMessage(client, message)
	}

	// DEV OPTION
	if message.Text == "/reset" && message.FromUser.ID == config.DevID {
		handleResetMessage(client, message)
	}
}

func handleStartMessage(client *api.Client, message api.UpdateMessage) {
	dbh, err := database.NewDBHandler()
	if err != nil {
		log.Println(err)
	}

	contains, err := dbh.ContainsUser(message.FromUser)
	if err != nil {
		log.Println(err)
	}

	if contains {
		_, err = client.SendMessage(api.Message{
			ChatID:       message.FromUser.ID,
			Text:         fmt.Sprintf("Hello, %s! Welcome back!", message.FromUser.FirstName),
			InlineMarkup: startMarkup,
		})
	} else {
		_, err = client.SendMessage(api.Message{
			ChatID:       message.FromUser.ID,
			Text:         fmt.Sprintf("Hello, %s! Welcome to CandyWarGO!", message.FromUser.FirstName),
			InlineMarkup: startMarkup,
		})

		// adding user to database if it is not there
		err = dbh.InsertUser(message.FromUser)
		if err != nil {
			log.Println(err)
		}
	}

	if err != nil {
		log.Println(err)
	}
}

func handleHelpMessage(client *api.Client, message api.UpdateMessage) {
	handleHelpAndRules(client, message)
}

func handleRulesMessage(client *api.Client, message api.UpdateMessage) {
	handleHelpAndRules(client, message)
}

func handleHelpAndRules(client *api.Client, message api.UpdateMessage){
	// todo refer to user lang code in next releases
	langCode := "en"
	msg := getMessage(message.Text, langCode)

	_, err := client.SendMessage(api.Message {
		ChatID: message.FromUser.ID,
		Text: msg,
	})
	if err != nil {
		log.Println(err)
	}
}

func handleChangeNickNameMessage(client *api.Client, message api.UpdateMessage) {
	success := true
	dbh, err := database.NewDBHandler()
	if err != nil {
		log.Println(err)
		success = false
	}

	err = dbh.Update("users", "nickname", message.Text, "telegram_id", message.FromUser.ID)
	if err != nil {
		log.Println(err)
		success = false
	}

	err = dbh.Update("users", "state", config.StateNone, "telegram_id", message.FromUser.ID)
	if err != nil {
		log.Println(err)
		success = false
	}

	var msg string
	if success {
		msg = "Nickname changed successfully."
	} else {
		msg = "Sorry. Some error happened."
	}

	_, err = client.SendMessage(
		api.Message{
			ChatID:       message.FromUser.ID,
			Text:         msg,
			InlineMarkup: startMarkup,
		})
}

// --------------------
func handleResetMessage(client *api.Client, message api.UpdateMessage) {
	err := database.ResetTables()
	if err == nil {
		_, err = client.SendMessage(api.Message{
			ChatID: message.FromUser.ID,
			Text: "Ok. Tables are reset.",
		})
	} else {
		_, err = client.SendMessage(api.Message{
			ChatID: message.FromUser.ID,
			Text: "Ok. Tables are reset.",
		})
	}
	if err != nil {
		log.Println(err)
	}
}