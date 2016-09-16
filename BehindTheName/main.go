/* Acquires info about a user's name through the html of a website called behindthename.com
   and displays it intelligently */
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/stytch16/BehindTheName/NameData"
)

/* TODO: Handle issue with feminine and masculine names. They have be joined.
   Sat 03 Sep 2016 10:27:40 PM PDT */
func main() {
	/* Acquire user's first name, ignore anything else */
	clearconsole()
	fmt.Println("Hi, what's your name?")
	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	name := strings.TrimSpace(strings.Split(input, " ")[0])

	/* Process the name and check if it is valid. */
	if eCode, err := NameData.Process(name); err != nil { /* eCode = 0 implies name was found. Otherwise, eCode = 1 */
		log.Fatal(err)
	} else if eCode == 1 {
		clearconsole()
		fmt.Printf("What a unique name! Never heard that one before!\n")
		time.Sleep(2 * time.Second)
		return
	}

	/* Extract data */
	prefix, err := NameData.GetPrefixTitle() /* Mr. / Ms. */
	errorLog(err)

	namegender, err := NameData.GetGender() /* Masculine / Feminine */
	errorLog(err)

	nameusage, err := NameData.GetUsage() /* English / Japanese / ... */
	errorLog(err)

	namemeaning, err := NameData.GetMeaning() /* Full text of meaning of name */
	errorLog(err)

	/* Chat away! */
	chat := func(pause time.Duration) {
		clearconsole()
		fmt.Printf("Nice to meet you, %s %s!...\n", prefix, name)
		time.Sleep(pause)
		clearconsole()
		fmt.Printf("What a virtuous and %s name!...\n", namegender)
		time.Sleep(pause)
		clearconsole()
		fmt.Printf("Do you actually have %s roots?...\n", nameusage)
		time.Sleep(pause)
		clearconsole()
		fmt.Printf("Here's some stuff I know about your name, %s %s ... \n\n", prefix, name)
		time.Sleep(pause)
		clearconsole()

		var c int
		input := bufio.NewScanner(strings.NewReader(namemeaning))
		input.Split(bufio.ScanWords)
		for input.Scan() {
			fmt.Printf("%s ", input.Text())
			c++
			if c%7 == 0 { /* output 7 words for every line */
				fmt.Print("\n")
				time.Sleep(pause)
			}
		}
		time.Sleep(pause * 2)
		fmt.Printf("\nNice talking to you, %s %s! Hope we meet again, %[1]s %[2]s!\n", prefix, name)
	}
	chat(2 * time.Second)
}

/* clearconsole clears the console screen by calling the unix system command clear. */
func clearconsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "clearconsole: Can't clear console using unix command 'clear'")
	}
}

/*  errorLog checks to see if error is not nil. If so, logs it and ends program. */
func errorLog(e error) {
	if err != nil {
		log.Fatal(err)
	}
}
