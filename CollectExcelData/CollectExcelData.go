package CollectExcelData

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

var sheet_name string = ""
var key string = ""

func CollectExcelData(path string) int{
	f, _ := excelize.OpenFile(path)
    defer func() {
        if err := f.Close(); err != nil {
            fmt.Println(err)
        }
    }()

	// 有効行数取得
	rows, _ := f.GetRows(sheet_name)
	length := len(rows)
	
	cnt := 0
	// 行数分ループ（開始が1、終了がlengthを含む点に注意）
	for i := 1; i<=length; i++{
		cell, _ := excelize.CoordinatesToCellName(18, i) // R列
		value, _ := f.GetCellValue(sheet_name, cell)
		if value == key{
			cnt++
		}
	}
	
	return cnt
}