package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"
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

	var dictFileLine string

	for scanner.Scan() {
		dictFileLine = scanner.Text()
		fileHash := generateMd5Hash(dictFileLine)

		if fileHash == hash {
			hashExists = true
			break
		}
	}

	if hashExists == true {
		fmt.Println(dictFileLine)
	} else {
		fmt.Printf("no string in the dict file matched the hash: %s", hash)
	}
}

func crackHashList(hashListPath string, dictFilePath string) {
	hashFile, hashFileErr := os.Open(hashListPath)
	dictFile, dictFileErr := os.Open(dictFilePath)

	if hashFileErr != nil {
		fmt.Printf("Can't read file: %s", hashListPath)
		return
	}

	if dictFileErr != nil {
		fmt.Printf("Can't read file: %s", dictFilePath)
		return
	}

	defer hashFile.Close()
	defer dictFile.Close()

	hashFileScanner := bufio.NewScanner(hashFile)
	hashFileScanner.Split(bufio.ScanLines)

	var hashFileLine string
	var dictFileLine string
	hashExists := false

	for hashFileScanner.Scan() {
		hashFileLine = strings.ToLower(hashFileScanner.Text())

		dictFileScanner := bufio.NewScanner(dictFile)
		dictFileScanner.Split(bufio.ScanLines)
		for dictFileScanner.Scan() {
			dictFileLine = dictFileScanner.Text()
			dictFileHash := generateMd5Hash(dictFileLine)

			if dictFileHash == hashFileLine {
				hashExists = true
				break
			}
		}

		if hashExists == true {
			fmt.Printf("Hash %s matched %s \n", hashFileLine, dictFileLine)
		}

		// We must reset the file for the wordlist.
		dictFile.Seek(0, 0)
		hashExists = false
	}
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
