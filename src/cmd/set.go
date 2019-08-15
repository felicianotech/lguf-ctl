package cmd

import (
	"fmt"
	"log"

	"github.com/felicianotech/go-lguf/lguf"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set brightness level of the monitor",
	Long: `Set the brightness level of an LG UltraFine 4K monitor by passing 
the level as a percentage or an integer from 0 - 54000.

Examples of good values include (without the quotes) "20%", "20000", and "99%".
Examples of bad values include (without the quotes) "120%", "10,000", and 
"59000".`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		value, err := prepBrightnessInput(args[0])
		if err != nil {
			return err
		}

		conn, err := lguf.NewConnection()
		if err != nil {
			log.Fatalf("%v", err)
		}
		defer conn.Close()

		err = conn.SetBrightness(value)
		if err != nil {
			return fmt.Errorf("%v", err)
		}

		fmt.Println("Brightness successfully set.")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
