/*
Copyright © 2023 CFET CFET CFET

*/
package shuffle

import (
	"fmt"
	"os"
	"math/rand"
	"time"
	"strconv"

	"github.com/csid-cfet/shuffler/utils"

	"github.com/spf13/cobra"
	"github.com/common-nighthawk/go-figure"
)


func main() {
    Execute()
}


func Execute() {
    if err := ShuffleCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}


var ShuffleCmd = &cobra.Command{
	Use:   "shuffle",
	Short: "Generate the order for the Wednesday standup.",
	Long: `Run from it. Hide from it. But the shuffler will always get you... Unless you took the day off like a champ, then fair play and have a good one`,

	Run: func(cmd *cobra.Command, args []string) {
		data, err := utils.ReadFromFile()

        if err != nil {
            fmt.Println("File reading error", err)
            return
        }

		if len(data) > 0 {
			rand.Seed(time.Now().UnixNano())
			rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })
	
			utils.ClearConsole()
	
			for i := 5; i > 0; i-- {
				countdownArt := figure.NewColorFigure(strconv.Itoa(i), "colossal", "green", true)
				countdownArt.Print()
				fmt.Println("")
	
				time.Sleep(1 * time.Second)
			}
	
			utils.ClearConsole()
	
			for _, line := range data {
				asciiArt := figure.NewFigure(line, "colossal", true)
				asciiArt.Print()
			}
		} else {
			fmt.Println("Empty team. Please use `shuffler team add [NAME]`")
			return
		}
	},
}