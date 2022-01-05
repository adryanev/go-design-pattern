package factory

import "fmt"

type Button interface {
	Render() string
	OnClick(fn func())
}

type WindowsButton struct {
}

func (wb *WindowsButton) Render() string {
	fmt.Printf("Windows Button is rendered.")
	return "Windows Button"
}

func (wb *WindowsButton) OnClick(fn func()) {
	fmt.Printf("Windows Button's clicked.")
	fn()
}

type HtmlButton struct {
}

func (hb *HtmlButton) Render() string {
	fmt.Printf("Html button is rendered.")
	return "Html Button"
}

func (hb *HtmlButton) OnClick(fn func()) {
	fmt.Printf("Html Button's Clicked")
	fn()
}
