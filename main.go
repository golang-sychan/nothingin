package main

func main() {
	println("hello joey!")
    addsth1
}

func demo() {
	fmt.Println("demo")
}

//add sth
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
