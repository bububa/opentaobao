package tmallservice

import (
	"sync"
)

// SupplyConfigRuleDto 结构体
type SupplyConfigRuleDto struct {
	// 工人校验类型（高级工人/品牌品类工人）
	WorkerCheckTypeList []string `json:"worker_check_type_list,omitempty" xml:"worker_check_type_list>string,omitempty"`
	// 规则唯一编号
	UniqueNo string `json:"unique_no,omitempty" xml:"unique_no,omitempty"`
	// servicecode
	ServiceCode string `json:"service_code,omitempty" xml:"service_code,omitempty"`
	// 店铺名称
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 叶子类目�名称
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
}

var poolSupplyConfigRuleDto = sync.Pool{
	New: func() any {
		return new(SupplyConfigRuleDto)
	},
}

// GetSupplyConfigRuleDto() 从对象池中获取SupplyConfigRuleDto
func GetSupplyConfigRuleDto() *SupplyConfigRuleDto {
	return poolSupplyConfigRuleDto.Get().(*SupplyConfigRuleDto)
}

// ReleaseSupplyConfigRuleDto 释放SupplyConfigRuleDto
func ReleaseSupplyConfigRuleDto(v *SupplyConfigRuleDto) {
	v.WorkerCheckTypeList = v.WorkerCheckTypeList[:0]
	v.UniqueNo = ""
	v.ServiceCode = ""
	v.ShopName = ""
	v.CategoryName = ""
	poolSupplyConfigRuleDto.Put(v)
}
