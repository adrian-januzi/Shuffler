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
	"os/exec"
	"runtime"

	"github.com/csid-cfet/shuffler/utils"

	"github.com/spf13/cobra"
	"github.com/common-nighthawk/go-figure"
	"golang.org/x/term"
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
		data, err := utils.ReadFromFile("configs/team")

        if err != nil {
            fmt.Println("File reading error", err)
            return
        }

		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })

		clearConsole()

		for i := 5; i > 0; i-- {
			countdownArt := figure.NewColorFigure(strconv.Itoa(i), "colossal", "green", true)
			countdownArt.Print()
			fmt.Println("")

			time.Sleep(1 * time.Second)
		}

		clearConsole()

		for _, line := range data {
			asciiArt := figure.NewFigure(line, "colossal", true)
			asciiArt.Print()
		}
	},
}


func clearConsole() {
	var clearCmd *exec.Cmd

	if runtime.GOOS == "windows" {
		clearCmd = exec.Command("cmd", "/c", "cls")
	} else {
		clearCmd = exec.Command("clear")
	}

	clearCmd.Stdout = os.Stdout
	clearCmd.Run()
}


func getTerminalWidth() (int, error) {
	width, _, err := term.GetSize(int(os.Stdin.Fd()))

	if err != nil {
		return 0, err
	}
	
	return width, nil
}


func printCentered(s string) {
	width, _ := getTerminalWidth()

    padding := (width - len(s)) / 2
    fmt.Printf("%*s\n", len(s) + padding, s)
}