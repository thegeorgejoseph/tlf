package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/boltdb/bolt"
    "log"
    "os"
    "github.com/atotto/clipboard"
)

func getMethod(key, identifier string) string {
    var result string
    err := db.View(func(tx *bolt.Tx) error {
        bucket := tx.Bucket([]byte(identifier)) 
        if bucket == nil {
            return fmt.Errorf("You have not set any keys into tlf yet")
        }
        value := bucket.Get([]byte(key))
        if value != nil {
            result = string(value)
        } else {
            if identifier == "value" {
                return fmt.Errorf("You have not set any values for '%s' yet", key)
            }  else {
                return fmt.Errorf("You have not set any links for '%s' yet", key)
            }
        }
        return nil
    })
    if err != nil {
        log.Fatal(err)
    }
    return result
}
// getCmd represents the get command
var getCmd = &cobra.Command{
    Use:   "get",
    Short: "Command to get the value or link of a key in the clipboard history",
    Long: `get is used to get the value or link of a key in the clipboard history.
    For example:

    tlf get -v myKey will return the value of myKey
    tlf get -l myKey will return the link of myKey

    `,
    Run: func(cmd *cobra.Command, args []string) {
        if int(cmd.Flags().NFlag()) == 0 {
            err := fmt.Errorf("You need to specify either -v or -l")
            fmt.Println(err)
            fmt.Println("Usage:")
            fmt.Println(cmd.Flags().FlagUsages())
            os.Exit(1)
        }
        var clipboard_bool bool = false
        var result string = ""
        key := args[0]
        value, _ := cmd.Flags().GetBool("value")
        link, _ := cmd.Flags().GetBool("link")
        if int(cmd.Flags().NFlag()) == 1 {
            clipboard_bool = true
        }
        if value {
            result = getMethod(key, "value")
            if result != "" {
                fmt.Printf("Value for key '%s': %s\n", key, result)
            }
        }
        if link {
            result = getMethod(key, "link")
            if result != "" {
                fmt.Printf("Link for key '%s': %s\n", key, result)
            }
        }
        if clipboard_bool {
            err := clipboard.WriteAll(result)
            if err != nil {
                log.Fatal(err)
            }
        }
    },
}

func init() {
    rootCmd.AddCommand(getCmd)

    getCmd.Flags().BoolP("value", "v", false, "Get the value of a key in the clipboard history")
    getCmd.Flags().BoolP("link", "l", false, "Get the link of a key in the clipboard history")
}
