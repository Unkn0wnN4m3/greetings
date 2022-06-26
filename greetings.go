package greetings

import "fmt"

// In Go, a function whose name starts with a capital letter
// can be called ba a function not in the same package

// Hello returns a greeting for the named person
func Hello(name string) string {
  // Return a greeting tahat enbeds the name in a message
  message := fmt.Sprintf("✋ Hi, $v. Welcome!", name)
  return message
}
