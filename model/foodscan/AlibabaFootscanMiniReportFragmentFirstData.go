package foodscan

import (
	"sync"
)

// AlibabaFootscanMiniReportFragmentFirstData 结构体
type AlibabaFootscanMiniReportFragmentFirstData struct {
	// 扫描ID
	ScanId string `json:"scan_id,omitempty" xml:"scan_id,omitempty"`
	// 测量结果为0
	IsZero bool `json:"is_zero,omitempty" xml:"is_zero,omitempty"`
}

var poolAlibabaFootscanMiniReportFragmentFirstData = sync.Pool{
	New: func() any {
		return new(AlibabaFootscanMiniReportFragmentFirstData)
	},
}

// GetAlibabaFootscanMiniReportFragmentFirstData() 从对象池中获取AlibabaFootscanMiniReportFragmentFirstData
func GetAlibabaFootscanMiniReportFragmentFirstData() *AlibabaFootscanMiniReportFragmentFirstData {
	return poolAlibabaFootscanMiniReportFragmentFirstData.Get().(*AlibabaFootscanMiniReportFragmentFirstData)
}

// ReleaseAlibabaFootscanMiniReportFragmentFirstData 释放AlibabaFootscanMiniReportFragmentFirstData
func ReleaseAlibabaFootscanMiniReportFragmentFirstData(v *AlibabaFootscanMiniReportFragmentFirstData) {
	v.ScanId = ""
	v.IsZero = false
	poolAlibabaFootscanMiniReportFragmentFirstData.Put(v)
}
