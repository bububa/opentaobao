package viapi

import (
	"sync"
)

// AliyunViapiGoodstechRecognizeFurniturespuData 结构体
type AliyunViapiGoodstechRecognizeFurniturespuData struct {
	// 预测的类目中文名称
	Predprobability string `json:"predprobability,omitempty" xml:"predprobability,omitempty"`
	// 预测的类目名称ID
	Predcate string `json:"predcate,omitempty" xml:"predcate,omitempty"`
	// 预测的类目置信概率
	Predcateid string `json:"predcateid,omitempty" xml:"predcateid,omitempty"`
}

var poolAliyunViapiGoodstechRecognizeFurniturespuData = sync.Pool{
	New: func() any {
		return new(AliyunViapiGoodstechRecognizeFurniturespuData)
	},
}

// GetAliyunViapiGoodstechRecognizeFurniturespuData() 从对象池中获取AliyunViapiGoodstechRecognizeFurniturespuData
func GetAliyunViapiGoodstechRecognizeFurniturespuData() *AliyunViapiGoodstechRecognizeFurniturespuData {
	return poolAliyunViapiGoodstechRecognizeFurniturespuData.Get().(*AliyunViapiGoodstechRecognizeFurniturespuData)
}

// ReleaseAliyunViapiGoodstechRecognizeFurniturespuData 释放AliyunViapiGoodstechRecognizeFurniturespuData
func ReleaseAliyunViapiGoodstechRecognizeFurniturespuData(v *AliyunViapiGoodstechRecognizeFurniturespuData) {
	v.Predprobability = ""
	v.Predcate = ""
	v.Predcateid = ""
	poolAliyunViapiGoodstechRecognizeFurniturespuData.Put(v)
}
