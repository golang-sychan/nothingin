package main

/*
 add 1
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
