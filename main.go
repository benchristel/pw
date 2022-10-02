package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/benchristel/pw/pw"
)

func main() {
	confirm := flag.Bool("confirm", false, "prompt for master password twice, to guard against typos")
	flag.Parse()
	domain := flag.Arg(0)
	if domain == "" {
		fmt.Printf("You must pass the domain as an argument, e.g. %s example.com\n", os.Args[0])
		return
	}

	fmt.Print("Enter your master password: ")
	masterPassword, err := readPassword()
	if err != nil {
		log.Fatal(err)
	}

	if *confirm {
		fmt.Print("--confirm: Enter it again: ")
		confirmation, err := readPassword()
		if err != nil {
			log.Fatal(err)
		}
		if confirmation != masterPassword {
			fmt.Println("Passwords do not match.")
			return
		}
	}

	sitePassword := pw.SitePassword{
		Master: masterPassword,
		Domain: domain,
		Salt: "FIXME",
		DeriveKey: pw.Scryptor{Complexity: 18}.Key,
	}.String()

	err = exec.Command("bash", "-c", "echo -n '" + sitePassword + "' | xclip -sel c").Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Your password for %s has been copied to the clipboard. Press return to clear it.", domain)
	cmd := exec.Command("bash", "-c", "read _; echo -n '' | xclip -sel c")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func readPassword() (string, error) {
	password := strings.Builder{}
	cmd := exec.Command("bash", "-c", "read -s password && echo -n \"$password\"")
	cmd.Stdin = os.Stdin
	cmd.Stdout = &password
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	fmt.Println("")
	if err != nil {
		return "", err
	}
	return password.String(), nil
}