package logistic

import (
	"sync"
)

// ResourceGroupDto 结构体
type ResourceGroupDto struct {
	// 区块编码
	AreaCode string `json:"area_code,omitempty" xml:"area_code,omitempty"`
	// 区块名称
	AreaName string `json:"area_name,omitempty" xml:"area_name,omitempty"`
	// 网格仓外部编码
	FromOrgResourceCode string `json:"from_org_resource_code,omitempty" xml:"from_org_resource_code,omitempty"`
	// from资源来源
	FromOrgSource string `json:"from_org_source,omitempty" xml:"from_org_source,omitempty"`
	// 网格仓编码
	FromResourceCode string `json:"from_resource_code,omitempty" xml:"from_resource_code,omitempty"`
	// from资源名称
	FromResourceName string `json:"from_resource_name,omitempty" xml:"from_resource_name,omitempty"`
	// from资源类型
	FromResourceType string `json:"from_resource_type,omitempty" xml:"from_resource_type,omitempty"`
	// 商家
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 网络编码
	NetworkCode string `json:"network_code,omitempty" xml:"network_code,omitempty"`
	// 自提点地址
	PickupPointAddress string `json:"pickup_point_address,omitempty" xml:"pickup_point_address,omitempty"`
	// 自提点外部编码
	ToOrgResourceCode string `json:"to_org_resource_code,omitempty" xml:"to_org_resource_code,omitempty"`
	// 自提点来源
	ToOrgSource string `json:"to_org_source,omitempty" xml:"to_org_source,omitempty"`
	// 自提点编码
	ToResourceCode string `json:"to_resource_code,omitempty" xml:"to_resource_code,omitempty"`
	// to资源编码（自提点）
	ToResourceName string `json:"to_resource_name,omitempty" xml:"to_resource_name,omitempty"`
	// to资源类型
	ToResourceType string `json:"to_resource_type,omitempty" xml:"to_resource_type,omitempty"`
	// 是否测试
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
}

var poolResourceGroupDto = sync.Pool{
	New: func() any {
		return new(ResourceGroupDto)
	},
}

// GetResourceGroupDto() 从对象池中获取ResourceGroupDto
func GetResourceGroupDto() *ResourceGroupDto {
	return poolResourceGroupDto.Get().(*ResourceGroupDto)
}

// ReleaseResourceGroupDto 释放ResourceGroupDto
func ReleaseResourceGroupDto(v *ResourceGroupDto) {
	v.AreaCode = ""
	v.AreaName = ""
	v.FromOrgResourceCode = ""
	v.FromOrgSource = ""
	v.FromResourceCode = ""
	v.FromResourceName = ""
	v.FromResourceType = ""
	v.MerchantCode = ""
	v.NetworkCode = ""
	v.PickupPointAddress = ""
	v.ToOrgResourceCode = ""
	v.ToOrgSource = ""
	v.ToResourceCode = ""
	v.ToResourceName = ""
	v.ToResourceType = ""
	v.IsTest = 0
	poolResourceGroupDto.Put(v)
}
