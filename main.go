package main

import (
	"os"
	"path/filepath"
	"strings"
	"fmt"
	"ExcelDataCollect/CollectExcelData"
)

// 再起探索対象ディレクトリ
var dirs[3] string = [3]string{
	"aaa",
	"bbb",
	"ccc",
} 

// 探索対象Excelファイル
var targetExcel string = ""	

// 総数
var total int = 0

func main(){
	args := os.Args
	rootPath := args[1]
	
	for _, dir := range dirs {
		fmt.Printf("===== %s ==================\n", dir)
		// 再起探索(第2引数はファイル or ディレクトリの情報を受け取るコールバック関数)
		err := filepath.Walk(rootPath+"/"+dir, JudgeFilename)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("===== 総数 ==================")
	fmt.Println(total)
}

func JudgeFilename(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	filename := info.Name()
	if strings.Contains(filename, targetExcel){
		fmt.Println(path)
		cnt := CollectExcelData.CollectExcelData(path)
		fmt.Println(cnt)
		total += cnt
	}
	return nil
}