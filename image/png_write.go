package image

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// WritePNGChunks writes the PNG file and extracts its chunks.
func CompressPNGWithText(pngFilePath string, width int, height int) error {
	f, err := os.OpenFile(pngFilePath, os.O_RDWR, 0o644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rawChunks, err := ReadPNGChunks(f)
	if err != nil {
		return fmt.Errorf("reading PNG chunks: %w", err)
	}

	var buf bytes.Buffer
	f.Seek(0, io.SeekStart)
	if err := PngCompress(f, &buf, width, height); err != nil {
		return fmt.Errorf("compressing PNG: %w", err)
	}

	data := buf.Bytes()
	chunksWithoutText, err := ReadPNGChunks(bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("reading new PNG chunks: %w", err)
	}
	newChunks := make([]pngChunk, 0, len(chunksWithoutText)+1)
	newChunks = append(newChunks, chunksWithoutText[:len(chunksWithoutText)-1]...)
	for _, chunk := range rawChunks {
		if string(chunk.Type[:]) == "tEXt" {
			newChunks = append(newChunks, chunk)
			newChunks = append(newChunks, chunksWithoutText[len(chunksWithoutText)-1])
			break
		}
	}

	f.Seek(0, io.SeekStart)
	if err = f.Truncate(0); err != nil {
		return fmt.Errorf("truncating file: %w", err)
	}
	if _, err := f.Write([]byte(pngSignature)); err != nil {
		return fmt.Errorf("writing PNG signature: %w", err)
	}
	for _, chunk := range newChunks {
		if err = writeChunk(f, chunk); err != nil {
			return fmt.Errorf("writing chunk: %w", err)
		}
	}
	return nil
}

// writeChunk writes a single PNG chunk to the writer.
func writeChunk(w io.Writer, chunk pngChunk) error {
	if err := binary.Write(w, binary.BigEndian, chunk.Length); err != nil {
		return err
	}
	if _, err := w.Write(chunk.Type[:]); err != nil {
		return err
	}
	if _, err := w.Write(chunk.Data); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, chunk.CRC); err != nil {
		return err
	}
	return nil
}
