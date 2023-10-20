package tbk

import (
	"sync"
)

// WordMapData 结构体
type WordMapData struct {
	// 链接-商品相关关联词落地页地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 商品相关的关联词
	Word string `json:"word,omitempty" xml:"word,omitempty"`
}

var poolWordMapData = sync.Pool{
	New: func() any {
		return new(WordMapData)
	},
}

// GetWordMapData() 从对象池中获取WordMapData
func GetWordMapData() *WordMapData {
	return poolWordMapData.Get().(*WordMapData)
}

// ReleaseWordMapData 释放WordMapData
func ReleaseWordMapData(v *WordMapData) {
	v.Url = ""
	v.Word = ""
	poolWordMapData.Put(v)
}
