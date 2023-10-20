package happytrip

import (
	"sync"
)

// CostDetailInfo 结构体
type CostDetailInfo struct {
	// 费用金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 费用类型名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 费用类型，如果不需要显示直接设为null，为0也会显示出来。discount_fee:折扣金额;refund_price：退款;empty_fee：远途费或超出套餐部分远途费; highway_fee：高速费;  bridge_fee：路桥费; low_speed_fee：低速费或超出套餐低速费;  night_fee：夜间费用或超出套餐部分夜间费; normal_fee：正常行驶费用或超出套餐部分行驶距离费; other_fee：其他费用;  park_fee：停车费; start_price：起步价格;  tip_fee：加价费用;  limit_fee：车最低消费; combo_fee：套餐费用;  normal_time_fee：快车时长费;  cancel_fee：违约费; dynamic_price：动态调价费用;  wait_fee：等候费; clear_fee：清洁费; cross_city_fee：跨城费
	Type string `json:"type,omitempty" xml:"type,omitempty"`
}

var poolCostDetailInfo = sync.Pool{
	New: func() any {
		return new(CostDetailInfo)
	},
}

// GetCostDetailInfo() 从对象池中获取CostDetailInfo
func GetCostDetailInfo() *CostDetailInfo {
	return poolCostDetailInfo.Get().(*CostDetailInfo)
}

// ReleaseCostDetailInfo 释放CostDetailInfo
func ReleaseCostDetailInfo(v *CostDetailInfo) {
	v.Amount = ""
	v.Name = ""
	v.Type = ""
	poolCostDetailInfo.Put(v)
}
