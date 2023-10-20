package hotel

import (
	"sync"
)

// GetMixRateListResult 结构体
type GetMixRateListResult struct {
	// 商品评论列表
	MixRates []MixRateVo `json:"mix_rates,omitempty" xml:"mix_rates>mix_rate_vo,omitempty"`
	// 页面布局信息
	ConfigInfo string `json:"config_info,omitempty" xml:"config_info,omitempty"`
	// debugInfo
	DebugInfo string `json:"debug_info,omitempty" xml:"debug_info,omitempty"`
	// errCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 服务器主机名称，排查问题使用
	HostName string `json:"host_name,omitempty" xml:"host_name,omitempty"`
	// 附加属性
	Attributes *Attributes `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 是否可以向下翻页，0不可以，1可以
	HasNextPage int64 `json:"has_next_page,omitempty" xml:"has_next_page,omitempty"`
	// 统计项
	ItemStatistic *ItemStatisticVo `json:"item_statistic,omitempty" xml:"item_statistic,omitempty"`
	// 总记录数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// degrade
	Degrade bool `json:"degrade,omitempty" xml:"degrade,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolGetMixRateListResult = sync.Pool{
	New: func() any {
		return new(GetMixRateListResult)
	},
}

// GetGetMixRateListResult() 从对象池中获取GetMixRateListResult
func GetGetMixRateListResult() *GetMixRateListResult {
	return poolGetMixRateListResult.Get().(*GetMixRateListResult)
}

// ReleaseGetMixRateListResult 释放GetMixRateListResult
func ReleaseGetMixRateListResult(v *GetMixRateListResult) {
	v.MixRates = v.MixRates[:0]
	v.ConfigInfo = ""
	v.DebugInfo = ""
	v.ErrCode = ""
	v.ErrMsg = ""
	v.HostName = ""
	v.Attributes = nil
	v.HasNextPage = 0
	v.ItemStatistic = nil
	v.TotalNum = 0
	v.Degrade = false
	v.Success = false
	poolGetMixRateListResult.Put(v)
}
