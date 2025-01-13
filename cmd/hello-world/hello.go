package helloworld

const englishHelloPrefix = "Hello, "

func Hello(personName string) string {
	if personName == "" {
		personName = "World"
	}
	return englishHelloPrefix + personName
}
