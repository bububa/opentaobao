package alihealth2

import (
	"sync"
)

// StoreAuditVo 结构体
type StoreAuditVo struct {
	// 审核状态
	CheckStatusString string `json:"check_status_string,omitempty" xml:"check_status_string,omitempty"`
	// 审核原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 门店ID
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 审核状态码
	CheckStatus int64 `json:"check_status,omitempty" xml:"check_status,omitempty"`
	// 门店审核ID
	StoreAuditId int64 `json:"store_audit_id,omitempty" xml:"store_audit_id,omitempty"`
}

var poolStoreAuditVo = sync.Pool{
	New: func() any {
		return new(StoreAuditVo)
	},
}

// GetStoreAuditVo() 从对象池中获取StoreAuditVo
func GetStoreAuditVo() *StoreAuditVo {
	return poolStoreAuditVo.Get().(*StoreAuditVo)
}

// ReleaseStoreAuditVo 释放StoreAuditVo
func ReleaseStoreAuditVo(v *StoreAuditVo) {
	v.CheckStatusString = ""
	v.Reason = ""
	v.StoreId = 0
	v.CheckStatus = 0
	v.StoreAuditId = 0
	poolStoreAuditVo.Put(v)
}
