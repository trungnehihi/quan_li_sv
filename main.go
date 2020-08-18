package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func  tb (t float64,l float64,h float64) float64 {
	return 	(t + l + h) / 3
}

type students struct {
	name 		string
	diemtoan 	float64
	diemli   	float64
	diemhoa		float64
	dtb 		float64
}

func read (path string) ([]string , error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var bang []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bang = append(bang, scanner.Text())
	}
	return bang, scanner.Err()
}

func string_conv(sv students) string {
	s := sv.name
	s += fmt.Sprintf("- %.2f - %.2f - %.2f - %.2f\n",sv.diemtoan, sv.diemli, sv.diemhoa, sv.dtb)
	return s
}

func xuat (sv [] students )  {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()
	for _,v := range sv {
		file.WriteString(string_conv(v))
	}
}

func main() {
	bang,err := read("C:\\Users\\ACER\\go\\src\\main\\sv.txt")
	if err != nil {
		fmt.Println("loi")
		fmt.Println(err)
	}
	var sv []students
	scf := strconv.ParseFloat
	for _,line := range bang {
		data := strings.Split(line,"-")
		t,_ := scf(data[1],32)
		l,_ := scf(data[2],32)
		h,_ := scf(data[3],32)
		student := students{data[0], t,l, h,tb(t,l,h), }
		sv = append(sv, student)
	}
	fmt.Println(sv)
	sort.SliceStable(sv, func(i int,j int) bool{
		return sv[i].dtb > sv[j].dtb
	})
	fmt.Println(sv)
	xuat(sv)




}
