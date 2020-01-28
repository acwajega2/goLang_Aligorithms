/*
The parameter of accum is a string which includes only letters from a..z and A..Z.
accum("abcd") -> "A-Bb-Ccc-Dddd"
accum("RqaEzty") -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
accum("cwAt") -> "C-Ww-Aaa-Tttt"
*/

package main
import(
	"strings"
)

func main() {
	println(Accum("abcd"))
	println(Accum("RqaEzty"))
	println(Accum("cwAt"))
}

func Accum(s string) string {
	// your code
	new_word :=""
	for pos,char:= range s {

				
		for i:=0; i <= pos; i++{
			if i == 0{
				new_word +=strings.ToUpper(string(char))
			
			}else {
				new_word +=strings.ToLower(string(char))
				

			}
		}
		new_word +="-"

		

	
	}
	new_word = strings.TrimSuffix(new_word,"-")

	return new_word
}