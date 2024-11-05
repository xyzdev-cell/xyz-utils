package file

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"os"
	"strings"

	toml "github.com/pelletier/go-toml/v2"
)

func LoadCsvFile(fileName string) (rows [][]string, err error) {
	// 获取数据，按照文件
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// 读取文件数据
	rows, err = csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}
	return
}

func LoadTomlFile(fileName string, structPointer any) error {
	fb, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer fb.Close()
	err = toml.NewDecoder(fb).Decode(structPointer)
	if err != nil {
		return err
	}
	return nil
}

// 直接覆盖文件,
// 如果要编辑, 先load完整文件内容
func WriteTomlFile(fileName string, structPointer any) error {
	fb, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer fb.Close()
	err = toml.NewEncoder(fb).Encode(structPointer)
	if err != nil {
		return err
	}
	return nil
}

func LoadXmlFile(fileName string, structPointer any) error {
	fb, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer fb.Close()
	err = xml.NewDecoder(fb).Decode(structPointer)
	if err != nil {
		return err
	}
	return nil
}

func BackupFile(fileName string, newName string) error {
	fileinfo, err := os.Stat(fileName)
	if err != nil {
		// 文件不存在 无需改名
		return nil
	}
	if !fileinfo.IsDir() {
		os.Rename(fileName, newName)
	} else {
		return fmt.Errorf("该路径是个目录: %s", fileName)
	}
	return nil
}

// 写入文件, 注意文件已经存在会覆盖
func WriteFile(fileName string, rawContent []byte) error {
	fb, err := os.Create(fileName)
	if err != nil {
		return err
	}
	writeLen, err := fb.Write(rawContent)
	if err != nil {
		return err
	} else if writeLen != len(rawContent) {
		return fmt.Errorf("write file %s: len err, expect %d but: %d", fileName, len(rawContent), writeLen)
	}
	return nil
}

/*
为了统一 test mode 和 普通 mode 的不同基准目录.
下面这个总是获取当前文件的路径
但是编译后使用的是编译时的*绝对路径*, 所以只能用于 test mode

	runtime.Caller(0)
*/
// func GetBaseDir() string {
// 	baseDir, err := os.Getwd()
// 	if err != nil {
// 		panic("os.Getw 获取当前路径出错")
// 	}
// 	if IsRunInTestMode() {
// 		_, filename, _, ok := runtime.Caller(0)
// 		if !ok {
// 			panic("runtime.Caller(0) 获取当前路径出错")
// 		}
// 		baseDir = filepath.Join(path.Dir(filename), "../")
// 	}
// 	return baseDir
// }

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
func IsRunInTestMode() bool {
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "-test.") {
			return true
		}
	}
	return false
}
