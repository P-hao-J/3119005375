package hash

import "errors"

type HaoHasher struct {
}

func NewHaoHasher() *HaoHasher {
	return &HaoHasher{}
}

func (hasher *HaoHasher) Hash64(data string) (uint64, error) {
	return computeHash(data)
}

//MD5的hash方法
func computeHash(data string) (uint64, error) {
	bytes := []byte(data)
	var err error
	if len(bytes) == 0 {
		err = errors.New("输入文本数据为空")
		return 0, err
	}
	var a, b, c uint64
	a, b = 0x9e3779b9, 0x9e3779b9
	c = 0
	i := 0

	for i = 0; i < len(bytes)-12; {
		a += uint64(bytes[i]) | uint64(bytes[i+1]<<8) | uint64(bytes[i+2]<<16) | uint64(bytes[i+3]<<24)
		i += 4
		b += uint64(bytes[i]) | uint64(bytes[i+1]<<8) | uint64(bytes[i+2]<<16) | uint64(bytes[i+3]<<24)
		i += 4
		c += uint64(bytes[i]) | uint64(bytes[i+1]<<8) | uint64(bytes[i+2]<<16) | uint64(bytes[i+3]<<24)

		a, b, c = mix(a, b, c)
	}
	c += uint64(len(bytes))
	param := [3]uint64{a, b, c}
	for index, _ := range param {
		for j := 0; j < 4; j++ {
			if i < len(bytes) {
				param[index] += uint64(bytes[i]) << (8 * j)
				i++
			}
		}
	}
	a, b, c = mix(param[0], param[1], param[2])
	return c, err
}
func mix(a, b, c uint64) (uint64, uint64, uint64) {
	a -= b
	a -= c
	a ^= c >> 13
	b -= c
	b -= a
	b ^= a << 8
	c -= a
	c -= b
	c ^= b >> 13
	a -= b
	a -= c
	a ^= c >> 12
	b -= c
	b -= a
	b ^= a << 16
	c -= a
	c -= b
	c ^= b >> 5
	a -= b
	a -= c
	a ^= c >> 3
	b -= c
	b -= a
	b ^= a << 10
	c -= a
	c -= b
	c ^= b >> 15
	return a, b, c
}
