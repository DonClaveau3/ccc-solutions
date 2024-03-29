/*
Copyright © 2024 donclaveau3

*/
package cmd

import (
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
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		message := filePath
		if ByteCountRequested {
			byteCount := calculateBytes(filePath)
			message = fmt.Sprintf("%s %s", strconv.FormatInt(byteCount, 10), filePath)
		}
		if LineCountRequested {
			lineCount := getLineCount(filePath)
			message = fmt.Sprintf("%s %s", strconv.Itoa(lineCount), filePath)
		}
		if WordCountRequested {
			wordCount := getWordCount(filePath)
			message = fmt.Sprintf("%s %s", strconv.Itoa(wordCount), filePath)
		}
		if CharCountRequested {
			charCount := getCharCount(filePath)
			message = fmt.Sprintf("%s %s", strconv.Itoa(charCount), filePath)
		}
		fmt.Println(message)
	},
}

func getCharCount(filePath string) int {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	c, err := charCounter(f)
	if err != nil {
		fmt.Println(err)
	}

	return c
}

func charCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	sep := []byte("")

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], sep) // "1 + number of Unicode code points" https://golangdoc.github.io/pkg/1.8/bytes/index.html#example_Count
		count -= 1
		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func calculateBytes(filePath string) int64 {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	fileInfo, err := f.Stat()
	if err != nil {
		fmt.Println(err)
	}
	return fileInfo.Size()
}

func getLineCount(filePath string) int {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	c, err := lineCounter(f)
	if err != nil {
		fmt.Println(err)
	}

	return c
}

// Credit to JimB - https://stackoverflow.com/a/24563853
func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func getWordCount(filePath string) int {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	c, err := wordCounter(f)
	if err != nil {
		fmt.Println(err)
	}

	return c
}
func wordCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	startOfThisIsSpace := false
	endOfLastWasSpace := true
	for {
		c, err := r.Read(buf)
		count += len(bytes.Fields(bytes.TrimSpace(buf[:c])))
		startOfThisIsSpace = unicode.IsSpace(rune(buf[0]))
		if !endOfLastWasSpace && !startOfThisIsSpace {
			count -= 1 //word was split between buffers and counted twice
		}
		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
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
