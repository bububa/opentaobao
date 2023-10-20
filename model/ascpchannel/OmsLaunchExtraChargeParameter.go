package ascpchannel

import (
	"sync"
)

// OmsLaunchExtraChargeParameter 结构体
type OmsLaunchExtraChargeParameter struct {
	// BFC单号
	WdsCoordinationOrderId string `json:"wds_coordination_order_id,omitempty" xml:"wds_coordination_order_id,omitempty"`
	// 描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 扩展字段
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 发起调整金额 单位：分
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 原因类型�: 1为服务费用� 2为二次上门费用� 3为代付费用� 4为配件费用� 5为拆旧费用� 6为维修费用� 7为空跑费用� 8为其它费用�
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolOmsLaunchExtraChargeParameter = sync.Pool{
	New: func() any {
		return new(OmsLaunchExtraChargeParameter)
	},
}

// GetOmsLaunchExtraChargeParameter() 从对象池中获取OmsLaunchExtraChargeParameter
func GetOmsLaunchExtraChargeParameter() *OmsLaunchExtraChargeParameter {
	return poolOmsLaunchExtraChargeParameter.Get().(*OmsLaunchExtraChargeParameter)
}

// ReleaseOmsLaunchExtraChargeParameter 释放OmsLaunchExtraChargeParameter
func ReleaseOmsLaunchExtraChargeParameter(v *OmsLaunchExtraChargeParameter) {
	v.WdsCoordinationOrderId = ""
	v.Desc = ""
	v.Feature = ""
	v.Price = 0
	v.Type = 0
	poolOmsLaunchExtraChargeParameter.Put(v)
}
