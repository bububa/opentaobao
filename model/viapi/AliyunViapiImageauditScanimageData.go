package viapi

import (
	"sync"
)

// AliyunViapiImageauditScanimageData 结构体
type AliyunViapiImageauditScanimageData struct {
	// 图片检测结果
	Results []AliyunViapiImageauditScanimageResult `json:"results,omitempty" xml:"results>aliyun_viapi_imageaudit_scanimage_result,omitempty"`
}

var poolAliyunViapiImageauditScanimageData = sync.Pool{
	New: func() any {
		return new(AliyunViapiImageauditScanimageData)
	},
}

// GetAliyunViapiImageauditScanimageData() 从对象池中获取AliyunViapiImageauditScanimageData
func GetAliyunViapiImageauditScanimageData() *AliyunViapiImageauditScanimageData {
	return poolAliyunViapiImageauditScanimageData.Get().(*AliyunViapiImageauditScanimageData)
}

// ReleaseAliyunViapiImageauditScanimageData 释放AliyunViapiImageauditScanimageData
func ReleaseAliyunViapiImageauditScanimageData(v *AliyunViapiImageauditScanimageData) {
	v.Results = v.Results[:0]
	poolAliyunViapiImageauditScanimageData.Put(v)
}
