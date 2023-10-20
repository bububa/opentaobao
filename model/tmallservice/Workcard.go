package tmallservice

import (
	"sync"
)

// Workcard 结构体
type Workcard struct {
	// 签到时间
	GmtSignIn string `json:"gmt_sign_in,omitempty" xml:"gmt_sign_in,omitempty"`
	// 工单状态编码
	StatusCode string `json:"status_code,omitempty" xml:"status_code,omitempty"`
	// 核销时间
	GmtIdentify string `json:"gmt_identify,omitempty" xml:"gmt_identify,omitempty"`
	// 预约时间开始
	GmtReserveStart string `json:"gmt_reserve_start,omitempty" xml:"gmt_reserve_start,omitempty"`
	// 买家期望服务时间
	BuyerExpectDate string `json:"buyer_expect_date,omitempty" xml:"buyer_expect_date,omitempty"`
	// 分配工人时间
	GmtAssignWorker string `json:"gmt_assign_worker,omitempty" xml:"gmt_assign_worker,omitempty"`
	// 预约结束时间
	GmtReserveEnd string `json:"gmt_reserve_end,omitempty" xml:"gmt_reserve_end,omitempty"`
	// 工单创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 额外属性，包含：{   servPrice: 服务采购价,   newService: 切服务化之后的服务单标识,   lbxNo: 淘宝的物流订单号,   discountFee: 优惠金额,   reserveSource: 预约来源,   serviceDate: 预约时间,   gmtExpire: 过期时间,   bizCode: 业务身份,   型号: IKEA00000237S,   serviceAuctionTitle: 服务商品标题,   reserveTimeEnd: 预约时间段的结束时间,   reassign: 是否是改派工单,   masterSkuId: 实物sku id,   serviceSkuDesc: 服务sku描述,   operator: 工单预约操作者 1表示消费者主动预约，2表示服务回传预约时间,   settlement: 是否进行线上结算,   mainServiceOrder: 双主服务主订单key,   signedTime: 签收时间,   masterAuctionTitle: 实物商品标题,   msfRervDate: 瞄师傅预约日期,   类目: 书架,   品牌: IKEA/宜家,   auctionImageUrl: 实物图片链接,   masterSkuDesc: 实物sku描述,   af: 交易订单实付金额,   auto_dispatch_order: 是否智能派单,   usedServiceCount: &#34;工单总使用次数&#34;,   reserveTimeStart: 预约时间段的开始时间,   masterParentOrderId: 实物订单主订单,   address_city: 买家城市,   reserveRemark: 预约备注,   父类目: 架类,   arrivalTimeNum: 物品签收时间,   outerIdSKU: 商家编码 }
	AttributeMap string `json:"attribute_map,omitempty" xml:"attribute_map,omitempty"`
	// 已使用的次数
	UsedServiceCount int64 `json:"used_service_count,omitempty" xml:"used_service_count,omitempty"`
	// 服务订单数据
	ServiceTradeOrder *ServiceTradeOrder `json:"service_trade_order,omitempty" xml:"service_trade_order,omitempty"`
	// 工单服务总次数
	ServiceCount int64 `json:"service_count,omitempty" xml:"service_count,omitempty"`
	// 工单id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 剩余次数
	LeftServiceCount int64 `json:"left_service_count,omitempty" xml:"left_service_count,omitempty"`
	// 主订单号
	ParentBizOrderId int64 `json:"parent_biz_order_id,omitempty" xml:"parent_biz_order_id,omitempty"`
	// 服务提供者
	ServiceProvider *ServiceProvider `json:"service_provider,omitempty" xml:"service_provider,omitempty"`
	// 实物订单信息
	MasterTradeOrder *MasterTradeOrder `json:"master_trade_order,omitempty" xml:"master_trade_order,omitempty"`
	// 服务定义
	ServiceDefinition *ServiceDefinition `json:"service_definition,omitempty" xml:"service_definition,omitempty"`
	// 核销/工单完结请求中直接带回该字段
	Sequence int64 `json:"sequence,omitempty" xml:"sequence,omitempty"`
	// 服务单id
	SpServiceOrderId int64 `json:"sp_service_order_id,omitempty" xml:"sp_service_order_id,omitempty"`
}

var poolWorkcard = sync.Pool{
	New: func() any {
		return new(Workcard)
	},
}

// GetWorkcard() 从对象池中获取Workcard
func GetWorkcard() *Workcard {
	return poolWorkcard.Get().(*Workcard)
}

// ReleaseWorkcard 释放Workcard
func ReleaseWorkcard(v *Workcard) {
	v.GmtSignIn = ""
	v.StatusCode = ""
	v.GmtIdentify = ""
	v.GmtReserveStart = ""
	v.BuyerExpectDate = ""
	v.GmtAssignWorker = ""
	v.GmtReserveEnd = ""
	v.GmtCreate = ""
	v.AttributeMap = ""
	v.UsedServiceCount = 0
	v.ServiceTradeOrder = nil
	v.ServiceCount = 0
	v.Id = 0
	v.LeftServiceCount = 0
	v.ParentBizOrderId = 0
	v.ServiceProvider = nil
	v.MasterTradeOrder = nil
	v.ServiceDefinition = nil
	v.Sequence = 0
	v.SpServiceOrderId = 0
	poolWorkcard.Put(v)
}
