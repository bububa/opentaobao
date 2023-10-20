package mozivds

import (
	"sync"
)

// RemoveTenantAdminsRequest 结构体
type RemoveTenantAdminsRequest struct {
	// 人员Code列表
	EmployeeCodes []string `json:"employee_codes,omitempty" xml:"employee_codes>string,omitempty"`
	// 操作人
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 请求元数据
	RequestMetaData string `json:"request_meta_data,omitempty" xml:"request_meta_data,omitempty"`
	// 租户Id
	TenantId int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
}

var poolRemoveTenantAdminsRequest = sync.Pool{
	New: func() any {
		return new(RemoveTenantAdminsRequest)
	},
}

// GetRemoveTenantAdminsRequest() 从对象池中获取RemoveTenantAdminsRequest
func GetRemoveTenantAdminsRequest() *RemoveTenantAdminsRequest {
	return poolRemoveTenantAdminsRequest.Get().(*RemoveTenantAdminsRequest)
}

// ReleaseRemoveTenantAdminsRequest 释放RemoveTenantAdminsRequest
func ReleaseRemoveTenantAdminsRequest(v *RemoveTenantAdminsRequest) {
	v.EmployeeCodes = v.EmployeeCodes[:0]
	v.Operator = ""
	v.RequestMetaData = ""
	v.TenantId = 0
	poolRemoveTenantAdminsRequest.Put(v)
}
