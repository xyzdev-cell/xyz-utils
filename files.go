package utils

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

)


func BackupFile(filename string, newName string) error {
	s := AbsPath(filename)
	fileinfo, err := os.Stat(s)
	if err != nil {
		// 文件不存在 无需改名
		return nil
	}
	if !fileinfo.IsDir() {
		os.Rename(s, AbsPath(newName))
	} else {
		return fmt.Errorf("该路径是个目录: %s", filename)
	}
	return nil
}

// 写入文件, 注意文件已经存在会覆盖
func WriteFile(filename string, rawContent []byte) error {
	fb, err := os.Create(AbsPath(filename))
	if err != nil {
		return err
	}
	writeLen, err := fb.Write(rawContent)
	if err != nil {
		return err
	} else if writeLen != len(rawContent) {
		return fmt.Errorf("write file %s: len err, expect %d but: %d", filename, len(rawContent), writeLen)
	}
	return nil
}

/*
RelativePath 应该以 **main.go** 所在文件夹为基准
return 全路径.
*/
func AbsPath(RelativePath string) (allfilepath string) {
	return filepath.Join(baseDir, RelativePath)
}

var baseDir = GetBaseDir() // main.go 所在文件夹

/*
为了统一 test mode 和 普通 mode 的不同基准目录.
下面这个总是获取当前文件的路径
但是编译后使用的是编译时的*绝对路径*, 所以只能用于 test mode

	runtime.Caller(0)
*/
func GetBaseDir() string {
	baseDir, err := os.Getwd()
	if err != nil {
		panic("os.Getw 获取当前路径出错")
	}
	if isRunInTestMode() {
		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			panic("runtime.Caller(0) 获取当前路径出错")
		}
		baseDir = filepath.Join(path.Dir(filename), "../")
	}
	return baseDir
}

/*
// 1.8 版本之后失效, 可能test的参数附加不再走flag.

	flag.Lookup("test.v")

// 取决于运行主体, main 和 test 不同.

	os.Getwd()

// 编译之后的exe路径, 如果直接运行go, 多半会在temp 目录下, 不通用.
// 但是 test 运行路径结尾一定是 .test.exe .
// 考虑操作系统差异, 不通用 .

	os.Args[0]

// 测试状态将会带有以下参数

	os.Args

	C:\Users\user\AppData\Local\Temp\go-build1687529845\b001\vr_client.test.exe
	-test.testlogfile=C:\Users\user\AppData\Local\Temp\go-build1687529845\b001\testlog.txt
	-test.paniconexit0
	-test.timeout=30s
	-test.v=true // 可能没有, 取决于设置
	-test.count=1  // 可能没有, 取决于设置
	-test.run=^TestIsQueueOK$
*/
func isRunInTestMode() bool {
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "-test.") {
			return true
		}
	}
	return false
}
