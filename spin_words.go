/*
Write a function that takes in a string of one or more words, and returns the same string, but with all five or more letter words reversed (Just like the name of this Kata). Strings passed in will consist of only letters and spaces. Spaces will be included only when more than one word is present.

Examples: spinWords( "Hey fellow warriors" ) => returns "Hey wollef sroirraw" spinWords( "This is a test") => returns "This is a test" spinWords( "This is another test" )=> returns "This is rehtona test"


*/

package main

import(
	"regexp"
	"strings"
)

func main() {

	println(SpinWords( "Hey fellow warriors"))
	println(SpinWords( "This is a test"))
	println(SpinWords( "This is another test"))
	

}

func SpinWords(str string) string {
	re :=regexp.MustCompile(" ").Split(str, -1)
	word:=[]string{}

	for _,v :=range re{
		if len(v) >= 5{
			reversed_word :=""
						
			for i:=len(v) -1; i >= 0;i--{
				reversed_word  += string(v[i])
			}

			word = append(word,reversed_word)
			

		} else{
			
			word = append(word,v)
		}
	}

	new_word:=strings.Join(word, " ")
	return new_word
  }