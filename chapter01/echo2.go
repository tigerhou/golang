package main
import(
	"fmt"
	"os"
)

func main(){
	s,sep:="",""
	for _,arg:=range os.Args[1:]{
		s+=sep+arg
		sep=" "
	}
	fmt.Println(s)
	for index,arg:=range os.Args[1:]{
		fmt.Print(index)
		fmt.Println(arg)
	}
}