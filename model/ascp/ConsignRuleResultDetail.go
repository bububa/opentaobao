package ascp

import (
	"sync"
)

// ConsignRuleResultDetail 结构体
type ConsignRuleResultDetail struct {
	// wms货主id
	WmsOwnerCode string `json:"wms_owner_code,omitempty" xml:"wms_owner_code,omitempty"`
}

var poolConsignRuleResultDetail = sync.Pool{
	New: func() any {
		return new(ConsignRuleResultDetail)
	},
}

// GetConsignRuleResultDetail() 从对象池中获取ConsignRuleResultDetail
func GetConsignRuleResultDetail() *ConsignRuleResultDetail {
	return poolConsignRuleResultDetail.Get().(*ConsignRuleResultDetail)
}

// ReleaseConsignRuleResultDetail 释放ConsignRuleResultDetail
func ReleaseConsignRuleResultDetail(v *ConsignRuleResultDetail) {
	v.WmsOwnerCode = ""
	poolConsignRuleResultDetail.Put(v)
}
