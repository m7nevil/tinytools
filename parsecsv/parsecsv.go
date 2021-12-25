package parsecsv

import (
	"bufio"
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var fields = map[string]string{
	",苯,":   "苯",
	",甲苯,":  "甲苯",
	"对间二甲苯": "对间二甲苯",
	"邻二甲苯":  "邻二甲苯",
	"总计":    "总计",
}

var indexes = []string{
	",苯,",
	",甲苯,",
	"对间二甲苯",
	"邻二甲苯",
	"总计",
}

var data = map[string][]string{}

func Run() {
	filename := "./4.csv"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	for key, _ := range fields {
		data[key] = []string{}
	}

	getDataFromCsv(filename)

	log.Println(data)
	outputCsv(data)
}

func outputCsv(data map[string][]string) {
	output, _ := os.Create("result.csv")
	defer output.Close()
	output.WriteString("\xEF\xBB\xBF")

	str := ""
	for _, key := range indexes {
		str += "\"" + fields[key] + "\"" + ","
	}
	str = strings.TrimRight(str, ",")
	output.WriteString(str + "\n")

	max := 0.0
	for _, key := range indexes {
		max = math.Max(max, float64(len(data[key])))
	}
	tmp := int(max)

	for i := 0; i < tmp; i++ {
		str = ""
		for _, key := range indexes {
			if i >= len(data[key]) {
				str += ""
			} else {
				str += data[key][i]
			}
			str += ","
		}
		str = strings.TrimRight(str, ",")
		output.WriteString(str + "\n")
	}
}

func getDataFromCsv(filename string) {

	input, e := os.Open(filename)
	defer input.Close()
	if e != nil {
		log.Fatalln(e)
	}

	buf := bufio.NewScanner(input)
	for {
		if !buf.Scan() {
			break
		}
		line := buf.Text()
		line, _ = GbkToUtf8(line)
		parseLine(line)
	}
}

func parseLine(line string) {
	for key, _ := range fields {
		if !strings.Contains(line, key) {
			continue
		}

		items := strings.Split(line, ",")
		len := len(items)
		for i := len - 1; i >= 0; i-- {
			_, e := strconv.ParseFloat(items[i], 64)
			if e == nil {
				data[key] = append(data[key], items[i])
				return
			}
		}
	}
}

func GbkToUtf8(s string) (string, error) {
	reader := transform.NewReader(bytes.NewReader([]byte(s)), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return "", e
	}
	return string(d), nil
}

func Utf8ToGbk(s string) (string, error) {
	reader := transform.NewReader(bytes.NewReader([]byte(s)), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return "", e
	}
	return string(d), nil
}
