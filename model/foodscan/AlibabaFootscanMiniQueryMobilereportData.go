package foodscan

import (
	"sync"
)

// AlibabaFootscanMiniQueryMobilereportData 结构体
type AlibabaFootscanMiniQueryMobilereportData struct {
	// 右脚趾围
	RightZhiwei string `json:"right_zhiwei,omitempty" xml:"right_zhiwei,omitempty"`
	// 右脚跗围
	RightFuwei string `json:"right_fuwei,omitempty" xml:"right_fuwei,omitempty"`
	// 右脚宽
	RightWidth string `json:"right_width,omitempty" xml:"right_width,omitempty"`
	// 右脚长
	RightLength string `json:"right_length,omitempty" xml:"right_length,omitempty"`
	// 左脚趾围
	LeftZhiwei string `json:"left_zhiwei,omitempty" xml:"left_zhiwei,omitempty"`
	// 左脚跗围
	LeftFuwei string `json:"left_fuwei,omitempty" xml:"left_fuwei,omitempty"`
	// 左脚宽
	LeftWidth string `json:"left_width,omitempty" xml:"left_width,omitempty"`
	// 左脚长
	LeftLength string `json:"left_length,omitempty" xml:"left_length,omitempty"`
}

var poolAlibabaFootscanMiniQueryMobilereportData = sync.Pool{
	New: func() any {
		return new(AlibabaFootscanMiniQueryMobilereportData)
	},
}

// GetAlibabaFootscanMiniQueryMobilereportData() 从对象池中获取AlibabaFootscanMiniQueryMobilereportData
func GetAlibabaFootscanMiniQueryMobilereportData() *AlibabaFootscanMiniQueryMobilereportData {
	return poolAlibabaFootscanMiniQueryMobilereportData.Get().(*AlibabaFootscanMiniQueryMobilereportData)
}

// ReleaseAlibabaFootscanMiniQueryMobilereportData 释放AlibabaFootscanMiniQueryMobilereportData
func ReleaseAlibabaFootscanMiniQueryMobilereportData(v *AlibabaFootscanMiniQueryMobilereportData) {
	v.RightZhiwei = ""
	v.RightFuwei = ""
	v.RightWidth = ""
	v.RightLength = ""
	v.LeftZhiwei = ""
	v.LeftFuwei = ""
	v.LeftWidth = ""
	v.LeftLength = ""
	poolAlibabaFootscanMiniQueryMobilereportData.Put(v)
}
