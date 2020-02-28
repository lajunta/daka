package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	studentsFile *os.File
	recordsFile  *os.File
	datesFile    *os.File
	err          error
	dk_dates     []string //打卡日期
)

func init() {

	recordsFile, err = os.Open("./current.csv")
	if err != nil {
		fmt.Println("打开数据文件出错 ::", err)
		fmt.Println("数据文件名为 current.csv")
		os.Exit(1)
	}

	studentsFile, err = os.Open("./students.csv")
	if err != nil {
		fmt.Println("打开学生名单文件出错:", err)
		fmt.Println("学生名单文件名为 students.csv")
		os.Exit(1)
	}

	datesFile, err = os.Open("./dates.txt")
	if err != nil {
		fmt.Println("打开日期文件出错:", err)
		fmt.Println("日期文件名为 dates.txt")
		os.Exit(1)
	}

}

func main() {
	stary := studentList()
	rowary := recordList()
	datesList()
	anait(stary, rowary)
}

func csvWriter(csvData [][]string) {
	// 1. Open the file
	outFileName := "output.csv"
	outFile, err := os.OpenFile(outFileName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatalln("写文件失败")
	}
	defer outFile.Close()
	outFile.WriteString("\xEF\xBB\xBF")
	// prevent chinese wrong display
	if err != nil {
		fmt.Println("创建输出文件时出错: ", err)
	}

	// 2. Initialize the writer
	writer := csv.NewWriter(outFile)

	// 3. Write all the records
	err = writer.WriteAll(csvData) // returns error
	if err != nil {
		fmt.Println("保存文件时出错:", err)
	}
}

// rowary每一行格式为：
// 打卡日期,班级,姓名,...
func anait(stary []string, rowary [][]string) {
	ary := [][]string{}
	header := []string{"班级", "姓名", "学号", "打卡次数", "打卡日期", "缺少日期", "缺少天数", "连续打卡"}
	ary = append(ary, header)
	for i := 0; i < len(stary); i++ {
		stinfo := strings.Split(stary[i], "-")
		aryinner := []string{}
		count := 0
		dates := []string{}

		for j := 0; j < len(rowary); j++ {
			banji := rowary[j][1]
			name := rowary[j][2]
			if banji == stinfo[0] && name == stinfo[1] {
				count++
				dates = append(dates, rowary[j][0])
			}
		}
		missing_dates := missing(dk_dates, dates)
		missing_days := strconv.Itoa(len(missing_dates))
		unique_dates := strings.Join(unique(dates), " ")
		miss_dates := strings.Join(missing_dates, " ")
		seriality := "否"
		if len(missing_dates) == 0 {
			seriality = "是"
		}
		aryinner = append(aryinner, stinfo[0], stinfo[1], stinfo[2], strconv.Itoa(count), unique_dates, miss_dates, missing_days, seriality)
		//fmt.Println(i, stinfo[0], stinfo[1], stinfo[2], count, unique(dates), missing_dates, missing_days)
		ary = append(ary, aryinner)
		csvWriter(ary)
	}

}

func missing(full []string, b []string) (ary []string) {
	for _, vf := range full {
		n := 0
		for _, vb := range b {
			if vf == vb {
				n++
			}
		}
		if n == 0 {
			ary = append(ary, vf)
		}
	}
	sort.Strings(ary)
	return ary
}

func unique(elements []string) []string {

	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []string{}
	for key, _ := range encountered {
		result = append(result, key)
	}

	sort.Strings(result)
	return result

}

func studentList() []string {
	defer studentsFile.Close()
	reader := csv.NewReader(studentsFile)
	if _, err := reader.Read(); err != nil {
		panic(err)
	}

	ary := []string{}
	for i := 0; ; i = i + 1 {
		row, err := reader.Read()
		if err == io.EOF {
			break // reached end of the file
		} else if err != nil {
			fmt.Println("An error encountered :", err)
			os.Exit(1)
		}
		str := row[0] + "-" + row[1] + "-" + row[2]
		ary = append(ary, str)
	}
	return ary
}

func datesList() {
	defer datesFile.Close()
	scanner := bufio.NewScanner(datesFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		dk_dates = append(dk_dates, scanner.Text())
	}

}

func recordList() [][]string {
	defer recordsFile.Close()

	ary := [][]string{}

	reader := csv.NewReader(recordsFile)
	reader.Read()

	for i := 0; ; i = i + 1 {
		record, err := reader.Read()
		if err == io.EOF {
			break // reached end of the file
		} else if err != nil {
			fmt.Printf("第 %d 行有错误发生 \n", i)
			continue
		}
		ary = append(ary, record)
	}

	return ary
}
