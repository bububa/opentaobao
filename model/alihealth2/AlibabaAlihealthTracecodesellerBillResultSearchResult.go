package alihealth2

import (
	"sync"
)

// AlibabaAlihealthTracecodesellerBillResultSearchResult 结构体
type AlibabaAlihealthTracecodesellerBillResultSearchResult struct {
	// 单据编号
	BillCode string `json:"bill_code,omitempty" xml:"bill_code,omitempty"`
	// 单据处理时间
	GmtModified int64 `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 单据处理状态
	TaskStatus int64 `json:"task_status,omitempty" xml:"task_status,omitempty"`
	// 单据上传时间
	GmtCreate int64 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 商家id
	OperEntId int64 `json:"oper_ent_id,omitempty" xml:"oper_ent_id,omitempty"`
}

var poolAlibabaAlihealthTracecodesellerBillResultSearchResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthTracecodesellerBillResultSearchResult)
	},
}

// GetAlibabaAlihealthTracecodesellerBillResultSearchResult() 从对象池中获取AlibabaAlihealthTracecodesellerBillResultSearchResult
func GetAlibabaAlihealthTracecodesellerBillResultSearchResult() *AlibabaAlihealthTracecodesellerBillResultSearchResult {
	return poolAlibabaAlihealthTracecodesellerBillResultSearchResult.Get().(*AlibabaAlihealthTracecodesellerBillResultSearchResult)
}

// ReleaseAlibabaAlihealthTracecodesellerBillResultSearchResult 释放AlibabaAlihealthTracecodesellerBillResultSearchResult
func ReleaseAlibabaAlihealthTracecodesellerBillResultSearchResult(v *AlibabaAlihealthTracecodesellerBillResultSearchResult) {
	v.BillCode = ""
	v.GmtModified = 0
	v.TaskStatus = 0
	v.GmtCreate = 0
	v.OperEntId = 0
	poolAlibabaAlihealthTracecodesellerBillResultSearchResult.Put(v)
}
