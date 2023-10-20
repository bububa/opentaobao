package xhotelonlineorder

import (
	"sync"
)

// TopOrderPackageDetail 结构体
type TopOrderPackageDetail struct {
	// 景区地址
	ScenicAddress string `json:"scenic_address,omitempty" xml:"scenic_address,omitempty"`
	// 景区名称
	ScenicName string `json:"scenic_name,omitempty" xml:"scenic_name,omitempty"`
	// 图文详情
	HowToPlay string `json:"how_to_play,omitempty" xml:"how_to_play,omitempty"`
	// 套餐数量单位
	AmountUnit string `json:"amount_unit,omitempty" xml:"amount_unit,omitempty"`
	// 附加产品使用维度   1:每间维度 2:每晚维度 3：入住人数维度
	DimensionText string `json:"dimension_text,omitempty" xml:"dimension_text,omitempty"`
	// 套餐类型名称
	TypeName string `json:"type_name,omitempty" xml:"type_name,omitempty"`
	// 套餐名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 景点封面图
	ScenicCoverImg string `json:"scenic_cover_img,omitempty" xml:"scenic_cover_img,omitempty"`
	// 景区id
	ScenicId int64 `json:"scenic_id,omitempty" xml:"scenic_id,omitempty"`
	// 套餐数量
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 套餐优惠金额
	DiscountPrice int64 `json:"discount_price,omitempty" xml:"discount_price,omitempty"`
	// 套餐类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 套餐ID
	PackageId int64 `json:"package_id,omitempty" xml:"package_id,omitempty"`
}

var poolTopOrderPackageDetail = sync.Pool{
	New: func() any {
		return new(TopOrderPackageDetail)
	},
}

// GetTopOrderPackageDetail() 从对象池中获取TopOrderPackageDetail
func GetTopOrderPackageDetail() *TopOrderPackageDetail {
	return poolTopOrderPackageDetail.Get().(*TopOrderPackageDetail)
}

// ReleaseTopOrderPackageDetail 释放TopOrderPackageDetail
func ReleaseTopOrderPackageDetail(v *TopOrderPackageDetail) {
	v.ScenicAddress = ""
	v.ScenicName = ""
	v.HowToPlay = ""
	v.AmountUnit = ""
	v.DimensionText = ""
	v.TypeName = ""
	v.Name = ""
	v.ScenicCoverImg = ""
	v.ScenicId = 0
	v.Amount = 0
	v.DiscountPrice = 0
	v.Type = 0
	v.PackageId = 0
	poolTopOrderPackageDetail.Put(v)
}
