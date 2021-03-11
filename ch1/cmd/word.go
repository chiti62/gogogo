package cmd

import (
	"log"
	"strings"

	"github.com/chiti62/gogogo/ch1/internal/word"
	"github.com/spf13/cobra"
)

const (
	ModeUpper = iota + 1
	ModeLower
	ModeWordCount
)

var str string
var mode int8
var desc = strings.Join([]string{
	"these are commands supported:",
	"1: convert word to upper case",
	"2: convert word to lower case",
	"3: count the letters of the word",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "word magic",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeWordCount:
			content = word.WordCountStr(str)
		default:
			log.Fatalf("not supported, see help word")
		}
		log.Printf("output: %s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "enter word")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "enter mode")
	//StringVarP added single letter paramter, can be used in -[n]
}
