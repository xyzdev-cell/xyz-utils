package utils

import (
	"math/rand"
	"time"
)

type anyUInt interface { // 目前没有支持 uint64
	uint | uint8 | uint16 | uint32
}

type anyInt interface {
	int | int8 | int16 | int32 | int64 | anyUInt
}

// var rands = newRandomSource()

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 在any切片中随机一个值
func RandChoice[T any](slice []T) T {
	return slice[rand.Intn(len(slice))]
}

// 随机抽取切片中 n 个不重复元素.
// n 大于切片长度时 panic.
// return 新的切片.
func RandSliceSample[T any](slice []T, num int) []T {
	newSlice := make([]T, num)
	copy(newSlice, slice)
	RandShuffleSlice(newSlice)
	return newSlice[:num]
}

// 原地打乱一个 slice 元素顺序
// 无返回
func RandShuffleSlice[T any](slice []T) {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

// 随机 int, 包括最大数自己
func RandIntRange[T anyInt](num1 T, num2 T) T {
	min, max := int(num1), int(num2)+1
	if num1 == num2 {
		return num1
	} else if num1 > num2 {
		max, min = int(num1)+1, int(num2)
	}
	return T(rand.Intn(max-min) + min)
}

func RandAnyInt[T anyInt](max T) T {
	return T(rand.Intn(int(max)))
}

// 随机源, 单次程序运行使用一次即可
// func newRandomSource() *rand.Rand {
// 	return rand.New(rand.NewSource(time.Now().Unix()))
// }
