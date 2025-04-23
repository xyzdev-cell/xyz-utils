package image

import (
	"fmt"
	"os"
	"testing"

	xyz_hash "github.com/xyzdev-cell/xyz-utils/hash"
)

func TestExtractPngString(t *testing.T) {
	inputPath := "../output/input.png" // 输入文件路径
	f, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	fmt.Println(ExtractPngString(f))
}

func TestPngCompress(t *testing.T) {
	inputPath2 := "../output/input2.png" // 输入文件路径
	inf2, err := os.Open(inputPath2)
	if err != nil {
		panic(err)
	}
	defer inf2.Close()
	raw_chunks, err := ReadPNGChunks(inf2)
	if err != nil {
		panic(err)
	}

	hash1 := "1"
	for _, chunk := range raw_chunks {
		if string(chunk.Type[:]) == "tEXt" {
			hash1 = xyz_hash.HashString(string(chunk.Data), "sha256")

			break
		}
	}
	// Write the PNG signature
	inputPath1 := "../output/input.png" // 输入文件路径
	err = CompressPNGWithText(inputPath1, 400, 0)
	if err != nil {
		fmt.Println("压缩PNG失败:", err)
		return
	}

	outf, err := os.OpenFile(inputPath1, os.O_RDWR, 0o644)
	if err != nil {
		panic(err)
	}
	defer outf.Close()

	hash2 := "2"
	dist_chunks2, err := ReadPNGChunks(outf)
	if err != nil {
		panic(err)
	}
	for _, chunk := range dist_chunks2 {
		if string(chunk.Type[:]) == "tEXt" {
			hash2 = xyz_hash.HashString(string(chunk.Data), "sha256")
			break
		}
	}
	if hash1 != hash2 {
		fmt.Println("hash不一致")
	} else {
		fmt.Println("hash一致")
	}
}
