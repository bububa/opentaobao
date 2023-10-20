package ascp

import (
	"sync"
)

// WarehouseResultDetail 结构体
type WarehouseResultDetail struct {
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 服务商自定义的仓编码，服务商+业务身份下唯一，string（50）
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
}

var poolWarehouseResultDetail = sync.Pool{
	New: func() any {
		return new(WarehouseResultDetail)
	},
}

// GetWarehouseResultDetail() 从对象池中获取WarehouseResultDetail
func GetWarehouseResultDetail() *WarehouseResultDetail {
	return poolWarehouseResultDetail.Get().(*WarehouseResultDetail)
}

// ReleaseWarehouseResultDetail 释放WarehouseResultDetail
func ReleaseWarehouseResultDetail(v *WarehouseResultDetail) {
	v.Code = ""
	v.Message = ""
	v.WarehouseCode = ""
	poolWarehouseResultDetail.Put(v)
}
