package main

import (
	"flag"
)

func main() {
	var (
		hash     = flag.String("h", "", "md5 hash")
		hashList = flag.String("hl", "", "path to hash list")
		wordList = flag.String("w", "", "path to word list")
	)
}

func crackSingleMd5Hash() {}

func setoptHelp() {
	help := `Usage: [HashFilePath] [WordListFilePath] [OPTION]...\n`
}
