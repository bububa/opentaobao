package idle

import (
	"sync"
)

// AutoTradeAPIParam 结构体
type AutoTradeAPIParam struct {
	// 场景
	Scenario string `json:"scenario,omitempty" xml:"scenario,omitempty"`
	// AT交易产品标识
	XGlobalBizCode string `json:"x_global_biz_code,omitempty" xml:"x_global_biz_code,omitempty"`
	// 额外参数
	ExtraParam string `json:"extra_param,omitempty" xml:"extra_param,omitempty"`
	// API版本
	Version string `json:"version,omitempty" xml:"version,omitempty"`
	// 订单id
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}

var poolAutoTradeAPIParam = sync.Pool{
	New: func() any {
		return new(AutoTradeAPIParam)
	},
}

// GetAutoTradeAPIParam() 从对象池中获取AutoTradeAPIParam
func GetAutoTradeAPIParam() *AutoTradeAPIParam {
	return poolAutoTradeAPIParam.Get().(*AutoTradeAPIParam)
}

// ReleaseAutoTradeAPIParam 释放AutoTradeAPIParam
func ReleaseAutoTradeAPIParam(v *AutoTradeAPIParam) {
	v.Scenario = ""
	v.XGlobalBizCode = ""
	v.ExtraParam = ""
	v.Version = ""
	v.BizOrderId = 0
	poolAutoTradeAPIParam.Put(v)
}
