package modules

import (
	"fmt"
	"os"
)

func WriteFile(inputLogin, inputPassword string) {
	text := inputLogin + "\n" + inputPassword
	file, err := os.Create("login.txt")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(text)

	fmt.Println("Done.")
}
