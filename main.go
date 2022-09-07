package main

func main() {
	println("hello joey!")
}

func demo() {
	fmt.Println("demo")
}

//add st 1111h
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
