package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
    "log"
    "net/url"
    "github.com/boltdb/bolt"
)

func setMethod(key, value, link string) {
    err := db.Update(func(tx *bolt.Tx) error {
        valueBucket, err := tx.CreateBucketIfNotExists([]byte("value"))
        if err != nil {
            log.Fatal(err)
        }
        return valueBucket.Put([]byte(key), []byte(value))
    })
    if err != nil {
        log.Fatal(err)
    }

    err = db.Update(func(tx *bolt.Tx) error {
        linkBucket, err := tx.CreateBucketIfNotExists([]byte("link"))
        if err != nil {
            log.Fatal(err)
        }
        return linkBucket.Put([]byte(key), []byte(link))
    })
    if err != nil {
        log.Fatal(err)
    }
}

func isValidURL(input string) bool {
	u, err := url.Parse(input)
	return err == nil && u.Path != ""
}

var setCmd = &cobra.Command{
    Use:   "set",
    Short: "Command to set a key value pair in the clipboard history",
    Long: `set is used to set a key value pair in the clipboard history.
    For example:

    tlf set -k myKey -v myValue -l myLink
    `,
	Run: func(cmd *cobra.Command, args []string) {
        key, _ := cmd.Flags().GetString("key")
        value, _ := cmd.Flags().GetString("value")
        link, _ := cmd.Flags().GetString("link")
        
        fmt.Printf("%d Flags founds", int(cmd.Flags().NFlag()))
        if key != "" {
            fmt.Println("Key: ", key)
        }
        if value != "" {
            fmt.Println("Value: ", value)
        }
        if link != "" {
            url_link := fmt.Sprintf("\033]8;;%s\a%s\033]8;;\a", link, link)
            fmt.Println("Link: ", url_link)
        }

        setMethod(key, value, link)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

    setCmd.Flags().StringP("key", "k", "", "Key to set in the clipboard history")
    setCmd.MarkFlagRequired("key")
    setCmd.Flags().StringP("value", "v", "", "Value to set in the clipboard history")
    setCmd.Flags().StringP("link", "l", "", "Link to set in the clipboard history")

}
