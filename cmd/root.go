/*
Copyright © 2024 donclaveau3

*/
package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"

	"github.com/spf13/cobra"
)

var ByteCountRequested bool
var LineCountRequested bool
var WordCountRequested bool
var CharCountRequested bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ccwc [path to file]",
	Short: "This is a custom version of wc for a coding challenge",
	Long: `ccwc is a CLI app that imitates the behavior of wc.
This application is a solution to a coding challenge found here:
https://codingchallenges.fyi/challenges/challenge-wc`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		openFile := os.Stdin
		message := ""
		if len(args) > 0 {
			filePath := args[0]
			message = filePath
			f, err := os.Open(filePath)
			if err != nil {
				fmt.Println(err)
			}
			openFile = f
		}

		b, c, w, l, err := scanForStats(openFile)
		if err != nil {
			fmt.Println(err)
		}

		if !(cmd.Flags().NFlag() > 0) {
			ByteCountRequested = true
			LineCountRequested = true
			WordCountRequested = true
		}

		if ByteCountRequested {
			message = fmt.Sprintf("%s %s", strconv.Itoa(b), message)
		}
		if CharCountRequested {
			message = fmt.Sprintf("%s %s", strconv.Itoa(c), message)
		}
		if WordCountRequested {
			message = fmt.Sprintf("%s %s", strconv.Itoa(w), message)
		}
		if LineCountRequested {
			lineFmtString := "%s %s"
			if !(cmd.Flags().NFlag() > 0) {
				lineFmtString = "  %s  %s"
			}
			message = fmt.Sprintf(lineFmtString, strconv.Itoa(l), message)
		}

		fmt.Println(message)
	},
}

func scanForStats(file io.Reader) (int, int, int, int, error) {
	r := bufio.NewReader(file)
	buf := make([]byte, 32*1024)
	byteCount := 0
	charCount := 0
	charSep := []byte("")
	wordCount := 0
	startOfThisIsSpace := false
	endOfLastWasSpace := true
	lineCount := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		switch {
		case err == io.EOF:
			return byteCount, charCount, wordCount, lineCount, nil

		case err != nil:
			return byteCount, charCount, wordCount, lineCount, err
		}
		byteCount += c
		charCount += bytes.Count(buf[:c], charSep) // "1 + number of Unicode code points" https://golangdoc.github.io/pkg/1.8/bytes/index.html#example_Count
		charCount -= 1
		lineCount += bytes.Count(buf[:c], lineSep) // Credit to JimB - https://stackoverflow.com/a/24563853
		wordCount += len(bytes.Fields(bytes.TrimSpace(buf[:c])))
		startOfThisIsSpace = unicode.IsSpace(rune(buf[0]))
		wordWasSplitBetweenBuffers := !endOfLastWasSpace && !startOfThisIsSpace
		if wordWasSplitBetweenBuffers {
			wordCount -= 1 // word was counted twice
		}
		endOfLastWasSpace = unicode.IsSpace(rune(buf[c-1]))
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ccwc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolVarP(&ByteCountRequested, "bytes", "c", false, "print the byte counts")
	rootCmd.Flags().BoolVarP(&LineCountRequested, "lines", "l", false, "print the newline counts")
	rootCmd.Flags().BoolVarP(&WordCountRequested, "words", "w", false, "print the word counts")
	rootCmd.Flags().BoolVarP(&CharCountRequested, "chars", "m", false, "print the character counts")
}
