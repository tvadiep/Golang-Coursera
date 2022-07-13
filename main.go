package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main(){
	var str string;
	fmt.Println("Enter your string: ");


	input := bufio.NewReader(os.Stdin)

	lineString,_  := input.ReadString('\n')
	str = strings.ToLower(lineString)[0:len(lineString)-1]; // to lower and remove the enter at the end of the string


	if(len(str)<3 || str[0]!=105 || str[len(str)-1]!= 110){
		fmt.Println("Not Found!");
		fmt.Println(len(str))
		return;
	}
	for i:=1; i<len(str)-1; i++ {
		if(str[i]==97){
			fmt.Println("Found!");
			return;
		}
	}
}