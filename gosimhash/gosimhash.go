package gosimhash

import (
	"PapeCheck/utils"
	"errors"
	jieba "github.com/yanyiwu/gojieba" //分词的包
)

type Simhasher struct {
	extractor *jieba.Jieba
	hasher    utils.Hasher
}

type HashWeight struct {
	hash   uint64
	weight float64
}

func NewSimhasher() *Simhasher {
	newHasher := utils.NewHaoHasher()
	var (
		dict      string
		hmm       string
		userDict  string
		stopwords string
		idf       string
	)
	getDicPath(&dict, &hmm, &userDict, &idf, &stopwords)
	return &Simhasher{
		extractor: jieba.NewJieba(dict, hmm, userDict, idf, stopwords),
		hasher:    newHasher,
	}
}

func (simhasher *Simhasher) MakeSimHasher(data string, topk int) (uint64, error) {
	fws := simhasher.extractor.ExtractWithWeight(data, topk)
	var err error
	if len(fws) == 0 {
		err = errors.New("输入文本数据为空,无数据可提取")
		return 0, err
	}
	hws := simhasher.ConvertFeatureToHash(fws)
	var one uint64 = 1
	var vector [64]float64
	for _, hw := range hws {
		for i := 0; i < 64; i++ {
			if (one << uint(i) & hw.hash) > 0 {
				vector[i] += hw.weight
			} else {
				vector[i] -= hw.weight
			}
		}
	}
	var res uint64 = 0
	for i, val := range vector {
		if val > 0.0 {
			res |= one << uint(i)
		}
	}
	return res, err
}

func (simhasher *Simhasher) ConvertFeatureToHash(fws []jieba.WordWeight) []HashWeight {
	size := len(fws)
	hws := make([]HashWeight, size, size)
	for index, fw := range fws {
		hws[index].hash, _ = simhasher.hasher.Hash64(fw.Word)
		hws[index].weight = fw.Weight
	}
	return hws
}

//输出文件
func getDicPath(dict, hmm, userDict, idf, stopWords *string) {
	if *dict == "" {
		*dict = jieba.DICT_PATH
	}
	if *hmm == "" {
		*hmm = jieba.HMM_PATH
	}
	if *userDict == "" {
		*userDict = jieba.USER_DICT_PATH
	}
	if *idf == "" {
		*idf = jieba.IDF_PATH
	}
	if *stopWords == "" {
		*stopWords = jieba.STOP_WORDS_PATH
	}
}
func GetHammingDis(data1 uint64, data2 uint64) int {
	xor := data1 ^ data2
	distance := 0
	for xor != 0 {
		xor &= xor - 1
		distance++
	}
	return distance
}

func GetSimilarity(data1 uint64, data2 uint64) (float64, error) {
	if data1 == 0 || data2 == 0 {
		err := errors.New("存在空串，无法比对")
		return 0, err
	}
	distance := GetHammingDis(data1, data2)
	return 0.01 * (100 - float64(distance)*100/128), nil
}
