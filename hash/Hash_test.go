package hash

import (
	"fmt"
	"testing"
)

func TestHaoHasher_Hash64(t *testing.T) {
	s1 := "今天是星期一，下午要去上课"
	s2 := "今天是 期日，一整天不用上课"
	s3 := "昨天是%#期一，上@不用!上课"
	s4 := ""
	hasher := NewHaoHasher()
	r1, err := hasher.Hash64(s1)
	if err != nil {
		t.Log(err)
	} else {
		fmt.Printf("Test1 succeed:Hash %s result is %v\n", s1, r1)
	}
	r2, err := hasher.Hash64(s2)
	if err != nil {
		t.Log(err)
	} else {
		fmt.Printf("Test2 succeed:Hash %s result is %v\n", s2, r2)
	}
	r3, err := hasher.Hash64(s3)
	if err != nil {
		t.Log(err)
	} else {
		fmt.Printf("Test3 succeed:Hash %s result is %v\n", s3, r3)
	}
	r4, err := hasher.Hash64(s4)
	if err != nil {
		t.Logf("Test4 failed:%v\n", err)
	} else {
		fmt.Printf("Test4 succeed:Hash %s result is %v\n", s4, r4)
	}
}
