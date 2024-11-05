package image

import (
	"fmt"
	"os"
	"testing"
)

func TestExtractPngString(t *testing.T) {
	inputPath := "../input.png" // 输入文件路径
	f, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	fmt.Println(ExtractPngString(f))
}
