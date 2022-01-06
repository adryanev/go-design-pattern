package factory_test

import (
	"testing"

	"github.com/adryanev/go-design-pattern/factory"
	"github.com/stretchr/testify/assert"
)

func TestFactory(t *testing.T) {
	dialogCreator := &factory.DialogCreator{}
	t.Run("Test Dialog Type is HTML", func(t *testing.T) {
		// setup

		//do
		htmlDialog, _ := dialogCreator.Render("html")
		//assert
		assert.IsType(t, &factory.HtmlDialog{}, htmlDialog)
	})

	t.Run("Test Dialog Type is Window", func(t *testing.T) {
		//setup

		//do
		windowDialog, _ := dialogCreator.Render("windows")

		//assert
		assert.IsType(t, &factory.WindowsDialog{}, windowDialog)
	})

	t.Run("Test Window Dialog can render a button", func(t *testing.T) {

		//setup
		windowDialog, _ := dialogCreator.Render("windows")

		//do
		windowButton := windowDialog.CreateButton()
		//assert
		assert.IsType(t, &factory.WindowsButton{}, windowButton)
	})

	t.Run("Test Html Dialog can render a button", func(t *testing.T) {

		//setup
		htmlDialog, _ := dialogCreator.Render("html")

		//do
		htmlButton := htmlDialog.CreateButton()
		//assert
		assert.IsType(t, &factory.HtmlButton{}, htmlButton)
	})

}
