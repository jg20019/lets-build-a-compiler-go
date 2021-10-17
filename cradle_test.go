package main

import (
  "io"
  "testing"
)

func TestIsDigit(t *testing.T) {
  input := '1'
  if !IsDigit(input) {
    t.Errorf("IsDigit - should recognized %c", input)
  }

  input = 'a'
  if IsDigit(input) {
    t.Errorf("IsDigit - should not recognize %c", input)
  }
}

func TestIsAlpha(t *testing.T) {
  input := 'a'
  if !IsAlpha(input) {
    t.Errorf("IsAlpha - should recognized %c", input)
  }

  input = '1'
  if IsAlpha(input) {
    t.Errorf("IsAplha - should not recognize %c", input)
  }
}

func TestGetRune(t *testing.T) {
  err := InitFromString("a") 

  if err != nil {
    t.Errorf("GetRune should have read a rune.")
  }

  if look != 'a' {
    t.Errorf("GetRune should have read the rune 'a'")
  }

  err = GetRune()
  if err == nil {
    t.Errorf("GetRune should have returned an error")
  }

  if err != io.EOF {
    t.Errorf("GetRune should have returned EOF, got %v", err)
  }

  err = InitFromString("")
  if err != io.EOF {
    t.Errorf("Initialize on an empty string should return EOF")
  }
}

func TestGetName(t *testing.T) {
  var err error
  err = InitFromString("a1")

  name, err := GetName()
  if err != nil || name != 'A' {
    t.Errorf("GetName should have returned 'A'")
    t.Errorf("\tname was '%c'", name)
    t.Errorf("\terr was %v", err)
  }

  name, err = GetName()
  if err == nil {
    t.Errorf("GetName should have returned an error")
    t.Errorf("\tname was '%c'", name)
    t.Errorf("\terr was %v", err)
  }
}

func TestGetNumber(t *testing.T) {
  InitFromString("1a")

  number, err := GetNumber()
  if err != nil {
    t.Errorf("GetNumber should have returned '1'")
    t.Errorf("\tnumber was '%c'", number)
    t.Errorf("\terr was %v", err)
  }

  number, err = GetNumber()
  if err == nil {
    t.Errorf("GetNumber should have returned an error")
    t.Errorf("\tnumber was '%c'", number)
    t.Errorf("\terr was %v", err)
  }
}


