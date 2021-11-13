//Declare a greetings package to collect related functions.
package greetings

import (
    "errors"
    "fmt"
    "math/rand"
    "time"
)

// Hello returns a greeting for the named person.
//takes a name parameter whose type is string
//a function whose name starts with a capital letter can be called by a function not in the same package. This is known in Go as an exported name.
// For example, Pizza is an exported name
//pizza do not start with a capital letter, so they are not exported i.e "unexported"
// When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.
// if you use food.pizza --> error because it is not exported, but if you use food.Pizza --> it would work
// fun <name of function>(<function parameter and its type>)<return type>
//in case of multiple return type use (<type1>, <type2>) else just use <type>
func Hello(name string) (string, error) {
	 // If no name was given, return an error with a message.
	 if name == "" {
        return "", errors.New("empty name")
    }
    // Return a greeting that embeds the name in a message.
	//declare a variable message
	//the := operator is a shortcut for declaring and initializing a variable in one line. the variable type is defined based on value on right side
    message := fmt.Sprintf(randomFormat(), name)
	//or
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)
	//return the value
    return message,nil
}

// init sets initial values for variables used in the function.
//Go executes init functions automatically at program startup, after global variables have been initialized.
func init() {
    rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
//Note that randomFormat starts with a lowercase letter, making it accessible only to code in its own package (in other words, it's not exported).
func randomFormat() string {
    // A slice of message formats.
	//When declaring a slice, you omit its size in the brackets, like this: []string. This tells Go that the size of the array underlying the slice can be dynamically changed.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}