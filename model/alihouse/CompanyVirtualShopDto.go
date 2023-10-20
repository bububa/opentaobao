package alihouse

import (
	"sync"
)

// CompanyVirtualShopDto 结构体
type CompanyVirtualShopDto struct {
	// 外部虚拟店铺ID
	OuterShopId string `json:"outer_shop_id,omitempty" xml:"outer_shop_id,omitempty"`
	// 目标ID
	OuterTargetId string `json:"outer_target_id,omitempty" xml:"outer_target_id,omitempty"`
	// 店铺名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 店铺logo
	Logo string `json:"logo,omitempty" xml:"logo,omitempty"`
	// 背景图
	BackgroundImage string `json:"background_image,omitempty" xml:"background_image,omitempty"`
	// banner图
	BannerImage string `json:"banner_image,omitempty" xml:"banner_image,omitempty"`
	// 虚拟店类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 是否删除 0 - 否 1 - 是
	IsDeleted int64 `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
}

var poolCompanyVirtualShopDto = sync.Pool{
	New: func() any {
		return new(CompanyVirtualShopDto)
	},
}

// GetCompanyVirtualShopDto() 从对象池中获取CompanyVirtualShopDto
func GetCompanyVirtualShopDto() *CompanyVirtualShopDto {
	return poolCompanyVirtualShopDto.Get().(*CompanyVirtualShopDto)
}

// ReleaseCompanyVirtualShopDto 释放CompanyVirtualShopDto
func ReleaseCompanyVirtualShopDto(v *CompanyVirtualShopDto) {
	v.OuterShopId = ""
	v.OuterTargetId = ""
	v.Name = ""
	v.Logo = ""
	v.BackgroundImage = ""
	v.BannerImage = ""
	v.Type = 0
	v.IsDeleted = 0
	poolCompanyVirtualShopDto.Put(v)
}
