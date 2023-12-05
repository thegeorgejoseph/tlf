package cmd

import (
	"os"
	"github.com/spf13/cobra"
    "github.com/boltdb/bolt"
    "log"
)

var db *bolt.DB

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tlf",
	Short: "A super simple CLI clipboard history manager",
    Long: `This is a CLI clipboard history manager.
You can use it to save and retrieve key, [value, link] pairs.
For example:

tlf set -k myKey -v myValue -l myLink
tlf get myKey -l
tlf get myKey -v
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

    var er error
    db, er = bolt.Open("tlf.db", 0600, nil)
    if er != nil {
        log.Fatal(er)
    }
    defer db.Close()

    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


