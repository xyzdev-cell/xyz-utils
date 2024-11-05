package image

import (
	"fmt"
	"testing"
)

func TestExtractPngString(t *testing.T) {
	inputPath := "../input.png" // 输入文件路径
	fmt.Println(ExtractPngString(inputPath))
}
