package main

import "fmt"

const (
  spanish = "Spanish"
  french = "French"
  englishHelloPrefix = "Hello, "
  spanishHelloPrefix = "Hola, "
  frenchHelloPrefix = "Bonjour, "
)

func Hello(name, language string) string {
  if name == "" {
    name = "World"
  }
  return greetingPrefix(language) + name
}

//In our function signature we have made a named return value (prefix string)
//This will create a variable called prefix in your function.
func greetingPrefix(language string) (prefix string) {
  switch language {
    case french:
      prefix = frenchHelloPrefix
    case spanish:
      prefix = spanishHelloPrefix
    default:
      prefix = englishHelloPrefix
  }
  return
}

func main() {
	fmt.Println(Hello("Andrei", "English"))
}
