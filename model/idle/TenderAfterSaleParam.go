package idle

import (
	"sync"
)

// TenderAfterSaleParam 结构体
type TenderAfterSaleParam struct {
	// 履约事件
	EventId string `json:"event_id,omitempty" xml:"event_id,omitempty"`
	// 打款支付宝流水号
	AlipayTradeNo string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// 订单号
	MainOrderId string `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
	// 提示用户需补充的申请信息
	NeedMakeUpNotice string `json:"need_make_up_notice,omitempty" xml:"need_make_up_notice,omitempty"`
	// 售后申请单id
	AfterSaleApplyId string `json:"after_sale_apply_id,omitempty" xml:"after_sale_apply_id,omitempty"`
	// 售后被拒绝，退回物流单号
	SendBackMailNo string `json:"send_back_mail_no,omitempty" xml:"send_back_mail_no,omitempty"`
	// 售后处理方案信息
	ServicePlanInfo *ServicePlanInfo `json:"service_plan_info,omitempty" xml:"service_plan_info,omitempty"`
}

var poolTenderAfterSaleParam = sync.Pool{
	New: func() any {
		return new(TenderAfterSaleParam)
	},
}

// GetTenderAfterSaleParam() 从对象池中获取TenderAfterSaleParam
func GetTenderAfterSaleParam() *TenderAfterSaleParam {
	return poolTenderAfterSaleParam.Get().(*TenderAfterSaleParam)
}

// ReleaseTenderAfterSaleParam 释放TenderAfterSaleParam
func ReleaseTenderAfterSaleParam(v *TenderAfterSaleParam) {
	v.EventId = ""
	v.AlipayTradeNo = ""
	v.MainOrderId = ""
	v.NeedMakeUpNotice = ""
	v.AfterSaleApplyId = ""
	v.SendBackMailNo = ""
	v.ServicePlanInfo = nil
	poolTenderAfterSaleParam.Put(v)
}
