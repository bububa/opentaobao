package alsc

import (
	"sync"
)

// RechargeAccountFlowDetailOpenInfo 结构体
type RechargeAccountFlowDetailOpenInfo struct {
	// 操作储值资产信息列表
	AffectedPropertyList []RechargeAccountPropertyOpenInfo `json:"affected_property_list,omitempty" xml:"affected_property_list>recharge_account_property_open_info,omitempty"`
	// 操作后储值资产信息列表
	AfterPropertyList []RechargeAccountPropertyOpenInfo `json:"after_property_list,omitempty" xml:"after_property_list>recharge_account_property_open_info,omitempty"`
	// 操作前储值资产信息列表
	BeforePropertyList []RechargeAccountPropertyOpenInfo `json:"before_property_list,omitempty" xml:"before_property_list>recharge_account_property_open_info,omitempty"`
	// 储值账户id
	AccountId string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	// 品牌Id
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 储值账户流水id
	FlowId string `json:"flow_id,omitempty" xml:"flow_id,omitempty"`
	// 储值相关的交易类型
	FlowType string `json:"flow_type,omitempty" xml:"flow_type,omitempty"`
	// 操作人ID
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 操作人
	OperatorName string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
	// 交易时间
	OrderTime string `json:"order_time,omitempty" xml:"order_time,omitempty"`
	// 外部交易单号id
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
	// 理由
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 交易门店ID
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 交易门店名称
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 卡ID
	CardId string `json:"card_id,omitempty" xml:"card_id,omitempty"`
	// 原支付单号
	OuterPayId string `json:"outer_pay_id,omitempty" xml:"outer_pay_id,omitempty"`
	// 扩展信息
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 交易后剩余总金额
	CurrentValue int64 `json:"current_value,omitempty" xml:"current_value,omitempty"`
	// 外部订单来源
	OrderSrc int64 `json:"order_src,omitempty" xml:"order_src,omitempty"`
	// 交易总金额，增加为正数，减少为负数
	OrderValue int64 `json:"order_value,omitempty" xml:"order_value,omitempty"`
}

var poolRechargeAccountFlowDetailOpenInfo = sync.Pool{
	New: func() any {
		return new(RechargeAccountFlowDetailOpenInfo)
	},
}

// GetRechargeAccountFlowDetailOpenInfo() 从对象池中获取RechargeAccountFlowDetailOpenInfo
func GetRechargeAccountFlowDetailOpenInfo() *RechargeAccountFlowDetailOpenInfo {
	return poolRechargeAccountFlowDetailOpenInfo.Get().(*RechargeAccountFlowDetailOpenInfo)
}

// ReleaseRechargeAccountFlowDetailOpenInfo 释放RechargeAccountFlowDetailOpenInfo
func ReleaseRechargeAccountFlowDetailOpenInfo(v *RechargeAccountFlowDetailOpenInfo) {
	v.AffectedPropertyList = v.AffectedPropertyList[:0]
	v.AfterPropertyList = v.AfterPropertyList[:0]
	v.BeforePropertyList = v.BeforePropertyList[:0]
	v.AccountId = ""
	v.BrandId = ""
	v.FlowId = ""
	v.FlowType = ""
	v.Operator = ""
	v.OperatorName = ""
	v.OrderTime = ""
	v.OuterOrderId = ""
	v.Remark = ""
	v.ShopId = ""
	v.ShopName = ""
	v.CardId = ""
	v.OuterPayId = ""
	v.ExtInfo = ""
	v.CurrentValue = 0
	v.OrderSrc = 0
	v.OrderValue = 0
	poolRechargeAccountFlowDetailOpenInfo.Put(v)
}
