package cainiaolocker

import (
	"sync"
)

// TradeOrderInfoDto 结构体
type TradeOrderInfoDto struct {
	// 物流服务值（详见https://support-cnkuaidi.taobao.com/doc.htm#?docId=106156&amp;docType=1，如无特殊服务请置空）
	LogisticsServices string `json:"logistics_services,omitempty" xml:"logistics_services,omitempty"`
	// 带面单号模式取号，目前仅顺丰支持
	WaybillCode string `json:"waybill_code,omitempty" xml:"waybill_code,omitempty"`
	// 云打印标准模板URL（组装云打印结果使用，值格式http://cloudprint.cainiao.com/template/standard/${模板ID}）
	TemplateUrl string `json:"template_url,omitempty" xml:"template_url,omitempty"`
	// &lt;a href=&#34;http://open.taobao.com/docs/doc.htm?docType=1&amp;articleId=105086&amp;treeId=17&amp;platformId=17#6&#34;&gt;请求ID&lt;/a&gt;
	ObjectId string `json:"object_id,omitempty" xml:"object_id,omitempty"`
	// 收件人信息
	Recipient *UserInfoDto `json:"recipient,omitempty" xml:"recipient,omitempty"`
	// 订单信息
	OrderInfo *OrderInfoDto `json:"order_info,omitempty" xml:"order_info,omitempty"`
	// 包裹信息
	PackageInfo *PackageInfoDto `json:"package_info,omitempty" xml:"package_info,omitempty"`
	// 使用者ID（使用电子面单账号的实际商家ID，如存在一个电子面单账号多个店铺使用时，请传入店铺的商家ID）
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

var poolTradeOrderInfoDto = sync.Pool{
	New: func() any {
		return new(TradeOrderInfoDto)
	},
}

// GetTradeOrderInfoDto() 从对象池中获取TradeOrderInfoDto
func GetTradeOrderInfoDto() *TradeOrderInfoDto {
	return poolTradeOrderInfoDto.Get().(*TradeOrderInfoDto)
}

// ReleaseTradeOrderInfoDto 释放TradeOrderInfoDto
func ReleaseTradeOrderInfoDto(v *TradeOrderInfoDto) {
	v.LogisticsServices = ""
	v.WaybillCode = ""
	v.TemplateUrl = ""
	v.ObjectId = ""
	v.Recipient = nil
	v.OrderInfo = nil
	v.PackageInfo = nil
	v.UserId = 0
	poolTradeOrderInfoDto.Put(v)
}
