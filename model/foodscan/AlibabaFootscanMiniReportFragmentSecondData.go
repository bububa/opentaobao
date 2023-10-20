package foodscan

import (
	"sync"
)

// AlibabaFootscanMiniReportFragmentSecondData 结构体
type AlibabaFootscanMiniReportFragmentSecondData struct {
	// 扫描ID
	ScanId string `json:"scan_id,omitempty" xml:"scan_id,omitempty"`
	// 测量结果为0
	IsZero bool `json:"is_zero,omitempty" xml:"is_zero,omitempty"`
}

var poolAlibabaFootscanMiniReportFragmentSecondData = sync.Pool{
	New: func() any {
		return new(AlibabaFootscanMiniReportFragmentSecondData)
	},
}

// GetAlibabaFootscanMiniReportFragmentSecondData() 从对象池中获取AlibabaFootscanMiniReportFragmentSecondData
func GetAlibabaFootscanMiniReportFragmentSecondData() *AlibabaFootscanMiniReportFragmentSecondData {
	return poolAlibabaFootscanMiniReportFragmentSecondData.Get().(*AlibabaFootscanMiniReportFragmentSecondData)
}

// ReleaseAlibabaFootscanMiniReportFragmentSecondData 释放AlibabaFootscanMiniReportFragmentSecondData
func ReleaseAlibabaFootscanMiniReportFragmentSecondData(v *AlibabaFootscanMiniReportFragmentSecondData) {
	v.ScanId = ""
	v.IsZero = false
	poolAlibabaFootscanMiniReportFragmentSecondData.Put(v)
}
