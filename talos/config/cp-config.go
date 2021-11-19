package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {

	var src string
	var dest string

	if len(os.Args) != 2 {
		fmt.Println("Please provide the number of copies you want!")
		return
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Err: Input must be a number!")
		return
	}

	for i := 1; i <= number; i++ {
		//controlplane.yaml part

		src = "controlplane.yaml"
		dest = "controlplane-" + strconv.Itoa(i) + ".yaml"

		bytesRead, err := ioutil.ReadFile(src)

		if err != nil {
			log.Fatal(err)
		}

		err = ioutil.WriteFile(dest, bytesRead, 0644)

		if err != nil {
			log.Fatal(err)
		}

		//worker.yaml part

		src = "worker.yaml"
		dest = "worker-" + strconv.Itoa(i) + ".yaml"

		bytesRead, err = ioutil.ReadFile(src)

		if err != nil {
			log.Fatal(err)
		}

		err = ioutil.WriteFile(dest, bytesRead, 0644)

		if err != nil {
			log.Fatal(err)
		}

	}
}
