package main

import (
    "flag"
    "fmt"
    "math/big"
    "os"
    "os/exec"
    "runtime"
)

const (
    RFC3526_PRIME = "FFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74020BBEA63B139B22514A08798E3404DDEF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6F44C42E9A637ED6B0BFF5CB6F406B7EDEE386BFB5A899FA5AE9F24117C4B1FE649286651ECE45B3DC2007CB8A163BF0598DA48361C55D39A69163FA8FD24CF5F83655D23DCA3AD961C62F356208552BB9ED529077096966D670C354E4ABC9804F1746C08CA18217C32905E462E36CE3BE39E772C180E86039B2783A2EC07A28FB5C55DF06F4C52C9DE2BCBF695581718399547CEA956AE515D2261898FA051015728E5A8AACAA68FFFFFFFFFFFFFFFF"
    RFC3526_GENERATOR = "2"
)

func main() {
    initCmd := flag.Bool("i", false, "Run initialization function")
    initCmdLong := flag.Bool("init", false, "Run initialization function")
    proverCmd := flag.Bool("p", false, "Run prover function")
    proverCmdLong := flag.Bool("prover", false, "Run prover function")
    verifierCmd := flag.Bool("v", false, "Run verifier function")
    verifierCmdLong := flag.Bool("verifier", false, "Run verifier function")

    flag.Parse()

    if *initCmd || *initCmdLong {
        initialize()
    } else if *proverCmd || *proverCmdLong {
        prover()
    } else if *verifierCmd || *verifierCmdLong {
        verifier()
    } else {
        fmt.Println("Usage: zkp <-i --init> <-p --prover> <-v --verifier>")
        os.Exit(1)
    }
}

func initialize() {
    p, _ := new(big.Int).SetString(RFC3526_PRIME, 16)
    g, _ := new(big.Int).SetString(RFC3526_GENERATOR, 10)
    var x big.Int

    fmt.Print("Enter shared secret: ")
    if _, err := fmt.Scan(&x); err != nil {
        fmt.Println("Error reading shared secret:", err)
        return
    }

    y := new(big.Int).Exp(g, &x, p)

    clearScreen()

    fmt.Println("Please note the values.\n")
    fmt.Println("Your agreed prime:", p.String())
    fmt.Println("Your agreed integer:", g.String())
    fmt.Println("Result of y:", y.String())
}

func prover() {
    p, _ := new(big.Int).SetString(RFC3526_PRIME, 16)
    g, _ := new(big.Int).SetString(RFC3526_GENERATOR, 10)
    var x, r big.Int

    fmt.Print("Enter your shared secret: ")
    if _, err := fmt.Scan(&x); err != nil {
        fmt.Println("Error reading shared secret:", err)
        return
    }

    fmt.Print("Enter a random integer: ")
    if _, err := fmt.Scan(&r); err != nil {
        fmt.Println("Error reading random integer:", err)
        return
    }

    c := new(big.Int).Exp(g, &r, p)

    clearScreen()

    fmt.Println("Send below lines to your verifier.\n")
    fmt.Println("c =", c.String())
    val1 := new(big.Int).Exp(g, new(big.Int).Mod(new(big.Int).Add(&x, &r), new(big.Int).Sub(p, big.NewInt(1))), p)
    fmt.Println("Result:", val1.String())
}

func verifier() {
    p, _ := new(big.Int).SetString(RFC3526_PRIME, 16)
    var y, c, val1 big.Int

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

    val2 := new(big.Int).Mod(new(big.Int).Mul(&c, &y), p)

    clearScreen()

    if val1.Cmp(val2) == 0 {
        fmt.Println("Verified: Your partner knows the shared secret.")
    } else {
        fmt.Println("Verification failed: Partner does not know the shared secret!")
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
