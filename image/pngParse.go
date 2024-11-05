package image

import (
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"hash/crc32"
	"os"
)

// pngChunk represents a chunk in a PNG file.
type pngChunk struct {
	Length uint32
	Type   [4]byte
	Data   []byte
	CRC    uint32
}

// ReadPNGChunks reads the PNG file and extracts its chunks.
func ReadPNGChunks(filePath string) ([]pngChunk, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the PNG signature
	signature := make([]byte, 8)
	if _, err := file.Read(signature); err != nil {
		return nil, err
	}

	if string(signature) != "\x89PNG\r\n\x1a\n" {
		return nil, errors.New("not a valid PNG file")
	}

	var chunks []pngChunk
	for {
		var chunk pngChunk

		// Read chunk length
		if err := binary.Read(file, binary.BigEndian, &chunk.Length); err != nil {
			return nil, err
		}

		// Read chunk type
		if _, err := file.Read(chunk.Type[:]); err != nil {
			return nil, err
		}

		// Read chunk data
		chunk.Data = make([]byte, chunk.Length)
		if _, err := file.Read(chunk.Data); err != nil {
			return nil, err
		}

		// Read chunk CRC
		if err := binary.Read(file, binary.BigEndian, &chunk.CRC); err != nil {
			return nil, err
		}

		if !checkCRC(chunk) {
			return nil, errors.New("invalid CRC")
		}

		chunks = append(chunks, chunk)
		if string(chunk.Type[:]) == "IEND" {
			break
		}
	}

	return chunks, nil
}

func checkCRC(chunk pngChunk) bool {
	// 计算 CRC
	crc := crc32.NewIEEE()
	crc.Write(chunk.Type[:]) // 先写入类型
	crc.Write(chunk.Data)    // 然后写入数据
	return crc.Sum32() == chunk.CRC
}

// extractTextChunks extracts tEXt chunks from the PNG chunks.
func extractTextChunks(chunks []pngChunk) (map[string]string, error) {
	texts := make(map[string]string)

	for _, chunk := range chunks {
		if string(chunk.Type[:]) == "tEXt" {
			// Split the data into keyword and text
			keywordEnd := 0
			for i, b := range chunk.Data {
				if b == 0 {
					keywordEnd = i
					break
				}
			}

			if keywordEnd == 0 {
				continue
			}

			keyword := string(chunk.Data[:keywordEnd])
			text := string(chunk.Data[keywordEnd+1:]) // +1 to skip the null byte
			texts[keyword] = text
		}
	}

	if len(texts) == 0 {
		return nil, errors.New("no text chunks found")
	}

	return texts, nil
}

// base64ToUtf8 converts a base64 string to a UTF-8 string.
func base64ToUtf8(encoded string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ExtractPngString(filePath string) (string, error) {
	chunks, err := ReadPNGChunks(filePath)
	if err != nil {
		return "", fmt.Errorf("reading PNG chunks: %w", err)
	}

	texts, err := extractTextChunks(chunks)
	if err != nil {
		return "", fmt.Errorf("extracting text chunks: %w", err)
	}

	// Check for specific keywords and handle accordingly
	if ccv3, ok := texts["ccv3"]; ok {
		decoded, err := base64ToUtf8(ccv3)
		if err != nil {
			return decoded, nil
		}
	}

	if chara, ok := texts["chara"]; ok {
		decoded, err := base64ToUtf8(chara)
		if err == nil {
			return decoded, nil
		}
	}
	return "", errors.New("ccv3 or chara keyword not found in PNG file")
}
