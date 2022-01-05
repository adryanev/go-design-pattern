package main

import (
	"fmt"

	"github.com/adryanev/go-design-pattern/factory"
)

func main() {
	var dialogCreator = &factory.DialogCreator{}
	window, _ := dialogCreator.Render("windows")
	html, _ := dialogCreator.Render("html")

	windowButton := window.CreateButton()
	windowButton.OnClick(func() {
		fmt.Println("Anonymous Function Windows Button")
	})
	windowButton.Render()

	htmlButton := html.CreateButton()
	htmlButton.OnClick(func() {
		fmt.Println("Anonymous Function Html Button")
	})
	htmlButton.Render()

}
