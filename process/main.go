package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/russross/blackfriday/v2"
	"html/template"
	"os"
	"strconv"
)

var customDomain = "https://emoji.sky123.top"
var customFile = "onetwo"
var customName = "一二"
var size = 332
var page = 5

func main() {
	//generateHtml()
	//generateJson()
	transferMdToHtml()
}

func transferMdToHtml() {
	// 输入文件路径
	inputFile := "../README.md"
	outputFile := "../index.html"
	// 读取 README.md 文件内容
	mdContent, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", inputFile, err)
		return
	}
	// 使用 blackfriday 将 Markdown 转为 HTML
	htmlContent := blackfriday.Run(mdContent)
	// 将 HTML 写入 index.html 文件
	err = os.WriteFile(outputFile, htmlContent, 0644)
	if err != nil {
		fmt.Printf("Error writing file %s: %v\n", outputFile, err)
		return
	}
	fmt.Printf("Successfully converted %s to %s\n", inputFile, outputFile)
}

func generateJson() {
	pageSize, divSize, count := size/page, size%page, 1
	emojiList := make([]EmojiInclude, 0)
	for i := 1; i <= page; i++ {
		emojiInclude := EmojiInclude{Name: customName + "第" + strconv.Itoa(i) + "弹",
			Type: "image", Items: make([]map[string]string, 0)}
		// 开始处理
		for j := 1; j <= pageSize; j++ {
			emojiInclude.Items = append(emojiInclude.Items, map[string]string{
				"key": customFile + "-" + strconv.Itoa(count),
				"val": customDomain + "/" + customFile + "/data/" + strconv.Itoa(count) + ".gif",
			})
			count++
		}
		if i == page {
			for j := 1; j <= divSize; j++ {
				emojiInclude.Items = append(emojiInclude.Items, map[string]string{
					"key": customFile + "-" + strconv.Itoa(count),
					"val": customDomain + "/" + customFile + "/data/" + strconv.Itoa(count) + ".gif",
				})
				count++
			}
		}
		fileName := "../" + customFile + "/" + customFile + "-" + strconv.Itoa(i) + ".json"
		f, _ := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		// 格式化 JSON 数据
		prettyJSON, _ := json.MarshalIndent(emojiInclude, "", "	")
		_, _ = f.Write(prettyJSON)
		_ = f.Close()
		emojiList = append(emojiList, emojiInclude)
	}
	fileName := "../" + customFile + "/include.json"
	f, _ := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	// 格式化 JSON 数据
	prettyJSON, _ := json.MarshalIndent(emojiList, "", "	")
	_, _ = f.Write(prettyJSON)
	_ = f.Close()
}

func generateHtml() {
	col, currDir, count := 10, "./"+customFile+"/data/", 1
	var row int
	if size%col == 0 {
		row = size / col
	} else {
		row = size/col + 1
	}
	rowsList := make([]EmojiRow, 0)
	for i := 1; i <= row; i++ {
		emojiRow := EmojiRow{
			Emojis: make([]Emoji, 0),
		}
		for j := 1; j <= col; j++ {
			emojiRow.Emojis = append(emojiRow.Emojis, Emoji{
				Path:  currDir + strconv.Itoa(count) + ".gif",
				Width: 120,
				Alt:   customFile + "-" + strconv.Itoa(count),
			})
			count++
			if count > size {
				break
			}
		}
		rowsList = append(rowsList, emojiRow)
	}
	tmpl, _ := template.ParseFiles("template.html")
	var buf bytes.Buffer
	_ = tmpl.Execute(&buf, rowsList)
	fmt.Println(buf.String())
}
