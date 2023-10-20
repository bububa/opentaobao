package viapi

import (
	"sync"
)

// AliyunViapiOcrCharacterData 结构体
type AliyunViapiOcrCharacterData struct {
	// 返回识别信息
	Results []AliyunViapiOcrCharacterResult `json:"results,omitempty" xml:"results>aliyun_viapi_ocr_character_result,omitempty"`
}

var poolAliyunViapiOcrCharacterData = sync.Pool{
	New: func() any {
		return new(AliyunViapiOcrCharacterData)
	},
}

// GetAliyunViapiOcrCharacterData() 从对象池中获取AliyunViapiOcrCharacterData
func GetAliyunViapiOcrCharacterData() *AliyunViapiOcrCharacterData {
	return poolAliyunViapiOcrCharacterData.Get().(*AliyunViapiOcrCharacterData)
}

// ReleaseAliyunViapiOcrCharacterData 释放AliyunViapiOcrCharacterData
func ReleaseAliyunViapiOcrCharacterData(v *AliyunViapiOcrCharacterData) {
	v.Results = v.Results[:0]
	poolAliyunViapiOcrCharacterData.Put(v)
}
