package main

import "fmt"

type IText interface {
	GetText() string
	AboutText() string
}

type Text struct {
	text string
}

func (t *Text) GetText() string {
	return t.text
}

func (t *Text) AboutText() string {
	return "This text: "
}

type TextDecoratorBold struct {
	wrappee IText
}

func (t *TextDecoratorBold) GetText() string {
	return "<b>" + t.wrappee.GetText() + "</b>"
}

func (t *TextDecoratorBold) AboutText() string {
	return t.wrappee.AboutText() + "bold "
}

type TextDecoratorEmphasized struct {
	wrappee IText
}

func (t *TextDecoratorEmphasized) GetText() string {
	return "<em>" + t.wrappee.GetText() + "</em>"
}

func (t *TextDecoratorEmphasized) AboutText() string {
	return t.wrappee.AboutText() + "emphasized  "
}

type TextDecoratorUnderlined struct {
	wrappee IText
}

func (t *TextDecoratorUnderlined) GetText() string {
	return "<u>" + t.wrappee.GetText() + "</u>"
}

func (t *TextDecoratorUnderlined) AboutText() string {
	return t.wrappee.AboutText() + "underlined "
}

func main() {
	text := Text{text: "Hello World!"}
	textB := TextDecoratorBold{wrappee: &text}
	textBU := TextDecoratorUnderlined{wrappee: &textB}
	textBUE := TextDecoratorEmphasized{wrappee: &textBU}

	fmt.Println(textBUE.GetText() + "\n" + textBUE.AboutText())
}
