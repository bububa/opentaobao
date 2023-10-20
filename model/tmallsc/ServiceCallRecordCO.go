package tmallsc

import (
	"sync"
)

// ServiceCallRecordCO 结构体
type ServiceCallRecordCO struct {
	// 修改日期
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 录音文件链接
	RecordLink string `json:"record_link,omitempty" xml:"record_link,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 呼叫记录的call_id
	CallId string `json:"call_id,omitempty" xml:"call_id,omitempty"`
	// 主叫拨打时间
	CallTime string `json:"call_time,omitempty" xml:"call_time,omitempty"`
	// 唯一标识
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolServiceCallRecordCO = sync.Pool{
	New: func() any {
		return new(ServiceCallRecordCO)
	},
}

// GetServiceCallRecordCO() 从对象池中获取ServiceCallRecordCO
func GetServiceCallRecordCO() *ServiceCallRecordCO {
	return poolServiceCallRecordCO.Get().(*ServiceCallRecordCO)
}

// ReleaseServiceCallRecordCO 释放ServiceCallRecordCO
func ReleaseServiceCallRecordCO(v *ServiceCallRecordCO) {
	v.GmtModified = ""
	v.RecordLink = ""
	v.GmtCreate = ""
	v.CallId = ""
	v.CallTime = ""
	v.Id = 0
	poolServiceCallRecordCO.Put(v)
}
