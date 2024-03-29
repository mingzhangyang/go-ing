package models

import (
	"fmt"
	"log"
	"strings"
	"bufio"
	"os"
	"path/filepath"
)

// NaiveBayes inference
type NaiveBayes struct {
	categories []string
	traindata map[string][]int
}

// encode string into int
func encode(ss []string) (map[string]int, map[int]string, []int) {
	m1 := make(map[string]int)
	m2 := make(map[int]string)
	list := make([]int, 0)
	var counter int
	for _, s := range ss {
		idx, ok := m1[s]
		if ok {
			list[idx]++
		} else {
			m1[s] = counter
			list = append(list, 1)
			m2[counter] = s
			counter++
		}
	}
	return m1, m2, list
}

// ScanFile read a file, analyze the words line by line
func ScanFile(path string) (map[string]int) {
	path, _ = filepath.Abs(path)
	file, err := os.Open(path)
	if err != nil {
		log.Panic("can't open the file")
	}
	scanner := bufio.NewScanner(file)
	var line string
	var ss []string
	m := make(map[string]int)
	for scanner.Scan() {
		line = scanner.Text()
		ss = strings.Fields(line)
		for _, s := range ss {
			s = strings.Trim(strings.ToLower(s), ",.:;?!()")
			m[strings.ToLower(s)]++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		log.Panic("error happend when scanning the file")
	}
	return m
}