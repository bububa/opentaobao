package drugtrace

import (
	"sync"
)

// DetailInfoList 结构体
type DetailInfoList struct {
	// 批次编号
	ProduceBatchNo string `json:"produce_batch_no,omitempty" xml:"produce_batch_no,omitempty"`
	// 药品信息
	DrugInfo string `json:"drug_info,omitempty" xml:"drug_info,omitempty"`
	// 追溯码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 验证状态， 1 已验证 ； 其他是未验证
	CodeIsMatched string `json:"code_is_matched,omitempty" xml:"code_is_matched,omitempty"`
}

var poolDetailInfoList = sync.Pool{
	New: func() any {
		return new(DetailInfoList)
	},
}

// GetDetailInfoList() 从对象池中获取DetailInfoList
func GetDetailInfoList() *DetailInfoList {
	return poolDetailInfoList.Get().(*DetailInfoList)
}

// ReleaseDetailInfoList 释放DetailInfoList
func ReleaseDetailInfoList(v *DetailInfoList) {
	v.ProduceBatchNo = ""
	v.DrugInfo = ""
	v.Code = ""
	v.CodeIsMatched = ""
	poolDetailInfoList.Put(v)
}
