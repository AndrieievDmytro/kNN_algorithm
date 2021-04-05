package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	trainFile string = "data\\trains.csv" //data file
	testFile  string = "data\\test.csv"   //data file
	k         int                         // k number of neigbours
)

type Flower struct {
	Params    []float64  `json:"params"`
	Name      string     `json:"name"`
	Distances []Distance `json:"distance"`
}

type Flowers struct {
	Fl []Flower
}

type Distance struct {
	Index    int     `json:"index"`
	Distance float64 `json:"distance"`
	Finded   int     `json:"finded"`
}

type FlowersNamesCount struct {
	name  string
	count int
}

func init() {
	flag.IntVar(&k, "k", k, "number of neigbours to analyze")
	flag.StringVar(&trainFile, "tr", trainFile, "train-file")
	flag.StringVar(&testFile, "ts", testFile, "test-file")
}

func readCsv(path string) ([][]string, error) {
	dataFile, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer dataFile.Close()
	if err == nil {
		// Reading text from file
		buf, rerr := ioutil.ReadAll(dataFile)
		if rerr != nil {
			return nil, err
		}

		// Parsing from comma-separated text
		r := csv.NewReader(strings.NewReader(string(buf)))
		records, err := r.ReadAll()
		if err != nil {
			return nil, err
		}
		return records, nil
	}
	return nil, err
}

func convertStrArrayToJson(records [][]string) string {
	// Converting from array of string to JSON
	jsonData := ""
	strNum := 0
	paramsLength := len(records[1]) - 1
	for _, record := range records {
		wrongStr := false
		if len(record) < paramsLength || len(record) > paramsLength+1 {
			fmt.Println("Wrong parameters count")
			wrongStr = true
		}
		flName := record[len(record)-1] // Cutting flower name
		record = record[:len(record)-1]
		strNum++
		stringArray := "["                // Opening sq bracket
		for _, arrField := range record { // Filling string representation of array
			_, err := strconv.ParseFloat(arrField, 64)
			if err != nil {
				fmt.Println("Wrong parameters type in string: ", strNum)
				wrongStr = true
			}
			stringArray += arrField + ","
		}
		stringArray = stringArray[:len(stringArray)-1] // Removing last ',' character
		stringArray += "]"                             // Closing sq bracket
		if !wrongStr {
			jsonData += "{ \"name\": \"" + flName + "\", \"params\":" + stringArray + ", \"Distance\": [] }," // Converting to JSON
		}
	}
	jsonData = "[" + jsonData[:len(jsonData)-1] + "]"
	return jsonData

}

func (f *Flowers) readData(path string) {
	records, err := readCsv(path)
	if err == nil {
		jsonData := convertStrArrayToJson(records)
		json.Unmarshal([]byte(jsonData), &f.Fl)
	} else {
		fmt.Println("Read CSV error: " + err.Error())
		os.Exit(1)
	}
}

func (ts *Flowers) calcDistances(tr *Flowers) {
	for i := 0; i < len(ts.Fl); i++ {
		for j := 0; j < len(tr.Fl); j++ {
			distance, err := euclideanDistance(tr.Fl[j].Params, ts.Fl[i].Params)
			if err != nil {
				fmt.Println("Error")
			}
			k := 0
			// Inseting with sorting
			for ; k < len(ts.Fl[i].Distances) && ts.Fl[i].Distances[k].Distance < distance; k++ {
			}
			// Add new element to slice
			ts.Fl[i].Distances = append(ts.Fl[i].Distances, Distance{Index: -1, Distance: -1.0, Finded: 0})
			copy(ts.Fl[i].Distances[k+1:], ts.Fl[i].Distances[k:]) // Move part after element for inserting
			// Fill a new element
			ts.Fl[i].Distances[k].Distance = distance // Fill distance (link)
			ts.Fl[i].Distances[k].Index = j           // Fill linked element
		}
		// fmt.Println(ts.Fl[i].Name, ts.Fl[i].Params, ts.Fl[i].Distances)
	}
}

func (flower *Flower) groupByName(tr *Flowers, itt int) []FlowersNamesCount {
	var trFlowers []FlowersNamesCount
	// Calculate count flowers with the same name

	for i := 0; i < len(flower.Distances) && i <= itt; i++ {
		j := 0
		for ; j < len(trFlowers); j++ {
			ok := (trFlowers[j].name == tr.Fl[flower.Distances[i].Index].Name)
			if ok {
				trFlowers[j].count++
				break
			}
		}
		if j == len(trFlowers) {
			trFlowers = append(trFlowers, FlowersNamesCount{name: tr.Fl[flower.Distances[i].Index].Name, count: 1})
		}
	}
	return trFlowers
}

func euclideanDistance(x, y []float64) (float64, error) {
	var sum float64
	_len := len(x)
	if _len != len(y) {
		return 0, fmt.Errorf(("The length of vectors " + fmt.Sprint(x) + " and " + fmt.Sprint(y) + " has to be "))
	}
	for i := 0; i < _len; i++ {
		sum += math.Pow(x[i]-y[i], 2)
	}
	return math.Sqrt(sum), nil
}

func printResult() {
	var accuracy float64
	var accuracies []float64
	var counter int

	tr := new(Flowers)
	tr.readData(trainFile) // Filling train data
	ts := new(Flowers)
	ts.readData(testFile) // Filling testing data
	ts.calcDistances(tr)  // Calculating distances
	if k > len(tr.Fl) {
		k = len(tr.Fl)
	}
	for j := 0; j < k; j++ {
		for i := 0; i < len(ts.Fl); i++ {
			flower := ts.Fl[i]
			// Looking for most often founded (using amount of elements with the same name)
			trFlowers := flower.groupByName(tr, j+1)
			maxCount := 0
			maxName := ""
			for l := 0; l < len(trFlowers); l++ {
				if maxCount < trFlowers[l].count {
					maxName = trFlowers[l].name
					maxCount = trFlowers[l].count
				}
			}
			// if names in ts and tr are different then hilight string with `diff`
			diff := ""
			if maxName != flower.Name {
				diff = "diff"
				counter++
			}
			fmt.Println("line ", i+1, "\t"+maxName+"\tFound:\t", maxCount, "\tcheck name from test:\t"+flower.Name+"\t", diff)
		}
		accuracy = (float64((len(ts.Fl) - counter)) / float64(len(ts.Fl))) * 100.0
		accuracies = append(accuracies, accuracy)
		counter = 0
		s := fmt.Sprintf("%.2f", accuracy)
		fmt.Println("k = ", j+1, " ", s+"%")
	}
	for g := 0; g < len(accuracies); g++ {
		fmt.Println("Accuracy for k=", g+1, " ", accuracies[g], "%")
	}

}

func main() {
	flag.Parse() //Parse comandline arguments
	printResult()
}
