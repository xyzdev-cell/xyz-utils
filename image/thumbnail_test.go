package image

import (
	"fmt"
	"os"
	"testing"
)

func TestSample(t *testing.T) {
	inputPath := "input.png"      // 输入文件路径
	outputPath := "thumbnail.png" // 输出文件路径
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

	if err := Thumbnail(inputFile, outputFile, "image/png", 0, 200); err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println("缩略图生成成功:", outputPath)
	}
}
