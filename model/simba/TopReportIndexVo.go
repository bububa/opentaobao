package simba

import (
	"sync"
)

// TopReportIndexVo 结构体
type TopReportIndexVo struct {
	// 总花费
	TotalCharge string `json:"total_charge,omitempty" xml:"total_charge,omitempty"`
	// 消费者运营花费
	CrowdSceneCharge string `json:"crowd_scene_charge,omitempty" xml:"crowd_scene_charge,omitempty"`
	// 货品运营花费
	ItemSceneCharge string `json:"item_scene_charge,omitempty" xml:"item_scene_charge,omitempty"`
	// 活动运营花费
	ActivitySceneCharge string `json:"activity_scene_charge,omitempty" xml:"activity_scene_charge,omitempty"`
	// 人群场景花费
	DisplayCharge string `json:"display_charge,omitempty" xml:"display_charge,omitempty"`
	// 关键词花费
	SearchCharge string `json:"search_charge,omitempty" xml:"search_charge,omitempty"`
	// 内容场景花费
	ContentSceneCharge string `json:"content_scene_charge,omitempty" xml:"content_scene_charge,omitempty"`
	// 店铺运营场景花费
	ShopSceneCharge string `json:"shop_scene_charge,omitempty" xml:"shop_scene_charge,omitempty"`
	// 展现量
	AdPv string `json:"ad_pv,omitempty" xml:"ad_pv,omitempty"`
	// 点击量
	Click string `json:"click,omitempty" xml:"click,omitempty"`
	// 花费
	Charge string `json:"charge,omitempty" xml:"charge,omitempty"`
	// 点击率
	Ctr string `json:"ctr,omitempty" xml:"ctr,omitempty"`
	// 平均点击花费
	Ecpc string `json:"ecpc,omitempty" xml:"ecpc,omitempty"`
	// 千次展现花费
	Ecpm string `json:"ecpm,omitempty" xml:"ecpm,omitempty"`
	// 总预售成交金额
	PrepayInshopAmt string `json:"prepay_inshop_amt,omitempty" xml:"prepay_inshop_amt,omitempty"`
	// 总预售成交笔数
	PrepayInshopNum string `json:"prepay_inshop_num,omitempty" xml:"prepay_inshop_num,omitempty"`
	// 直接预售成交金额
	PrepayDirAmt string `json:"prepay_dir_amt,omitempty" xml:"prepay_dir_amt,omitempty"`
	// 直接预售成交笔数
	PrepayDirNum string `json:"prepay_dir_num,omitempty" xml:"prepay_dir_num,omitempty"`
	// 间接预售成交金额
	PrepayIndirAmt string `json:"prepay_indir_amt,omitempty" xml:"prepay_indir_amt,omitempty"`
	// 间接预售成交笔数
	PrepayIndirNum string `json:"prepay_indir_num,omitempty" xml:"prepay_indir_num,omitempty"`
	// 直接成交金额
	AlipayDirAmt string `json:"alipay_dir_amt,omitempty" xml:"alipay_dir_amt,omitempty"`
	// 间接成交金额
	AlipayIndirAmt string `json:"alipay_indir_amt,omitempty" xml:"alipay_indir_amt,omitempty"`
	// 总成交金额
	AlipayInshopAmt string `json:"alipay_inshop_amt,omitempty" xml:"alipay_inshop_amt,omitempty"`
	// 总成交笔数
	AlipayInshopNum string `json:"alipay_inshop_num,omitempty" xml:"alipay_inshop_num,omitempty"`
	// 直接成交笔数
	AlipayDirNum string `json:"alipay_dir_num,omitempty" xml:"alipay_dir_num,omitempty"`
	// 间接成交笔数
	AlipayIndirNum string `json:"alipay_indir_num,omitempty" xml:"alipay_indir_num,omitempty"`
	// 成交转化率
	Cvr string `json:"cvr,omitempty" xml:"cvr,omitempty"`
	// 投入产出比
	Roi string `json:"roi,omitempty" xml:"roi,omitempty"`
	// 总成交成本
	AlipayInshopCost string `json:"alipay_inshop_cost,omitempty" xml:"alipay_inshop_cost,omitempty"`
	// 总购物车数
	CartInshopNum string `json:"cart_inshop_num,omitempty" xml:"cart_inshop_num,omitempty"`
	// 直接购物车数
	CartDirNum string `json:"cart_dir_num,omitempty" xml:"cart_dir_num,omitempty"`
	// 间接购物车数
	CartIndirNum string `json:"cart_indir_num,omitempty" xml:"cart_indir_num,omitempty"`
	// 加购率
	CartRate string `json:"cart_rate,omitempty" xml:"cart_rate,omitempty"`
	// 收藏宝贝数
	ItemColInshopNum string `json:"item_col_inshop_num,omitempty" xml:"item_col_inshop_num,omitempty"`
	// 收藏店铺数
	ShopColDirNum string `json:"shop_col_dir_num,omitempty" xml:"shop_col_dir_num,omitempty"`
	// 店铺收藏成本
	ShopColInshopCost string `json:"shop_col_inshop_cost,omitempty" xml:"shop_col_inshop_cost,omitempty"`
	// 总收藏加购数
	ColCartNum string `json:"col_cart_num,omitempty" xml:"col_cart_num,omitempty"`
	// 总收藏加购成本
	ColCartCost string `json:"col_cart_cost,omitempty" xml:"col_cart_cost,omitempty"`
	// 宝贝收藏加购数
	ItemColCart string `json:"item_col_cart,omitempty" xml:"item_col_cart,omitempty"`
	// 宝贝收藏加购成本
	ItemColCartCost string `json:"item_col_cart_cost,omitempty" xml:"item_col_cart_cost,omitempty"`
	// 总收藏数
	ColNum string `json:"col_num,omitempty" xml:"col_num,omitempty"`
	// 宝贝收藏成本
	ItemColInshopCost string `json:"item_col_inshop_cost,omitempty" xml:"item_col_inshop_cost,omitempty"`
	// 宝贝收藏率
	ItemColInshopRate string `json:"item_col_inshop_rate,omitempty" xml:"item_col_inshop_rate,omitempty"`
	// 加购成本
	CartCost string `json:"cart_cost,omitempty" xml:"cart_cost,omitempty"`
	// 拍下订单笔数
	GmvInshopNum string `json:"gmv_inshop_num,omitempty" xml:"gmv_inshop_num,omitempty"`
	// 拍下订单金额
	GmvInshopAmt string `json:"gmv_inshop_amt,omitempty" xml:"gmv_inshop_amt,omitempty"`
	// 直接收藏宝贝数
	ItemColDirNum string `json:"item_col_dir_num,omitempty" xml:"item_col_dir_num,omitempty"`
	// 间接收藏宝贝数
	ItemColIndirNum string `json:"item_col_indir_num,omitempty" xml:"item_col_indir_num,omitempty"`
	// 优惠券领取量
	CouponShopNum string `json:"coupon_shop_num,omitempty" xml:"coupon_shop_num,omitempty"`
	// 购物金充值笔数
	ShoppingNum string `json:"shopping_num,omitempty" xml:"shopping_num,omitempty"`
	// 购物金充值金额
	ShoppingAmt string `json:"shopping_amt,omitempty" xml:"shopping_amt,omitempty"`
	// 旺旺咨询量
	WwNum string `json:"ww_num,omitempty" xml:"ww_num,omitempty"`
	// 引导访问量
	InshopPv string `json:"inshop_pv,omitempty" xml:"inshop_pv,omitempty"`
	// 引导访问人数
	InshopUv string `json:"inshop_uv,omitempty" xml:"inshop_uv,omitempty"`
	// 引导访问潜客数
	InshopPotentialUv string `json:"inshop_potential_uv,omitempty" xml:"inshop_potential_uv,omitempty"`
	// 引导访问潜客占比
	InshopPotentialUvRate string `json:"inshop_potential_uv_rate,omitempty" xml:"inshop_potential_uv_rate,omitempty"`
	// 入会率
	RhRate string `json:"rh_rate,omitempty" xml:"rh_rate,omitempty"`
	// 入会量
	RhNum string `json:"rh_num,omitempty" xml:"rh_num,omitempty"`
	// 引导访问率
	InshopPvRate string `json:"inshop_pv_rate,omitempty" xml:"inshop_pv_rate,omitempty"`
	// 深度访问量
	DeepInshopPv string `json:"deep_inshop_pv,omitempty" xml:"deep_inshop_pv,omitempty"`
	// 平均访问页面数
	AvgAccessPageNum string `json:"avg_access_page_num,omitempty" xml:"avg_access_page_num,omitempty"`
	// 成交新客数
	NewAlipayInshopUv string `json:"new_alipay_inshop_uv,omitempty" xml:"new_alipay_inshop_uv,omitempty"`
	// 成交新客占比
	NewAlipayInshopUvRate string `json:"new_alipay_inshop_uv_rate,omitempty" xml:"new_alipay_inshop_uv_rate,omitempty"`
	// 会员首购人数
	HySgUv string `json:"hy_sg_uv,omitempty" xml:"hy_sg_uv,omitempty"`
	// 会员成交金额
	HyPayAmt string `json:"hy_pay_amt,omitempty" xml:"hy_pay_amt,omitempty"`
	// 会员成交笔数
	HyPayNum string `json:"hy_pay_num,omitempty" xml:"hy_pay_num,omitempty"`
	// 自然流量转化金额
	NaturalPayAmt string `json:"natural_pay_amt,omitempty" xml:"natural_pay_amt,omitempty"`
	// 自然流量曝光量
	OrgNaturalPv string `json:"org_natural_pv,omitempty" xml:"org_natural_pv,omitempty"`
	// 成交人数
	AlipayInshopUv string `json:"alipay_inshop_uv,omitempty" xml:"alipay_inshop_uv,omitempty"`
	// 人均成交笔数
	AlipayInshopNumAvg string `json:"alipay_inshop_num_avg,omitempty" xml:"alipay_inshop_num_avg,omitempty"`
	// 人均成交金额
	AlipayInshopAmtAvg string `json:"alipay_inshop_amt_avg,omitempty" xml:"alipay_inshop_amt_avg,omitempty"`
}

var poolTopReportIndexVo = sync.Pool{
	New: func() any {
		return new(TopReportIndexVo)
	},
}

// GetTopReportIndexVo() 从对象池中获取TopReportIndexVo
func GetTopReportIndexVo() *TopReportIndexVo {
	return poolTopReportIndexVo.Get().(*TopReportIndexVo)
}

// ReleaseTopReportIndexVo 释放TopReportIndexVo
func ReleaseTopReportIndexVo(v *TopReportIndexVo) {
	v.TotalCharge = ""
	v.CrowdSceneCharge = ""
	v.ItemSceneCharge = ""
	v.ActivitySceneCharge = ""
	v.DisplayCharge = ""
	v.SearchCharge = ""
	v.ContentSceneCharge = ""
	v.ShopSceneCharge = ""
	v.AdPv = ""
	v.Click = ""
	v.Charge = ""
	v.Ctr = ""
	v.Ecpc = ""
	v.Ecpm = ""
	v.PrepayInshopAmt = ""
	v.PrepayInshopNum = ""
	v.PrepayDirAmt = ""
	v.PrepayDirNum = ""
	v.PrepayIndirAmt = ""
	v.PrepayIndirNum = ""
	v.AlipayDirAmt = ""
	v.AlipayIndirAmt = ""
	v.AlipayInshopAmt = ""
	v.AlipayInshopNum = ""
	v.AlipayDirNum = ""
	v.AlipayIndirNum = ""
	v.Cvr = ""
	v.Roi = ""
	v.AlipayInshopCost = ""
	v.CartInshopNum = ""
	v.CartDirNum = ""
	v.CartIndirNum = ""
	v.CartRate = ""
	v.ItemColInshopNum = ""
	v.ShopColDirNum = ""
	v.ShopColInshopCost = ""
	v.ColCartNum = ""
	v.ColCartCost = ""
	v.ItemColCart = ""
	v.ItemColCartCost = ""
	v.ColNum = ""
	v.ItemColInshopCost = ""
	v.ItemColInshopRate = ""
	v.CartCost = ""
	v.GmvInshopNum = ""
	v.GmvInshopAmt = ""
	v.ItemColDirNum = ""
	v.ItemColIndirNum = ""
	v.CouponShopNum = ""
	v.ShoppingNum = ""
	v.ShoppingAmt = ""
	v.WwNum = ""
	v.InshopPv = ""
	v.InshopUv = ""
	v.InshopPotentialUv = ""
	v.InshopPotentialUvRate = ""
	v.RhRate = ""
	v.RhNum = ""
	v.InshopPvRate = ""
	v.DeepInshopPv = ""
	v.AvgAccessPageNum = ""
	v.NewAlipayInshopUv = ""
	v.NewAlipayInshopUvRate = ""
	v.HySgUv = ""
	v.HyPayAmt = ""
	v.HyPayNum = ""
	v.NaturalPayAmt = ""
	v.OrgNaturalPv = ""
	v.AlipayInshopUv = ""
	v.AlipayInshopNumAvg = ""
	v.AlipayInshopAmtAvg = ""
	poolTopReportIndexVo.Put(v)
}
