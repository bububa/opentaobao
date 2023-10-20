package nropen

import (
	"sync"
)

// AlibabaAscpIndustryDisivisonQueryData 结构体
type AlibabaAscpIndustryDisivisonQueryData struct {
	// 区域信息了列表
	DivisionInfos []Divisioninfos `json:"division_infos,omitempty" xml:"division_infos>divisioninfos,omitempty"`
}

var poolAlibabaAscpIndustryDisivisonQueryData = sync.Pool{
	New: func() any {
		return new(AlibabaAscpIndustryDisivisonQueryData)
	},
}

// GetAlibabaAscpIndustryDisivisonQueryData() 从对象池中获取AlibabaAscpIndustryDisivisonQueryData
func GetAlibabaAscpIndustryDisivisonQueryData() *AlibabaAscpIndustryDisivisonQueryData {
	return poolAlibabaAscpIndustryDisivisonQueryData.Get().(*AlibabaAscpIndustryDisivisonQueryData)
}

// ReleaseAlibabaAscpIndustryDisivisonQueryData 释放AlibabaAscpIndustryDisivisonQueryData
func ReleaseAlibabaAscpIndustryDisivisonQueryData(v *AlibabaAscpIndustryDisivisonQueryData) {
	v.DivisionInfos = v.DivisionInfos[:0]
	poolAlibabaAscpIndustryDisivisonQueryData.Put(v)
}
