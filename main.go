package main

import (
        "bufio"
        "fmt"
        "log"
        "os"
        "os/exec"
        "strings"
)

const (
        GreenColor = "\033[0;32m"
        RedColor   = "\033[0;31m"
        NoColor    = "\033[0m"
)

func main() {
        if len(os.Args) < 2 {
                fmt.Println("Usage: rcheck domains.txt")
                os.Exit(1)
        }
        inputFile := os.Args[1]

        file, err := os.Open(inputFile)
        if err != nil {
                log.Fatal(err)
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                domain := scanner.Text()
                checkRecords(domain)
        }

        if err := scanner.Err(); err != nil {
                log.Fatal(err)
        }
}

func checkRecords(domain string) {
        spfCmd := exec.Command("dig", "+short", "TXT", domain)
        spfOutput, err := spfCmd.Output()
        if err != nil {
                log.Fatalf("error running dig command for SPF: %s", err)
        }
        spf := string(spfOutput)

        dmarcCmd := exec.Command("dig", "+short", "TXT", "_dmarc."+domain)
        dmarcOutput, err := dmarcCmd.Output()
        if err != nil {
                log.Fatalf("error running dig command for DMARC: %s", err)
        }
        dmarc := string(dmarcOutput)

        spfFound := strings.Contains(spf, "v=spf")
        dmarcFound := strings.Contains(dmarc, "v=DMARC")

        if spfFound {
                fmt.Printf("%sFound SPF record%s for %s\n", RedColor, NoColor, domain)
        } else {
                fmt.Printf("%sNot Found SPF record%s for %s\n", GreenColor, NoColor, domain)
        }

        if dmarcFound {
                fmt.Printf("%sFound DMARC record%s for %s\n", RedColor, NoColor, domain)
        } else {
                fmt.Printf("%sNot Found DMARC record%s for %s\n", GreenColor, NoColor, domain)
        }

        fmt.Println("--------------------------")
}
