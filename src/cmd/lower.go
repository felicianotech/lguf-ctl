package cmd

import (
	"fmt"
	"log"

	"github.com/felicianotech/go-lguf/lguf"
	"github.com/spf13/cobra"
)

var lowerCmd = &cobra.Command{
	Use:   "lower",
	Short: "Lower brightness level of the monitor",
	Long: `Lower the brightness level of an LG UltraFine 4K monitor by passing 
the delta as a percentage or an integer from 1 - 43999.

Without an argument, will try to lower by 10%.

Examples of good values include (without the quotes) "20%", "20000", and "99%".
Examples of bad values include (without the quotes) "120%", "10,000", and 
"59000".`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) != 1 {
			args = append(args, "10%")
		}

		value, err := prepBrightnessInput(args[0])
		if err != nil {
			return err
		}

		conn, err := lguf.NewConnection()
		if err != nil {
			log.Fatalf("%v", err)
		}
		defer conn.Close()

		curValue, err := conn.Brightness()
		if err != nil {
			return fmt.Errorf("%v", err)
		}

		if curValue == lguf.MinBrightness {

			fmt.Println("Brightness already at minimim level.")
			return nil
		}

		newValue := curValue - value

		if newValue < lguf.MinBrightness {
			newValue = lguf.MinBrightness
		}

		err = conn.SetBrightness(newValue)
		if err != nil {
			return fmt.Errorf("%v", err)
		}

		log.Printf("Brightness successfully lowered.")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(lowerCmd)
}
