package cmd

import (
	"fmt"
	"log"
	"math"

	"github.com/felicianotech/go-lguf/lguf"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets the current brightness level of the monitor",
	Run: func(cmd *cobra.Command, args []string) {

		conn, err := lguf.NewConnection()
		if err != nil {
			log.Fatalf("%v", err)
		}
		defer conn.Close()

		value, err := conn.Brightness()
		if err != nil {
			log.Fatalf("%v", err)
		}

		brightness := int(math.Round(float64(value) / float64(lguf.MaxBrightness) * 100))
		simpleMode, _ := cmd.Flags().GetBool("simple")

		if simpleMode {
			fmt.Printf("%v%%", brightness)
		} else {
			fmt.Printf("LG UltraFine 4K brightness is: %v%%", brightness)
		}
	},
}

func init() {

	rootCmd.AddCommand(getCmd)
	getCmd.Flags().BoolP("simple", "s", false, "Return only the unit, not a sentence")
	//getCmd.Flags().StringP("unit", "u", false, "Unit to return. Choice between percent and raw")
}
