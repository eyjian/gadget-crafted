// Wrote by yijian on 2024/08/10
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gadget-crafted/internal/types"
	"io"
	"net/http"
	"os"
	"strings"
)

var (
	tableName = flag.String("table_name", "", "table name")
	delimiter = flag.String("delimiter", "", "delimiter")
	inputFile = flag.String("input_file", "", "input file")
	url       = flag.String("url", "", "url")
)

func main() {
	flag.Parse()
	if !checkParam() {
		flag.Usage()
		os.Exit(1)
	}

	bytes, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Println("read file error:", err)
		os.Exit(2)
	}
	req := types.GenerateCreateTableSqlReq{
		TableName: *tableName,
		Delimiter: *delimiter,
		Text:      string(bytes),
	}
	reqBytes, err := json.Marshal(req)
	if err != nil {
	}

	httpResp, err := http.Post(*url, "application/json", strings.NewReader(string(reqBytes)))
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer httpResp.Body.Close()

	respBodyBytes, err := io.ReadAll(httpResp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println(string(respBodyBytes))
}

func checkParam() bool {
	// 参数值不能为空
	if *tableName == "" || *delimiter == "" || *inputFile == "" || *url == "" {
		return false
	}
	return true
}
