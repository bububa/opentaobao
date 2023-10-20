package viapi

import (
	"sync"
)

// AliyunViapiGoodstechRecognizeFurnitureAttributeData 结构体
type AliyunViapiGoodstechRecognizeFurnitureAttributeData struct {
	// 预测的风格名称ID
	Predstyle string `json:"predstyle,omitempty" xml:"predstyle,omitempty"`
	// 预测的风格置信概率
	Predprobability string `json:"predprobability,omitempty" xml:"predprobability,omitempty"`
	// 预测的风格中文名称
	Predstyleid string `json:"predstyleid,omitempty" xml:"predstyleid,omitempty"`
}

var poolAliyunViapiGoodstechRecognizeFurnitureAttributeData = sync.Pool{
	New: func() any {
		return new(AliyunViapiGoodstechRecognizeFurnitureAttributeData)
	},
}

// GetAliyunViapiGoodstechRecognizeFurnitureAttributeData() 从对象池中获取AliyunViapiGoodstechRecognizeFurnitureAttributeData
func GetAliyunViapiGoodstechRecognizeFurnitureAttributeData() *AliyunViapiGoodstechRecognizeFurnitureAttributeData {
	return poolAliyunViapiGoodstechRecognizeFurnitureAttributeData.Get().(*AliyunViapiGoodstechRecognizeFurnitureAttributeData)
}

// ReleaseAliyunViapiGoodstechRecognizeFurnitureAttributeData 释放AliyunViapiGoodstechRecognizeFurnitureAttributeData
func ReleaseAliyunViapiGoodstechRecognizeFurnitureAttributeData(v *AliyunViapiGoodstechRecognizeFurnitureAttributeData) {
	v.Predstyle = ""
	v.Predprobability = ""
	v.Predstyleid = ""
	poolAliyunViapiGoodstechRecognizeFurnitureAttributeData.Put(v)
}
