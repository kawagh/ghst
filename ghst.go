package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func main() {
	ghq_root_paths, err := exec.Command("ghq", "root", "-all").Output()
	if err != nil {
		log.Fatal(err)
	}
	grps := strings.Split(string(ghq_root_paths), "\n")
	for _, grp := range grps {
		fmt.Println(grp)
		files, err := ioutil.ReadDir(grp)
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range files {
			fmt.Println(f.Name())
		}
	}
}
