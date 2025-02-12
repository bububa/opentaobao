package tbk

import (
	"sync"
)

// ItemLinkInfoDto 结构体
type ItemLinkInfoDto struct {
	// 二合一长链接
	CouponLongUrl string `json:"coupon_long_url,omitempty" xml:"coupon_long_url,omitempty"`
	// 商品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// CPS链接长链
	CpsLongUrl string `json:"cps_long_url,omitempty" xml:"cps_long_url,omitempty"`
	// CPS链接的短口令
	CpsShortTpwd string `json:"cps_short_tpwd,omitempty" xml:"cps_short_tpwd,omitempty"`
	// 二合一链接的短口令
	CouponShortTpwd string `json:"coupon_short_tpwd,omitempty" xml:"coupon_short_tpwd,omitempty"`
	// CPS链接短链
	CpsShortUrl string `json:"cps_short_url,omitempty" xml:"cps_short_url,omitempty"`
	// 二合一链接长链
	CouponShortUrl string `json:"coupon_short_url,omitempty" xml:"coupon_short_url,omitempty"`
	// 二合一链接的长口令
	CouponFullTpwd string `json:"coupon_full_tpwd,omitempty" xml:"coupon_full_tpwd,omitempty"`
	// CPS链接的长口令
	CpsFullTpwd string `json:"cps_full_tpwd,omitempty" xml:"cps_full_tpwd,omitempty"`
	// 种草页链接，定向开放
	ShareUnitUrl string `json:"share_unit_url,omitempty" xml:"share_unit_url,omitempty"`
	// 根据入参工具服务商账号信息生成的新商品ID
	IsvMktid string `json:"isv_mktid,omitempty" xml:"isv_mktid,omitempty"`
	// 1—单品 2—店铺 3—会场
	MaterialType int64 `json:"material_type,omitempty" xml:"material_type,omitempty"`
}

var poolItemLinkInfoDto = sync.Pool{
	New: func() any {
		return new(ItemLinkInfoDto)
	},
}

// GetItemLinkInfoDto() 从对象池中获取ItemLinkInfoDto
func GetItemLinkInfoDto() *ItemLinkInfoDto {
	return poolItemLinkInfoDto.Get().(*ItemLinkInfoDto)
}

// ReleaseItemLinkInfoDto 释放ItemLinkInfoDto
func ReleaseItemLinkInfoDto(v *ItemLinkInfoDto) {
	v.CouponLongUrl = ""
	v.ItemId = ""
	v.CpsLongUrl = ""
	v.CpsShortTpwd = ""
	v.CouponShortTpwd = ""
	v.CpsShortUrl = ""
	v.CouponShortUrl = ""
	v.CouponFullTpwd = ""
	v.CpsFullTpwd = ""
	v.ShareUnitUrl = ""
	v.IsvMktid = ""
	v.MaterialType = 0
	poolItemLinkInfoDto.Put(v)
}
