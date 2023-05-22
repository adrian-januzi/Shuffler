package utils

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"bufio"
	"log"
	"os/exec"
	"runtime"
)


var defaultFonts = []string {
	"starwars",
	"standard",
	"thin",
	"speed",
	"shadow",
	"3-d",
	"5lineoblique",
	"alligator2",
	"avatar",
	"banner3-D",
	"big",
	"bulbhead",
	"cosmic",
	"jazmine",
	"larry3d",
	"o8",
	"ogre",
	"colossal",
}


func ensureShufflerFiles() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	shufflerDir := filepath.Join(usr.HomeDir, ".shuffler")
	teamFile := filepath.Join(shufflerDir, "team")
	fontsFile := filepath.Join(shufflerDir, "fonts")

	if _, err := os.Stat(teamFile); os.IsNotExist(err) {
		err := os.MkdirAll(shufflerDir, 0755)
		if err != nil {
			log.Fatal(err)
		}

		_, err = os.Create(teamFile)
		if err != nil {
			log.Fatal(err)
		}

		f, err := os.Create(fontsFile)
		if err != nil {
			log.Fatal(err)
		}

		w := bufio.NewWriter(f)
		for _, line := range defaultFonts {
			fmt.Fprintln(w, line)
		}

		if err = w.Flush(); err != nil {
			log.Fatal(err)
		}

		err = f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}


func GetTeamFilepath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("Error getting current user: %w", err)
	}

	return filepath.Join(usr.HomeDir, ".shuffler", "team"), nil
}


func ReadFromFile(filenames ...string) ([]string, error) {
	ensureShufflerFiles()

	usr, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("Error getting current user: %w", err)
	}

	filename := filepath.Join(usr.HomeDir, ".shuffler", "team")
	if len(filenames) > 0 {
		filename = filenames[0]
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %w", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error reading file: %w", err)
	}

	return lines, nil
}


func ClearConsole() {
	var clearCmd *exec.Cmd

	if runtime.GOOS == "windows" {
		clearCmd = exec.Command("cmd", "/c", "cls")
	} else {
		clearCmd = exec.Command("clear")
	}

	clearCmd.Stdout = os.Stdout
	clearCmd.Run()
}