package File

import (
	"bufio"
	"fmt"
	"os"
)

func Open() (*os.File, *os.File, *os.File) {
	//for index, arg := range os.Args {
	//	fmt.Println(index, " ", arg)
	//}
	f1, err := os.Open(os.Args[1])
	//f1, err := os.Open("D:/test/orig.txt")
	if err != nil {
		fmt.Printf("Open file1 error:%v\n", err)
	}
	f2, err := os.Open(os.Args[2])
	//f2, err := os.Open("D:/test/orig_0.8_del.txt")
	if err != nil {
		fmt.Printf("Open file2 error:%v\n", err)
	}
	f3, err := os.OpenFile(os.Args[3], os.O_WRONLY|os.O_CREATE, 0666)
	//f3, err := os.OpenFile("D:/test/out.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Open file3 error:%v\n", err)
	}
	return f1, f2, f3
}

func OutPut(f *os.File, data string) {
	write := bufio.NewWriter(f)
	_, err := write.WriteString(data)
	if err != nil {
		fmt.Printf("Writing file error:%v\n", err)
	}
	write.Flush()
}
