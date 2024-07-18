package main

import "github.com/phihc116/go-clean-architecture/presentation/initialize"

func main() {
	r := initialize.InitRouter()
	r.Run()
}
