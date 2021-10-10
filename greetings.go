package greetings

import (
	"fmt"
	"math/rand"
	"time"
)

func Greet(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("name is empty")
	} else {
		message := generateMessage(name)
		return fmt.Sprintf(message), nil
	}
}

func GreetV2(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Greet(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	return messages, nil
}

func generateMessage(name string) string {
	// fmt.Println(time.Now().Format("3:04PM"))
	rand.Seed(time.Now().UnixNano())
	messages := []string{
		"Hello, %s",
		"Hi, %s",
		"Hey, %s",
		"Howdy, %s",
		"Greetings, %s",
		"Salutations, %s",
		"Good day, %s",
		"Good evening, %s",
		"Good morning, %s",
		"Good night, %s",
	}
	return fmt.Sprintf(messages[rand.Intn(len(messages))], name)
}
