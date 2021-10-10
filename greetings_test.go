package greetings

import (
	"testing"
	"regexp"
)

func TestGreetName(t *testing.T){
	name:= "Rick"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg,err := Greet(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Greet(%s) = %q, %v, want match for %#q, nil`,name,msg,err,want)
	}
}


func TestGreetEmpty(t *testing.T){
	msg,err := Greet("")
	if msg!= "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`,msg,err)
	}
}