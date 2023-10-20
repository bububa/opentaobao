package travel

import (
	"sync"
)

// PontusTravelItemSaleInfo 结构体
type PontusTravelItemSaleInfo struct {
	// 关联商品与店铺类目 结构:&amp;quot;,cid1,cid2,...,&amp;quot;，如果店铺类目存在二级类目，必须传入子类目cids。  支持的最大列表长度为：256； 关于如何获取cid，请参考该接口：http://open.taobao.com/doc2/apiDetail.htm?apiId=65
	SellerCids []string `json:"seller_cids,omitempty" xml:"seller_cids>string,omitempty"`
	// 预约商品必填，普通商品不填。预约商品开始时间，格式：yyyy-MM-dd HH:mm
	BcStartDate string `json:"bc_start_date,omitempty" xml:"bc_start_date,omitempty"`
	// 预约商品必填，普通商品可不填。可选出发结束日期，格式：年-月-日 日期必须是最近300天内的，最早和最晚日期跨度最大为180天。对于普通商品，根据日历库存的最晚时间来自动计算；对于预约商品则为必填字段
	EndComboDate string `json:"end_combo_date,omitempty" xml:"end_combo_date,omitempty"`
	// 电子凭证码商，格式为：码商id:码商nick。电子凭证的卖家规则详见：http://bangpai.taobao.com/group/thread/14051111-283426841.htm
	Merchant string `json:"merchant,omitempty" xml:"merchant,omitempty"`
	// 电子凭证网点ID
	NetworkId string `json:"network_id,omitempty" xml:"network_id,omitempty"`
	// 商品秒杀，商品秒杀三个值：可选类型web_only(只能通过web网络秒杀)，wap_only(只能通过wap网络秒杀)，web_and_wap(既能通过web秒杀也能通过wap秒杀)
	SecondKill string `json:"second_kill,omitempty" xml:"second_kill,omitempty"`
	// 预约商品必填，普通商品可不填。可选出发开始日期，格式：yyyy-MM-dd。对于普通商品，根据日历库存的最早时间来自动计算。对于预约商品则为必填字段
	StartComboDate string `json:"start_combo_date,omitempty" xml:"start_combo_date,omitempty"`
	// 至少提前天数，最晚成团提前天数，最小0天，最大为300天；不传则为0
	Duration int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 商品售卖类型，1. 普通商品  2. 预约商品  3. 拍卖商品  4. 非日历商品；默认为1。
	SaleType int64 `json:"sale_type,omitempty" xml:"sale_type,omitempty"`
	// 扣库存方式。1：拍下减库存，2：付款减库存
	SubStock int64 `json:"sub_stock,omitempty" xml:"sub_stock,omitempty"`
	// 新版电子凭证参数结构
	ItemEleCertInfo *PontusTravelItemEleCertInfo `json:"item_ele_cert_info,omitempty" xml:"item_ele_cert_info,omitempty"`
	// 资源确认方式。1：即时确认，2：二次确认
	ConfirmType int64 `json:"confirm_type,omitempty" xml:"confirm_type,omitempty"`
	// 资源确认时长。1:2个工作小时内确认，2:6个工作小时内确认，3:9个工作小时内确认，4:18个工作小时内确认（只有境外目的地商品支持18小时确认，境内商品最长9小时）
	ConfirmTime int64 `json:"confirm_time,omitempty" xml:"confirm_time,omitempty"`
	// 是否支持会员打折。可选值：true，false；默认值：false(不打折)。不传的话默认为false
	HasDiscount bool `json:"has_discount,omitempty" xml:"has_discount,omitempty"`
	// 已废弃。是否提供发票 默认为false  仅C商家需要设置该值 B商家强制提供发票
	HasInvoice bool `json:"has_invoice,omitempty" xml:"has_invoice,omitempty"`
	// 是否橱窗推荐，可选值：true，false；默认值：false(不推荐)
	HasShowcase bool `json:"has_showcase,omitempty" xml:"has_showcase,omitempty"`
	// 电子凭证是否支持系统自动退款，true则表示支持
	SupportOnsaleAutoRefund bool `json:"support_onsale_auto_refund,omitempty" xml:"support_onsale_auto_refund,omitempty"`
}

var poolPontusTravelItemSaleInfo = sync.Pool{
	New: func() any {
		return new(PontusTravelItemSaleInfo)
	},
}

// GetPontusTravelItemSaleInfo() 从对象池中获取PontusTravelItemSaleInfo
func GetPontusTravelItemSaleInfo() *PontusTravelItemSaleInfo {
	return poolPontusTravelItemSaleInfo.Get().(*PontusTravelItemSaleInfo)
}

// ReleasePontusTravelItemSaleInfo 释放PontusTravelItemSaleInfo
func ReleasePontusTravelItemSaleInfo(v *PontusTravelItemSaleInfo) {
	v.SellerCids = v.SellerCids[:0]
	v.BcStartDate = ""
	v.EndComboDate = ""
	v.Merchant = ""
	v.NetworkId = ""
	v.SecondKill = ""
	v.StartComboDate = ""
	v.Duration = 0
	v.SaleType = 0
	v.SubStock = 0
	v.ItemEleCertInfo = nil
	v.ConfirmType = 0
	v.ConfirmTime = 0
	v.HasDiscount = false
	v.HasInvoice = false
	v.HasShowcase = false
	v.SupportOnsaleAutoRefund = false
	poolPontusTravelItemSaleInfo.Put(v)
}
