/* 
Complete the method/function so that 
it converts dash/underscore delimited words into camel casing. 
The first word within the output should be capitalized only if 
the original word was capitalized (known as Upper Camel Case, also often referred to as Pascal case).

Examples

ToCamelCase("the-stealth-warrior"); // returns "theStealthWarrior"

ToCamelCase("The_Stealth_Warrior"); // returns "TheStealthWarrior"

*/

package main
import(
	"unicode"
	"regexp"
	"strings"
	
)

func main() {
	new_word :=ToCamelCase("The_Stealth_Warrior")
	new_word2 :=getCamel("The_Stealth_Warrior")
	println(new_word)
	println(new_word2)
	
	
}

func ToCamelCase(s string) string  {
	new_word :=""
	//---looping through the whole word
	symbol_pos :=0
	for pos,char :=range s {
		
		if string(char) == "-" || string(char) == "_"{
			symbol_pos = pos
			continue
		} 
		//---next character

		if pos  ==  symbol_pos + 1 && pos !=1 {
			
			if unicode.IsLower(char){
				new_word = new_word + string(unicode.ToUpper(char))

			} else {
				new_word = new_word + string(char)

			}


		}	else {
			new_word = new_word + string(char)

		}


		//println(string(char))
	}

	return new_word


	//return new_word

}



//---Second implementation
func getCamel(s string) string{
	re := regexp.MustCompile("_|-").Split(s,-1)

	new_word :=re[0]
	for _,v:= range re[1:] {
		new_word +=strings.Title(v)
	}

	return new_word
}

