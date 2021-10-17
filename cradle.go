package main

import (
  "bufio"
  "errors"
  "fmt"
  "os"
  "strings"
  "unicode"
)

///////////////////////////////////////////////
// Program Input 
//////////////////////////////////////////////

// The lookahead character
var look rune 

// The input stream
var inputStream *bufio.Reader

// Read a rune from inputStream
func GetRune() error {
  var err error
  look, _, err = inputStream.ReadRune()
  return err
}

/////////////////////////////////////////////
// Error Handling
/////////////////////////////////////////////

// Outputs an error message
func Error(s string) string{
  return fmt.Sprintf("\n Error: %s.\n", s)
}


// Report what was expected. 
func Expected(s string) error {
  return errors.New(fmt.Sprintf("%s expected.", s))
}

///////////////////////////////////////////////////
// Recognizing Input 
//////////////////////////////////////////////////

// Match a specific input character
func Match(x rune) error {
  if look == x {
    return GetRune()
  }
  return Expected(fmt.Sprintf("'%c'", x))
}

// Recognize an alpha character
func IsAlpha(x rune) bool {
  return unicode.IsLetter(x)
}

// Recognize a decimal digit
func IsDigit(x rune) bool {
  return unicode.IsDigit(x)
}


// Get an identifier 
func GetName() (rune, error) {
  if !IsAlpha(look) {
    return 0, Expected("Name")
  }

  return unicode.ToUpper(look), GetRune()
}

// Get a number
func GetNumber() (rune, error) {
  if !IsDigit(look) {
    return 0, Expected("Integer")
  }

  return look, GetRune()
}

///////////////////////////////////////
// Emitters
//////////////////////////////////////

// Output a string with tab
func Emit(s string) {
  fmt.Printf("\t%s", s)
}

// Output a string with tab and newline
func EmitLn(s string) {
  fmt.Printf("\t%s\n", s)
}


// Initializes input stream 
func Init() error {
  inputStream = bufio.NewReader(os.Stdin)
  return GetRune()
}

// Intializes input stream from string
func InitFromString(s string) error {
  inputStream = bufio.NewReader(strings.NewReader(s))
  return GetRune()
}
