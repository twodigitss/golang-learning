package main
//https://go.dev/doc/tutorial/workspaces

import "fmt"
import "golang.org/x/example/hello/reverse"

func main(){
	fmt.Println(reverse.String("Hello"), reverse.Int(4455))
}
