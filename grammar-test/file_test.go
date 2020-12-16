//golang文件操作练习记录
package grammar

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var name = "./test.txt"
var name2 = "./test2.txt"
var name3 = "./test3.txt"

func QTest_File(t *testing.T) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("file=%v", file)
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			log.Fatal(err)
		}
		fmt.Print(str)
	}

}

func QTest_File2(t *testing.T) {
	reader, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reader)
	fmt.Println(string(reader))
}

func Test_File3(t *testing.T) {
	file, err := os.OpenFile(name2, os.O_WRONLY|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	str := "hello, world"
	writer := bufio.NewWriter(file)
	for x := 0; x < 5; x++ {
		writer.WriteString(str)
	}
	writer.Flush()
}

func Test_File5(t *testing.T) {
	file, err := os.Stat(name3)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal(err)
		} else {
			log.Println(err)
		}
	}
	fmt.Println(file)
}
