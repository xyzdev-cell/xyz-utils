package utils

import (
	"fmt"
	"os"
	"testing"
)

func TestSample(t *testing.T) {
	inputPath := "input.png"      // 输入文件路径
	outputPath := "thumbnail.png" // 输出文件路径
	width := 100                  // 缩略图宽度
	inputFile, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()
	outputFile, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	if err := Thumbnail(inputFile, outputFile, "image/png", width); err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println("缩略图生成成功:", outputPath)
	}
}
