package main

import (
	"bufio" // read input from STDIN
	"fmt"   // print
	"log"
	"math/rand" //random
	"os"
	"strconv" // convert Str to Int and vice versa

	// process string
	"time"
)

func main() {

	var array [3]string
	array[0] = "Huyen"
	array[1] = "Ling"
	array[2] = "Sakura"

	fmt.Println("Number of secret person 0: " + array[0] +
		", 1: " + array[1] +
		", 2: " + array[2])

	// random a number
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	secretPersonNumber := r.Intn(3)

	// read STDIN
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter a number: ")
	// input, _ := reader.ReadString('\n')
	// input = strings.TrimSuffix(input, "\n") // cut the last character '\n' in input string
	// answer, err := strconv.Atoi(input)      // convert Str to Int

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter a number: ")
	scanner.Scan()
	input := scanner.Text()
	answer, err := strconv.Atoi(input)

	if err != nil {
		log.Panic(err)
	}

	if answer == secretPersonNumber {
		fmt.Printf("Congratulations!! You answer is correct, %s is the secret person.\n", array[answer])
	} else {
		fmt.Printf("Sorry!! You answer is incorrect, "+strconv.Itoa(secretPersonNumber)+": %s is the secret person.\n", array[secretPersonNumber])
	}
}
