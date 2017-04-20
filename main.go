package main

import(
    "flag"
    "fmt"
    "path/filepath"
    "os"
)

var (
	outpath = flag.String("outpath", "./out", "outpath json path default = ./out")
)

func walkFunc(path string, info os.FileInfo, err error) error {
	//fmt.Printf("%s\n", path)
	fext := filepath.Ext(path)
	//排除打开文件
    fmt.Print("\n"+fext)
	// if strings.Index(path, "~$") != -1 {
	// 	return nil
	// }
	// if fext == ".xlsx" || fext == ".xls" {
	// 	DoFile(path)
	// }
	return nil
}

func main()  {
    flag.Parse()
    fmt.Println(*outpath)
    filepath.Walk("./",walkFunc)
}