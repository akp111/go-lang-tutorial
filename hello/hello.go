// Declare a main package (a package is a way to group functions, and it's made up of all the files in the same directory)
package main

// Import the popular fmt package, which contains functions for formatting text, including printing to the console.
import (
    "fmt"
    "log"
    "github.com/akp111/greetings"
)
// For using external packages
// import "rsc.io/quote"
//Implement a main function to print a message to the console. A main function executes by default when you run the main package.
func main() {
     // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)
    // Shows the time of error
    // log.SetFlags(1)
    // Request a greeting message.
    message, err := greetings.Hello("")
    //Request for success
    // message, err := greetings.Hello("")
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }
    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(message)
}