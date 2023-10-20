package examination

import (
	"sync"
)

// AlibabaAlihealthExaminationTodoorServiceinfoSyncData 结构体
type AlibabaAlihealthExaminationTodoorServiceinfoSyncData struct {
	// 结果
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 结果
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
}

var poolAlibabaAlihealthExaminationTodoorServiceinfoSyncData = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthExaminationTodoorServiceinfoSyncData)
	},
}

// GetAlibabaAlihealthExaminationTodoorServiceinfoSyncData() 从对象池中获取AlibabaAlihealthExaminationTodoorServiceinfoSyncData
func GetAlibabaAlihealthExaminationTodoorServiceinfoSyncData() *AlibabaAlihealthExaminationTodoorServiceinfoSyncData {
	return poolAlibabaAlihealthExaminationTodoorServiceinfoSyncData.Get().(*AlibabaAlihealthExaminationTodoorServiceinfoSyncData)
}

// ReleaseAlibabaAlihealthExaminationTodoorServiceinfoSyncData 释放AlibabaAlihealthExaminationTodoorServiceinfoSyncData
func ReleaseAlibabaAlihealthExaminationTodoorServiceinfoSyncData(v *AlibabaAlihealthExaminationTodoorServiceinfoSyncData) {
	v.ResponseCode = ""
	v.Msg = ""
	poolAlibabaAlihealthExaminationTodoorServiceinfoSyncData.Put(v)
}
