package main

import (
	"PapeCheck/gosimhash"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

//var sentence = flag.String("sentence", "今天的天气确实适合户外运动", "Sentence for simhash")
//var topN = flag.Int("top_n", 5, "Top n of the words separated by jieba")

func main() {
	for index, arg := range os.Args {
		fmt.Println(index, " ", arg)
	}
	f1, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Open file1 error")
	}
	f2, err := os.Open(os.Args[2])
	if err != nil {
		fmt.Println("Open file2 error")
	}
	f3, err := os.OpenFile(os.Args[3], os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Open file3 error")
	}
	defer f1.Close()
	defer f2.Close()
	defer f3.Close()
	data1, err := ioutil.ReadAll(f1)
	if err != nil {
		fmt.Println("Reading file1 error")
	}
	data2, err := ioutil.ReadAll(f2)
	if err != nil {
		fmt.Println("Reading file2 error")
	}

	simhasher := gosimhash.NewSimhasher()
	//defer simhasher.Free()
	s1 := string(data1)
	s2 := string(data2)
	a1, _ := simhasher.MakeSimHasher(s1, 1000000)
	a2, _ := simhasher.MakeSimHasher(s2, 1000000)
	res, err := gosimhash.GetSimilarity(a1, a2)
	if err != nil {
		fmt.Printf("GetSimilarity error:%v\n", err)
		return
	}
	res *= 100
	strRes := strconv.FormatFloat(res, 'g', 4, 64)
	fmt.Println(strRes)
	line1 := fmt.Sprintf("文本1:%s\n", f1.Name())
	line2 := fmt.Sprintf("文本2:%s\n", f2.Name())
	line3 := fmt.Sprintf("相似度:%v%%\n", strRes)
	write := bufio.NewWriter(f3)
	_, err = write.WriteString(line1 + line2 + line3)
	if err != nil {
		fmt.Printf("Writing file error:%v\n", err)
	}
	write.Flush()
}
