//Declare a greetings package to collect related functions.
package greetings

import "fmt"

// Hello returns a greeting for the named person.
//takes a name parameter whose type is string
//a function whose name starts with a capital letter can be called by a function not in the same package. This is known in Go as an exported name.
// For example, Pizza is an exported name
//pizza do not start with a capital letter, so they are not exported i.e "unexported"
// When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.
// if you use food.pizza --> error because it is not exported, but if you use food.Pizza --> it would work
// fun <name of function>(<function parameter and its type>)<return type>
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
	//declare a variable message
	//the := operator is a shortcut for declaring and initializing a variable in one line. the variable type is defined based on value on right side
    message := fmt.Sprintf("Hi, %v. Welcome to the Go Tutorial!", name)
	//or
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)
	//return the value
    return message
}
