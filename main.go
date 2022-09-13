package main

/*
 add 8
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
