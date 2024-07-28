package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	meetingCmd := flag.Bool("m", false, "Run meeting function")
	meetingCmdLong := flag.Bool("meeting", false, "Run meeting function")
	proverCmd := flag.Bool("p", false, "Run prover function")
	proverCmdLong := flag.Bool("prover", false, "Run prover function")
	verifierCmd := flag.Bool("v", false, "Run verifier function")
	verifierCmdLong := flag.Bool("verifier", false, "Run verifier function")

	flag.Parse()

	if *meetingCmd || *meetingCmdLong {
		meeting()
	} else if *proverCmd || *proverCmdLong {
		prover()
	} else if *verifierCmd || *verifierCmdLong {
		verifier()
	} else {
		fmt.Println("Usage: zkp <-m --meeting> <-p --prover> <-v --verifier>")
		os.Exit(1)
	}
}

func meeting() {
	var p, g, x big.Int

	fmt.Print("Enter upon agreed prime: ")
	if _, err := fmt.Scan(&p); err != nil {
		fmt.Println("Error reading prime:", err)
		return
	}

	fmt.Print("Enter upon agreed integer: ")
	if _, err := fmt.Scan(&g); err != nil {
		fmt.Println("Error reading integer:", err)
		return
	}

	fmt.Print("Enter secret integer: ")
	if _, err := fmt.Scan(&x); err != nil {
		fmt.Println("Error reading secret integer:", err)
		return
	}

	y := new(big.Int).Exp(&g, &x, &p)

	clearScreen()

	fmt.Println("Please note the values.\n")
	fmt.Println("Your agreed prime:", p.String())
	fmt.Println("Your agreed integer:", g.String())
	fmt.Println("Result of y:", y.String())
}

func prover() {
	var p, g, x, r big.Int

	fmt.Print("Enter noted prime: ")
	if _, err := fmt.Scan(&p); err != nil {
		fmt.Println("Error reading prime:", err)
		return
	}

	fmt.Print("Enter noted integer: ")
	if _, err := fmt.Scan(&g); err != nil {
		fmt.Println("Error reading integer:", err)
		return
	}

	fmt.Print("Enter your secret integer: ")
	if _, err := fmt.Scan(&x); err != nil {
		fmt.Println("Error reading secret integer:", err)
		return
	}

	fmt.Print("Enter a random integer: ")
	if _, err := fmt.Scan(&r); err != nil {
		fmt.Println("Error reading random integer:", err)
		return
	}

	c := new(big.Int).Exp(&g, &r, &p)

	clearScreen()

	fmt.Println("Send below lines to your verifier.\n")
	fmt.Println("c =", c.String())
	val1 := new(big.Int).Exp(&g, new(big.Int).Mod(new(big.Int).Add(&x, &r), new(big.Int).Sub(&p, big.NewInt(1))), &p)
	fmt.Println("Result:", val1.String())
}

func verifier() {
	var p, y, c, val1 big.Int

	fmt.Print("Enter noted prime: ")
	if _, err := fmt.Scan(&p); err != nil {
		fmt.Println("Error reading prime:", err)
		return
	}

	fmt.Print("Enter noted y: ")
	if _, err := fmt.Scan(&y); err != nil {
		fmt.Println("Error reading y:", err)
		return
	}

	fmt.Print("Enter the received c: ")
	if _, err := fmt.Scan(&c); err != nil {
		fmt.Println("Error reading c:", err)
		return
	}

	fmt.Print("Enter the received result: ")
	if _, err := fmt.Scan(&val1); err != nil {
		fmt.Println("Error reading result:", err)
		return
	}

	val2 := new(big.Int).Mod(new(big.Int).Mul(&c, &y), &p)

	clearScreen()

	if val1.Cmp(val2) == 0 {
		fmt.Println("Your communications partner knows x.")
	} else {
		fmt.Println("Your communications partner does not know x!")
	}
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
