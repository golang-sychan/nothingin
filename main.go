package main

/*
 add 8
add 20
 add 19
 add 17
 ad doadd 18
add 16
 add 15
 add 14
oadd 12
 add 10
 add 13
 add 9
 add 7
 add 1
 add 6
 add 5
 add 4
 add  2
 add 3
*/

func main() {
	println("hello joey!")
}

func demo() {
	fmt.Println("demo")
}

func Run() error {
  	fmt.Printf("App run ...\n")
	return nil
}

// wrap error for common
func wrap_err() []error {
  return nil
}

type data struct {

  Name string 
  Data []byte
  Addr string `json:"addr"`
}
