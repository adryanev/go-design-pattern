package factory

import "fmt"

type Dialog interface {
	CreateButton() Button
}

type DialogCreator struct{}

func (d *DialogCreator) Render(dialogType string) (Dialog, error) {
	switch dialogType {
	case "windows":
		return &WindowsDialog{}, nil
	case "html":
		return &HtmlDialog{}, nil

	}
	return nil, fmt.Errorf("wrong dialog type passed")
}

type WindowsDialog struct{}

func (d *WindowsDialog) CreateButton() Button {
	return &WindowsButton{}
}

type HtmlDialog struct{}

func (d *HtmlDialog) CreateButton() Button {
	return &HtmlButton{}
}
