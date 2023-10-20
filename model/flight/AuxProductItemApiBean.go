package flight

import (
	"sync"
)

// AuxProductItemApiBean 结构体
type AuxProductItemApiBean struct {
	// 产品名称。 最大允许64个字符，不允许*·#|等特殊符号，也不允许带空格换行等符号
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	// 外部ID，此辅营报价的唯一标识，后续用于校验生单；只允许数字字母组合，最大允许32个字符。 不允许包含空格、换行、|这类特殊符号
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 行李说明，当productType=4为必传
	Baggage *BaggageApiBean `json:"baggage,omitempty" xml:"baggage,omitempty"`
	// 柜台价，用于划价显示。 仅允许100的倍数。 币种：人民币。单位：分。
	CounterPrice int64 `json:"counter_price,omitempty" xml:"counter_price,omitempty"`
	// 销售类型 1:一次销售，2:二次销售，3:不限 一次销售是指辅营产品和机票同时购买； 二次销售是指买完机票后，再单独购买辅营
	SaleType int64 `json:"sale_type,omitempty" xml:"sale_type,omitempty"`
	// 辅营产品服务说明,仅当productType=1或productType=2为必传
	Service *ServiceApiBean `json:"service,omitempty" xml:"service,omitempty"`
	// 在线价，是对旅客展示的销售价格。 仅允许100的倍数。 币种：人民币。单位：分。
	OnlinePrice int64 `json:"online_price,omitempty" xml:"online_price,omitempty"`
	// 辅营产品销售规则
	SalesRule *SalesRuleApiBean `json:"sales_rule,omitempty" xml:"sales_rule,omitempty"`
	// 辅营产品退改规则
	RefundRule *AuxRefundApiBean `json:"refund_rule,omitempty" xml:"refund_rule,omitempty"`
	// 成本价。仅允许100的倍数。 币种：人民币。单位：分。
	BasePrice int64 `json:"base_price,omitempty" xml:"base_price,omitempty"`
	// 产品类型 1:贵宾厅，2:CIP，3:在线选座，4:付费行李，6:值机，7:餐食, 8:值机及选座 当前仅允许投放：4，6，7，8
	ProductType int64 `json:"product_type,omitempty" xml:"product_type,omitempty"`
	// 选座说明，当productType=3或8 为必传
	Seat *SeatApiBean `json:"seat,omitempty" xml:"seat,omitempty"`
}

var poolAuxProductItemApiBean = sync.Pool{
	New: func() any {
		return new(AuxProductItemApiBean)
	},
}

// GetAuxProductItemApiBean() 从对象池中获取AuxProductItemApiBean
func GetAuxProductItemApiBean() *AuxProductItemApiBean {
	return poolAuxProductItemApiBean.Get().(*AuxProductItemApiBean)
}

// ReleaseAuxProductItemApiBean 释放AuxProductItemApiBean
func ReleaseAuxProductItemApiBean(v *AuxProductItemApiBean) {
	v.ProductName = ""
	v.OuterId = ""
	v.Baggage = nil
	v.CounterPrice = 0
	v.SaleType = 0
	v.Service = nil
	v.OnlinePrice = 0
	v.SalesRule = nil
	v.RefundRule = nil
	v.BasePrice = 0
	v.ProductType = 0
	v.Seat = nil
	poolAuxProductItemApiBean.Put(v)
}
