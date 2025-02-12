package tbk

import (
	"sync"
)

// LinkInfoDto 结构体
type LinkInfoDto struct {
	// 二合一长链接
	CouponLongUrl string `json:"coupon_long_url,omitempty" xml:"coupon_long_url,omitempty"`
	// 淘口令原始链接，入参物料不为淘口令时，此字段返回null
	TpwdOriginUrl string `json:"tpwd_origin_url,omitempty" xml:"tpwd_origin_url,omitempty"`
	// 物料ID，需根据material_type判断物料类型， 可能为商品ID、卖家ID、会场ID
	MaterialId string `json:"material_id,omitempty" xml:"material_id,omitempty"`
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
	// 淘积木页面链接，定向开放
	TaoBrickUrl string `json:"tao_brick_url,omitempty" xml:"tao_brick_url,omitempty"`
	// 卖家ID
	SellerId string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 会场ID
	PageId string `json:"page_id,omitempty" xml:"page_id,omitempty"`
	// 入参推广链接中的pid，需配合入参fields使用，如果不属于当前调用的推广者，则不返回
	OriginPid string `json:"origin_pid,omitempty" xml:"origin_pid,omitempty"`
	// 根据入参工具服务商账号信息生成的新商品ID
	IsvMktid string `json:"isv_mktid,omitempty" xml:"isv_mktid,omitempty"`
	// 1—单品 2—店铺 3—会场 4-承接开放 5-优惠券 6-直播间 7-淘积木页
	MaterialType int64 `json:"material_type,omitempty" xml:"material_type,omitempty"`
	// 1-联盟口令，2-主站口令。入参物料不为淘口令时，此字段返回null
	TkBizType int64 `json:"tk_biz_type,omitempty" xml:"tk_biz_type,omitempty"`
}

var poolLinkInfoDto = sync.Pool{
	New: func() any {
		return new(LinkInfoDto)
	},
}

// GetLinkInfoDto() 从对象池中获取LinkInfoDto
func GetLinkInfoDto() *LinkInfoDto {
	return poolLinkInfoDto.Get().(*LinkInfoDto)
}

// ReleaseLinkInfoDto 释放LinkInfoDto
func ReleaseLinkInfoDto(v *LinkInfoDto) {
	v.CouponLongUrl = ""
	v.TpwdOriginUrl = ""
	v.MaterialId = ""
	v.CpsLongUrl = ""
	v.CpsShortTpwd = ""
	v.CouponShortTpwd = ""
	v.CpsShortUrl = ""
	v.CouponShortUrl = ""
	v.CouponFullTpwd = ""
	v.CpsFullTpwd = ""
	v.ShareUnitUrl = ""
	v.TaoBrickUrl = ""
	v.SellerId = ""
	v.PageId = ""
	v.OriginPid = ""
	v.IsvMktid = ""
	v.MaterialType = 0
	v.TkBizType = 0
	poolLinkInfoDto.Put(v)
}
