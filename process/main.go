package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	generateHtml()
	//generateJson()
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
