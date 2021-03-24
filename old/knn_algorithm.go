package old

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
// 	trainFile    string = "Data\\train.txt" //data file
// 	testFile     string = "Data\\test.txt"  //data file
// 	paramsNumber int    = 4                 //float parameters count
// )

// type Flower struct {
// 	Params    []float64  `json:"params"`
// 	Name      string     `json:"name"`
// 	Distances []Distance `json:"distance"`
// }

// type FlowerParams struct {
// 	inFileIndex int        `json:"fileidx"`
// 	Params      []float64  `json:"params"`
// 	Distances   []Distance `json:"distance"`
// }

// type FlowersWithParams struct {
// 	Fp []FlowerParams
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
// 	flag.IntVar(&paramsNumber, "p", paramsNumber, "float parameters count")
// 	flag.StringVar(&trainFile, "tr", trainFile, "train-file")
// 	flag.StringVar(&testFile, "ts", testFile, "test-file")
// }

// func (f *Flowers) readData(path string) {
// 	dataFile, err := os.OpenFile(path, os.O_RDONLY, 0666)
// 	if err != nil {
// 		return
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

// 		// Converting from array of string to JSON
// 		jsonData := ""
// 		strNum := 0
// 		for _, record := range records {
// 			wrongStr := false
// 			if len(record) < paramsNumber {
// 				fmt.Println("Wrong parameters count in " + path)
// 				wrongStr = true
// 			}
// 			if len(record) > paramsNumber+1 {
// 				fmt.Println("Wrong parameters count in " + path)
// 				wrongStr = true
// 			}
// 			var flName string
// 			if len(record) == paramsNumber+1 {
// 				flName = record[len(record)-1] // Cutting flower name
// 				record = record[:len(record)-1]
// 			}
// 			strNum++
// 			stringArray := "["                // Opening sq bracket
// 			for _, arrField := range record { // Filling string representation of array
// 				_, err := strconv.ParseFloat(arrField, 64)
// 				if err != nil {
// 					fmt.Println("Wrong parameters type "+path+" in string: ", strNum)
// 					wrongStr = true
// 				}
// 				stringArray += arrField + ","
// 			}
// 			stringArray = stringArray[:len(stringArray)-1] // Removing last ',' character
// 			stringArray += "]"                             // Closing sq bracket
// 			if !wrongStr {
// 				jsonData += "{ \"name\": \"" + flName + "\", \"params\":" + stringArray + ", \"Distance\": [] }," // Converting to JSON
// 			}
// 		}
// 		jsonData = "[" + jsonData[:len(jsonData)-1] + "]"
// 		// Converting from JSON to templated structure
// 		json.Unmarshal([]byte(jsonData), &f.Fl)
// 	} else {
// 		fmt.Println("Read file error: " + err.Error())
// 	}
// }

// func euclideanDistance(X, Y []float64) float64 {

// 	_len := len(X)
// 	if _len != len(Y) {
// 		fmt.Println(errors.New("the length of vectors " + fmt.Sprint(X) + " and " + fmt.Sprint(Y) + " has to be "))
// 	}

// 	var x, y float64

// 	var sum float64
// 	for i := 0; i < _len; i++ {
// 		x = X[i]
// 		y = Y[i]
// 		sum += math.Pow(x-y, 2)
// 	}
// 	return math.Sqrt(sum)
// }

// func main() {
// 	var distance float64
// 	flag.Parse() //Parse comandline arguments
// 	tr := new(Flowers)
// 	tr.readData(trainFile)
// 	ts := new(Flowers)
// 	ts.readData(testFile)

// 	// for i := 0; i < len(tr.Fl); i++ {
// 	// 	for j := 0; j < len(ts.Fl); j++ {
// 	// 		distance = euclideanDistance(tr.Fl[i].Params, ts.Fl[j].Params)
// 	// 		k := 0
// 	// 		for ; k < len(tr.Fl[i].Distances) && tr.Fl[i].Distances[k].Distance < distance; k++ {
// 	// 		}
// 	// 		tr.Fl[i].Distances = append(tr.Fl[i].Distances, Distance{Index: -1, Distance: -1.0, Finded: 0})
// 	// 		copy(tr.Fl[i].Distances[k+1:], tr.Fl[i].Distances[k:])
// 	// 		tr.Fl[i].Distances[k].Distance = distance
// 	// 		tr.Fl[i].Distances[k].Index = j
// 	// 	}
// 	// 	fmt.Println(tr.Fl[i].Name, tr.Fl[i].Params, tr.Fl[i].Distances)
// 	// }

// 	for i := 0; i < len(ts.Fl); i++ {
// 		for j := 0; j < len(tr.Fl); j++ {
// 			distance = euclideanDistance(tr.Fl[j].Params, ts.Fl[i].Params)
// 			k := 0
// 			for ; k < len(ts.Fl[i].Distances) && ts.Fl[i].Distances[k].Distance < distance; k++ {
// 			}
// 			ts.Fl[i].Distances = append(ts.Fl[i].Distances, Distance{Index: -1, Distance: -1.0, Finded: 0})
// 			copy(ts.Fl[i].Distances[k+1:], ts.Fl[i].Distances[k:])
// 			ts.Fl[i].Distances[k].Distance = distance
// 			ts.Fl[i].Distances[k].Index = j
// 		}
// 		fmt.Println(ts.Fl[i].Name, ts.Fl[i].Params, ts.Fl[i].Distances)
// 	}

// 	// Converting to map
// 	tsFlowers := make(map[string]FlowersWithParams)

// 	// for i := 0; i < len(ts.Fl); i++ {
// 	// 	fl := ts.Fl[i]
// 	// 	fp, ok := tsFlowers[fl.Name]
// 	// 	if ok {
// 	// 		fp.Fp = append(fp.Fp, FlowerParams{inFileIndex: i, Distances: fl.Distances, Params: fl.Params})
// 	// 	} else {
// 	// 		ns := make([]FlowerParams, 1)
// 	// 		ns[0] = FlowerParams{inFileIndex: i, Distances: fl.Distances, Params: fl.Params}
// 	// 		fp = FlowersWithParams{Fp: ns}
// 	// 	}
// 	// 	tsFlowers[fl.Name] = fp
// 	// }
// 	// for i := 0; i < len(ts.Fl); i++ {
// 	// 	fl := ts.Fl[i]
// 	// 	fp, ok := tsFlowers[strconv.Itoa(i)]
// 	// 	if ok {
// 	// 		fp.Fp = append(fp.Fp, FlowerParams{inFileIndex: i, Distances: fl.Distances, Params: fl.Params})
// 	// 	} else {
// 	// 		ns := make([]FlowerParams, 1)
// 	// 		ns[0] = FlowerParams{inFileIndex: i, Distances: fl.Distances, Params: fl.Params}
// 	// 		fp = FlowersWithParams{Fp: ns}
// 	// 	}
// 	// 	tsFlowers[strconv.Itoa(i)] = fp
// 	// }

// 	// for fln, fl := range tsFlowers {
// 	// 	fmt.Println("Name: ", fln, " count:", len(fl.Fp))
// 	// 	for i := 0; i < len(fl.Fp); i++ {
// 	// 		fmt.Println("\ttrIdx: ", i, " dist count: ", len(fl.Fp[i].Distances))
// 	// 		fmt.Print("\t\t")
// 	// 		for j := 0; j < len(fl.Fp[i].Distances); j++ {
// 	// 			lookingIdx := fl.Fp[i].Distances[j].Index
// 	// 			fmt.Print("lookingIdx: ", lookingIdx, " ")
// 	// 			for k := 0; k < len(fl.Fp); k++ {
// 	// 				if k != i {
// 	// 					for l := 0; l < len(fl.Fp[k].Distances) && l < 10; l++ {
// 	// 						if fl.Fp[k].Distances[l].Index == lookingIdx {
// 	// 							fl.Fp[i].Distances[j].Finded++
// 	// 							break
// 	// 						}
// 	// 					}
// 	// 				}
// 	// 			}
// 	// 			fmt.Print("Found: ", fl.Fp[i].Distances[j].Finded, "\t")
// 	// 		}
// 	// 		fmt.Println()
// 	// 	}
// 	// }
// 	for fln, fl := range tsFlowers {
// 		// fmt.Println("Name: ", fln, " count:", len(fl.Fp))
// 		trFlowers := make(map[string]int)
// 		for i := 0; i < len(fl.Fp); i++ {
// 			fmt.Println("trIdx: ", i, " dist count: ", len(fl.Fp[i].Distances))
// 			for j := 0; j < len(fl.Fp[i].Distances) && j < 12; j++ {
// 				fp, ok := trFlowers[tr.Fl[fl.Fp[i].Distances[j].Index].Name]
// 				if ok {
// 					fp++ // = append(fp, FlowerParams{inFileIndex: i, Distances: fl.Distances, Params: fl.Params})
// 				} else {
// 					fp = 0
// 				}
// 				trFlowers[tr.Fl[fl.Fp[i].Distances[j].Index].Name] = fp

// 				// fmt.Print("\t\t")
// 				// fmt.Println(tr.Fl[fl.Fp[i].Distances[j].Index].Name+" Found: ", fl.Fp[i].Distances[j].Index+1, "<->", fl.Fp[i].inFileIndex+1, " "+fln+" ", fp, "\t")
// 			}
// 			fmt.Println()
// 			maxCount := 0
// 			maxName := ""
// 			for trCurrName, trCurrVal := range trFlowers {
// 				if maxCount < trCurrVal {
// 					maxName = trCurrName
// 					maxCount = trCurrVal
// 				}
// 			}
// 			index, _ := strconv.Atoi(fln)
// 			fmt.Println(maxName+" Found: ", maxCount, " times at line", fl.Fp[i].inFileIndex+1, " "+ts.Fl[index].Name+" ", "\t")
// 		}
// 	}

// 	// os.Exit(0)

// 	// for fln, fl := range tsFlowers {
// 	// 	fmt.Println("Name: ", fln, " count:", len(fl.Fp))
// 	// 	for i := 0; i < len(fl.Fp); i++ {
// 	// 		fmt.Println("\ttrIdx: ", i, " dist count: ", len(fl.Fp[i].Distances))
// 	// 		for j := 0; j < len(fl.Fp[i].Distances); j++ {
// 	// 			if fl.Fp[i].Distances[j].Finded+1 == len(fl.Fp) {
// 	// 				fmt.Print("\t\t")
// 	// 				fmt.Println(ts.Fl[fl.Fp[i].Distances[j].Index].Name, "Found: ", fl.Fp[i].Distances[j].Index, "<->", fl.Fp[i].inFileIndex, "found count:", fl.Fp[i].Distances[j].Finded, "\t")
// 	// 			}
// 	// 		}
// 	// 	}
// 	// }

// }
