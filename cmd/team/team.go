/*
Copyright © 2023 CFET CFET CFET

*/
package team

import (
	"fmt"
	"os"
	"strings"

	"github.com/csid-cfet/shuffler/utils"

	"github.com/spf13/cobra"
)


func main() {
    Execute()
}


func init() {
	TeamCmd.AddCommand(listTeamCmd, addMemberCmd, removeMemberCmd)
}


func Execute() {
    if err := TeamCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}


var TeamCmd = &cobra.Command{
	Use:   "team",
	Short: "CFET Team",
	Long: `CFET Team Gang Gang`,

	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("Please provide a subcommand.")
	},
}


var listTeamCmd = &cobra.Command{
    Use:   "list",
    Short: "Lists the people to be shuffled.",
    Long:  `Lists all the members that will be shuffled for a classic wednesday standup where it devolves into ChatGPT mumbo jumbo, random articles and chit chat.`,

	Run: func(cmd *cobra.Command, args []string) {
		lines, err := utils.ReadFromFile("configs/team")
		if err != nil {
			fmt.Println(err)
			return
		}
		
		for _, line := range lines {
			fmt.Println(line)
		}
	},
}


var addMemberCmd = &cobra.Command{
	Use:   "add \"[member name]\"",
	Short: "Add someone to shuffler.",
	Long:  `Add another dreaded soul to the shuffler of doom 'n' gloom.`,

	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.OpenFile("configs/team", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)

		if err != nil {
			fmt.Println("Error opening file", err)
			return
		}

		defer f.Close()

		if _, err := f.WriteString("\n" + args[0]); err != nil {
			fmt.Println("Error writing to file", err)
		}

		fmt.Println(args[0], " has been added.")
	},
}


var removeMemberCmd = &cobra.Command{
	Use:   "remove \"[member name]\"",
	Short: "Add someone to shuffler.",
	Long:  `Release someone from the shuffler's evil grasps.`,

	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		lines, err := utils.ReadFromFile("configs/team")
		if err != nil {
			fmt.Println(err)
			return
		}
		
		var newLines []string
		for _, line := range lines {
			if strings.ToLower(line) != strings.ToLower(args[0]) {
				newLines = append(newLines, line)
			}
		}

		f, err := os.Create("configs/team")
		if err != nil {
			fmt.Println("Error opening file", err)
			return
		}
		defer f.Close()

		for _, line := range newLines {
			if _, err := f.WriteString(line + "\n"); err != nil {
				fmt.Println("Error writing to file", err)
				return
			}
		}

		fmt.Println("Removed member:", args[0])
	},
}