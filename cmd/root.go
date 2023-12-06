package cmd

import (
	"os"
	"github.com/spf13/cobra"
    "github.com/boltdb/bolt"
    "log"
    "path/filepath"
)

var db *bolt.DB

var rootCmd = &cobra.Command{
	Use:   "tlf",
	Short: "tiny little finder - your super tiny CLI clipboard history manager",
    Long: `tiny little finder.
Your super tiny CLI clipboard history manager.
You can use it to save and retrieve key, [value, link] pairs.
For example:

tlf set -k myKey -v myValue -l myLink
tlf get myKey -l
tlf get myKey -v
tlf get myKey -h 
`,
}

func Execute() {
    var err error
    
    home, err := os.UserHomeDir()
    if err != nil {
        log.Fatal(err)
    }
    tlsPath := filepath.Join(home, "tls")
    if err != nil {
        log.Fatal(err)
    }
    err = os.MkdirAll(tlsPath, os.ModePerm)
    if err != nil {
        log.Fatal(err)
    }
    err = os.Chdir(tlsPath)
    if err != nil {
        log.Fatal(err)
    }
    db, err = bolt.Open("tlf.db", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


