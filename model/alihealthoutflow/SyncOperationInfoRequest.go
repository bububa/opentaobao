package alihealthoutflow

import (
	"sync"
)

// SyncOperationInfoRequest 结构体
type SyncOperationInfoRequest struct {
	// 操作时间(非空)
	OperateTime string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	// 授权id(非空)
	AuthorizationId string `json:"authorization_id,omitempty" xml:"authorization_id,omitempty"`
	// 医院id(非空)
	HospitalId string `json:"hospital_id,omitempty" xml:"hospital_id,omitempty"`
	// 渠道(非空)（ALIPAY：支付宝）
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 患者id(非空)
	PatientId string `json:"patient_id,omitempty" xml:"patient_id,omitempty"`
	// 操作类型(非空)（ENTER_HOME_PAGE:进入首页）
	OperateType string `json:"operate_type,omitempty" xml:"operate_type,omitempty"`
	// 患者姓名(非空)
	PatientName string `json:"patient_name,omitempty" xml:"patient_name,omitempty"`
}

var poolSyncOperationInfoRequest = sync.Pool{
	New: func() any {
		return new(SyncOperationInfoRequest)
	},
}

// GetSyncOperationInfoRequest() 从对象池中获取SyncOperationInfoRequest
func GetSyncOperationInfoRequest() *SyncOperationInfoRequest {
	return poolSyncOperationInfoRequest.Get().(*SyncOperationInfoRequest)
}

// ReleaseSyncOperationInfoRequest 释放SyncOperationInfoRequest
func ReleaseSyncOperationInfoRequest(v *SyncOperationInfoRequest) {
	v.OperateTime = ""
	v.AuthorizationId = ""
	v.HospitalId = ""
	v.Channel = ""
	v.PatientId = ""
	v.OperateType = ""
	v.PatientName = ""
	poolSyncOperationInfoRequest.Put(v)
}
