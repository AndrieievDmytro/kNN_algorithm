// package main

// import (
// 	"encoding/csv"
// 	"encoding/json"
// 	"errors"
// 	"flag"
// 	"fmt"
// 	"io/ioutil"
// 	"math"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// var (
// 	trainFile    string = "data\\train.txt" //data file
// 	testFile     string = "data\\test.txt"  //data file
// 	paramsNumber int    = 4                 //float parameters count
// 	neighbours   int
// )

// type Flower struct {
// 	Params    []float64  `json:"params"`
// 	Name      string     `json:"name"`
// 	Distances []Distance `json:"distance"`
// }

// type Flowers struct {
// 	Fl []Flower
// }

// type Distance struct {
// 	Index    int     `json:"index"`
// 	Distance float64 `json:"distance"`
// 	Finded   int     `json:"finded"`
// }

// func init() {
// 	flag.IntVar(&neighbours, "k", neighbours, "test-file")
// 	flag.IntVar(&paramsNumber, "p", paramsNumber, "float parameters count")
// 	flag.StringVar(&trainFile, "tr", trainFile, "train-file")
// 	flag.StringVar(&testFile, "ts", testFile, "test-file")
// }

// func readCsv(path string) ([][]string, error) {
// 	dataFile, err := os.OpenFile(path, os.O_RDONLY, 0666)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer dataFile.Close()
// 	if err == nil {
// 		// Reading text from file
// 		buf, rerr := ioutil.ReadAll(dataFile)
// 		if rerr != nil {
// 			fmt.Println("Read CSV error: " + rerr.Error())
// 		}

// 		// Parsing from comma-separated text
// 		r := csv.NewReader(strings.NewReader(string(buf)))
// 		records, err := r.ReadAll()
// 		if err != nil {
// 			fmt.Println("Parse CSV error: " + err.Error())
// 		}
// 		return records, nil
// 	}
// 	return nil, err
// }

// func convertStrArrayToJson(records [][]string) string {
// 	// Converting from array of string to JSON
// 	jsonData := ""
// 	strNum := 0
// 	for _, record := range records {
// 		wrongStr := false
// 		if len(record) < paramsNumber || len(record) > paramsNumber+1 {
// 			// fmt.Println("Wrong parameters count in " + path)
// 			fmt.Println("Wrong parameters count")
// 			wrongStr = true
// 		}
// 		var flName string
// 		if len(record) == paramsNumber+1 {
// 			flName = record[len(record)-1] // Cutting flower name
// 			record = record[:len(record)-1]
// 		}
// 		strNum++
// 		stringArray := "["                // Opening sq bracket
// 		for _, arrField := range record { // Filling string representation of array
// 			_, err := strconv.ParseFloat(arrField, 64)
// 			if err != nil {
// 				// fmt.Println("Wrong parameters type "+path+" in string: ", strNum)
// 				fmt.Println("Wrong parameters type in string: ", strNum)
// 				wrongStr = true
// 			}
// 			stringArray += arrField + ","
// 		}
// 		stringArray = stringArray[:len(stringArray)-1] // Removing last ',' character
// 		stringArray += "]"                             // Closing sq bracket
// 		if !wrongStr {
// 			jsonData += "{ \"name\": \"" + flName + "\", \"params\":" + stringArray + ", \"Distance\": [] }," // Converting to JSON
// 		}
// 	}
// 	jsonData = "[" + jsonData[:len(jsonData)-1] + "]"
// 	return jsonData

// }

// func (f *Flowers) readData(path string) {
// 	records, err := readCsv(path)
// 	if err == nil {
// 		jsonData := convertStrArrayToJson(records)
// 		json.Unmarshal([]byte(jsonData), &f.Fl)
// 	} else {
// 		fmt.Println("Read CSV error: " + err.Error())
// 	}
// }

// func (ts *Flowers) calcDistances(tr *Flowers) {
// 	var distance float64
// 	for i := 0; i < len(ts.Fl); i++ {
// 		for j := 0; j < len(tr.Fl); j++ {
// 			distance = euclideanDistance(tr.Fl[j].Params, ts.Fl[i].Params)
// 			k := 0
// 			// Inseting with sorting
// 			for ; k < len(ts.Fl[i].Distances) && ts.Fl[i].Distances[k].Distance < distance; k++ {
// 			}
// 			// Add new element to slice
// 			ts.Fl[i].Distances = append(ts.Fl[i].Distances, Distance{Index: -1, Distance: -1.0, Finded: 0})
// 			copy(ts.Fl[i].Distances[k+1:], ts.Fl[i].Distances[k:]) // Move part after element for inserting
// 			// Fill a new element
// 			ts.Fl[i].Distances[k].Distance = distance // Fill distance (link)
// 			ts.Fl[i].Distances[k].Index = j           // Fill linked element
// 		}
// 		// fmt.Println(ts.Fl[i].Name, ts.Fl[i].Params, ts.Fl[i].Distances)
// 	}
// }

// func (flower *Flower) groupByName(tr *Flowers) map[string]int {
// 	trFlowers := make(map[string]int)
// 	// Calculate count flowers with the same name
// 	for j := 0; j < len(flower.Distances) && j < neighbours; j++ {
// 		fp, ok := trFlowers[tr.Fl[flower.Distances[j].Index].Name]
// 		if ok {
// 			fp++
// 		} else {
// 			fp = 1
// 		}
// 		trFlowers[tr.Fl[flower.Distances[j].Index].Name] = fp
// 	}
// 	return trFlowers
// }

// func euclideanDistance(x, y []float64) float64 {
// 	var sum float64
// 	_len := len(x)
// 	if _len != len(y) {
// 		fmt.Println(errors.New("the length of vectors " + fmt.Sprint(x) + " and " + fmt.Sprint(y) + " has to be "))
// 	}
// 	for i := 0; i < _len; i++ {
// 		sum += math.Pow(x[i]-y[i], 2)
// 	}
// 	return math.Sqrt(sum)
// }

// func printResult() {
// 	var counter int
// 	var accuracy float64
// 	tr := new(Flowers)
// 	tr.readData(trainFile) // Filling train data
// 	ts := new(Flowers)
// 	ts.readData(testFile) // Filling testing data
// 	ts.calcDistances(tr)  // Calculating distances

// 	for i := 0; i < len(ts.Fl); i++ {
// 		flower := ts.Fl[i]
// 		// Looking for most oftern founded (using amount of elements with the same name)
// 		trFlowers := flower.groupByName(tr)
// 		maxCount := 0
// 		maxName := ""

// 		for trCurrName, trCurrVal := range trFlowers {
// 			if maxCount < trCurrVal {
// 				maxName = trCurrName
// 				maxCount = trCurrVal
// 			}
// 		}
// 		// if names in ts and tr are different then hilight string with `diff`
// 		diff := ""
// 		if maxName != flower.Name {
// 			diff = "diff"
// 			counter++
// 		}
// 		fmt.Println("line ", i+1, "\t"+maxName+"\tFound:\t", maxCount, "\tcheck name from test:\t"+flower.Name+"\t", diff)
// 	}
// 	accuracy = (float64((len(ts.Fl) - counter)) / float64(len(ts.Fl))) * 100.0
// 	s := fmt.Sprintf("%.2f", accuracy)
// 	fmt.Println("k = ", neighbours, " ", s+"%")
// }

// // func printStatistics() {
// // 	for i := 0; i < neighbours; i++ {
// // 		printResult()
// // 	}
// // }

// func main() {

// 	flag.Parse() //Parse comandline arguments
// 	printResult()
// }
