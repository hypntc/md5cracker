package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		hash     = flag.String("h", "", "md5 hash")
		hashList = flag.String("hl", "", "path to hash list")
		dict     = flag.String("d", "", "path to dictionary")
	)
	flag.Parse()

	if len(os.Args) == 5 || *dict == "" {
		if *hash != "" {
			crackHash(*hash, *dict)
		}

		if *hashList != "" {
			crackHashList(*hashList, *dict)
		}
	} else {
		printUsage()
		return
	}
}

func crackHash(hash string, dictFilePath string) {
	hashExists := false
	dictFile, dictFileErr := os.Open(dictFilePath)

	if dictFileErr != nil {
		fmt.Printf("Can't read file: %s", dictFilePath)
	}

	defer dictFile.Close()
	scanner := bufio.NewScanner(dictFile)
	scanner.Split(bufio.ScanLines)

	var dictLine string

	for scanner.Scan() {
		dictLine = scanner.Text()
		fileHash := generateMd5Hash(dictLine)

		if fileHash == hash {
			hashExists = true
			break
		}
	}

	if hashExists == true {
		fmt.Println(dictLine)
	} else {
		fmt.Print("no string in the dict file matched the hash: " + hash)
	}
}

func crackHashList(hashListPath string, dictFilePath string) {
	hashfile, hashFileErr := os.Open(hashListPath)
	dictFile, dictFileErr := os.Open(dictFilePath)

	if hashFileErr != nil {
		fmt.Printf("Can't read file: %s", hashListPath)
	}

	if dictFileErr != nil {
		fmt.Printf("Can't read file: %s", dictFilePath)
	}

	defer hashfile.Close()
	defer dictFile.Close()
}

func generateMd5Hash(rawString string) string {
	hasher := md5.New()
	hasher.Write([]byte(rawString))

	return hex.EncodeToString(hasher.Sum(nil))
}

func printUsage() {
	usage := `Usage: md5cracker -h <hash> -hl <path to hashlist> -w <path to wordlist> [options...]
	-t, --number of threads
	`

	fmt.Printf(usage)
}
