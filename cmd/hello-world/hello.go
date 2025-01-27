package helloworld

const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(personName, language string) string {
	if personName == "" {
		personName = "World"
	}

	if language == spanish {
		return spanishHelloPrefix + personName
	}

	return englishHelloPrefix + personName
}
