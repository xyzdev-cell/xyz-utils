package xyz_rand

import (
	"fmt"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
)

func TestRandChoice(t *testing.T) {
	slice1 := []int{1, 2, 3, 4}
	fmt.Println(RandChoice(slice1))
}

func TestRandShuffleSlice(t *testing.T) {
	slice1 := []int{1, 2, 3, 4}
	RandShuffleSlice(slice1)
	fmt.Println(slice1)
}

func BenchmarkRandShuffleSlice(b *testing.B) {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	for i := 0; i < b.N; i++ {
		RandShuffleSlice(slice1)
	}
}

func TestRandSample(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	newSlice := RandSliceSample(slice1, 5)
	fmt.Println(newSlice)
	fmt.Println(slice1)
}

func TestRandomIntRange(t *testing.T) {
	r1 := RandIntRange(3, 3)
	if r1 != 3 {
		t.Errorf("RandomIntRange error:%d", r1)
	}
	r2 := RandIntRange(3, 5)
	if r2 < 3 || r2 > 5 {
		t.Errorf("RandomIntRange error:%d", r2)
	}
	r3 := RandIntRange(5, 3)
	if r3 < 3 || r3 > 5 {
		t.Errorf("RandomIntRange error:%d", r3)
	}
}

type testTypeUser struct {
	Csv []string `toml:"Csv"`
}

func TestSliceRemove(t *testing.T) {
	rawlist := []int{1, 2, 3, 4}
	dellist := []int{2, 4}
	newlist := SliceRemove(rawlist, dellist)
	fmt.Println(newlist)
	assert.Equal(t, 2, len(newlist))
}

// 重置切片的测试
func TestSliceClean(t *testing.T) {
	array := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliceIntA := array[0:6]
	fmt.Printf("array %p\n", &array)
	fmt.Println(array)
	fmt.Printf("sliceIntA %p\n", sliceIntA)
	fmt.Println(sliceIntA)
	fmt.Printf("len of sliceIntA:%d,cap of sliceIntA:%d\n", len(sliceIntA), cap(sliceIntA))
	sliceIntA = sliceIntA[0:0]
	fmt.Printf("array %p\n", &array)
	fmt.Println(array)
	fmt.Printf("sliceIntA %p\n", sliceIntA)
	fmt.Println(sliceIntA)
	fmt.Printf("len of sliceIntA:%d,cap of sliceIntA:%d\n", len(sliceIntA), cap(sliceIntA))
}

// set 测试
func TestSet(t *testing.T) {
	newSet := mapset.NewSet[int]()
	for i := 0; i < 10; i++ {
		newSet.Add(i)
	}
	assert.Equal(t, 10, newSet.Cardinality())
	for j := range newSet.Iterator().C {
		fmt.Print(j)
	}
}

func TestSliceDedup(t *testing.T) {
	rawlist := []int{1, 1, 2, 2, 3, 3, 3, 4}
	newlist := SliceDeduplicate(rawlist)
	fmt.Println(newlist)
	assert.Equal(t, 4, len(newlist))
}

func TestRandStrNum(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandStrNum(i))
	}
}
