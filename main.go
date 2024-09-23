package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"sync"

	"github.com/xuri/excelize/v2"
)

type Tudenn struct {
	Anv   string
	Niver int32
}

type Rann struct {
	Tudennou []Tudenn
}

type Hollad struct {
	Tudennou []Tudenn
}

const sheetName = "Sheet1"
const totalLabel = "TOTAL"

func main() {
	inputDir := flag.String("i", "", "Input directory")
	outputFile := flag.String("o", "", "Output file")
	flag.Parse()

	if *inputDir == "" || *outputFile == "" {
		fmt.Println("Usage: lignaj -i [directory] -o [output-file]")
		os.Exit(1)
	}

	inputFiles, err := listFiles(*inputDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Create a channel to collect Rann results
	rannChan := make(chan Rann, len(inputFiles))
	var wg sync.WaitGroup

	// Process each file concurrently
	for _, file := range inputFiles {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			rann := processFile(f)
			rannChan <- rann
		}(file)
	}

	// Wait for all go routines to finish and then close the channel
	go func() {
		wg.Wait()
		close(rannChan)
	}()

	// Collect all Rann from the channel
	var rannou []Rann
	for rann := range rannChan {
		rannou = append(rannou, rann)
	}

	hollad := sum(rannou)

	if err := writeOutputFile(*outputFile, hollad); err != nil {
		fmt.Println("Error writing output:", err)
	}
}

func listFiles(directory string) ([]string, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	var output []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".xlsx" {
			output = append(output, filepath.Join(directory, file.Name()))
		}
	}
	return output, nil
}

func processFile(filename string) Rann {
	input, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filename, err)
		return Rann{}
	}
	defer input.Close()

	rows, err := input.GetRows(sheetName)
	if err != nil {
		fmt.Printf("Error reading sheet %s: %v\n", sheetName, err)
		return Rann{}
	}

	var rann Rann
	for _, row := range rows {
		if len(row) < 2 || row[0] == totalLabel {
			continue
		}

		if niver, err := strconv.Atoi(row[1]); err == nil && niver != 0 {
			rann.Tudennou = append(rann.Tudennou, Tudenn{Anv: row[0], Niver: int32(niver)})
		}
	}
	return rann
}

func sum(rannou []Rann) Hollad {
	tmp := make(map[string]int32)
	for _, r := range rannou {
		for _, v := range r.Tudennou {
			tmp[v.Anv] += v.Niver
		}
	}

	hollad := make([]Tudenn, 0, len(tmp))
	for k, v := range tmp {
		hollad = append(hollad, Tudenn{Anv: k, Niver: v})
	}
	sort.Slice(hollad, func(i, j int) bool {
		return hollad[i].Niver > hollad[j].Niver
	})

	return Hollad{Tudennou: hollad}
}

func writeOutputFile(outputPath string, h Hollad) error {
	output, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer output.Close()

	w := bufio.NewWriter(output)
	defer w.Flush()

	w.WriteString("Anv,Niver\n")
	for _, v := range h.Tudennou {
		if _, err := w.WriteString(fmt.Sprintf("%s,%d\n", v.Anv, v.Niver)); err != nil {
			return err
		}
	}
	return nil
}
