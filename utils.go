package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func PrettyJSON(jsonByte []byte) string {
	prettyJSON := &bytes.Buffer{}
	json.Indent(prettyJSON, jsonByte, "", "    ")
	return prettyJSON.String()
}

func StartCPUProfile(file string, duration time.Duration) (err error) {
	f, err := os.Create(file)
	if err != nil {
		return err
	}

	log.Println("Start cpu profiling for", duration)
	err = pprof.StartCPUProfile(f)
	if err != nil {
		f.Close()
		return err
	}

	time.AfterFunc(duration, func() {
		pprof.StopCPUProfile()
		f.Close()
		log.Println("Stop CPU profiling after", duration)
	})
	return nil
}

func StartMemoryProfile(file string, duration time.Duration) (err error) {
	f, err := os.Create(file)
	if err != nil {
		return err
	}

	log.Println("Start memory profiling for", duration)
	time.AfterFunc(duration, func() {
		err = pprof.WriteHeapProfile(f)
		if err != nil {
			log.Println(err)
		}
		f.Close()
		log.Println("Stop memory profiling after", duration)
	})
	return nil
}
