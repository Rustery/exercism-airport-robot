// Airport Robot is a simple package that can greet you in different languages.
package airportrobot

import "fmt"

// Greeter is the interface that wraps the basic Greet method.
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

// Italian is a type that implements the Greeter interface.
type Italian struct {
}

func (s Italian) LanguageName() string {
	return "Italian"
}

func (s Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

// Portuguese is a type that implements the Greeter interface.
type Portuguese struct {
}

func (s Portuguese) LanguageName() string {
	return "Portuguese"
}

func (s Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}

// SayHello is a function that returns a greeting message in the language of the given Greeter.
func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}
