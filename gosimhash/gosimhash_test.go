package gosimhash

import (
	"fmt"
	"testing"
)

var (
	r1  uint64
	r2  uint64
	r3  uint64
	r4  uint64
	err error
)

func TestSimhasher_MakeSimHasher(t *testing.T) {
	s1 := "今天是星期一，下午要去上课"
	s2 := "今天是星期日，一整天不用上课"
	s3 := "昨天是星期一，上午不用去上课"
	s4 := ""
	hasher := NewSimhasher()
	r1, err = hasher.MakeSimHasher(s1, 100)
	if err != nil {
		t.Log(err)
	} else {
		fmt.Printf("Test1 succeed:Hash %s result is %v\n", s1, r1)
	}
	r2, err = hasher.MakeSimHasher(s2, 100)
	if err != nil {
		t.Log(err)
	} else {
		fmt.Printf("Test2 succeed:Hash %s result is %v\n", s2, r2)
	}
	r3, err = hasher.MakeSimHasher(s3, 100)
	if err != nil {
		t.Log(err)
	} else {
		fmt.Printf("Test3 succeed:Hash %s result is %v\n", s3, r3)
	}
	r4, err = hasher.MakeSimHasher(s4, 100)
	if err != nil {
		t.Logf("Test4 failed:%v\n", err)
	} else {
		fmt.Printf("Test4 succeed:Hash %s result is %v\n", s4, r4)
	}

}

func TestGetSimilarity(t *testing.T) {
	fmt.Printf("%v %v %v %v\n", r1, r2, r3, r4)
	similarity1, err := GetSimilarity(r1, r2)
	if err != nil {
		t.Logf("similarity1 failed:%v\n", err)
	} else {
		fmt.Printf("Similarity1:%v\n", similarity1)
	}
	similarity2, err := GetSimilarity(r1, r3)
	if err != nil {
		t.Logf("similarity2 failed:%v\n", err)
	} else {
		fmt.Printf("Similarity2:%v\n", similarity2)
	}
	similarity3, err := GetSimilarity(r1, r4)
	if err != nil {
		t.Logf("similarity3 failed:%v\n", err)
	} else {
		fmt.Printf("Similarity3:%v\n", similarity3)
	}
	similarity4, err := GetSimilarity(r2, r3)
	if err != nil {
		t.Logf("similarity4 failed:%v\n", err)
	} else {
		fmt.Printf("Similarity4:%v\n", similarity4)
	}
	similarity5, err := GetSimilarity(r2, r4)
	if err != nil {
		t.Logf("similarity5 failed:%v\n", err)
	} else {
		fmt.Printf("Similarity5:%v\n", similarity5)
	}
	similarity6, err := GetSimilarity(r3, r4)
	if err != nil {
		t.Logf("similarity6 failed:%v\n", err)
	} else {
		fmt.Printf("Similarity6:%v\n", similarity6)
	}

}
