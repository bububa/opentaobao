package traderate

import (
	"sync"
)

// GetMixRateListResult 结构体
type GetMixRateListResult struct {
	// 评价明细信息
	MixRates []MixRateVo `json:"mix_rates,omitempty" xml:"mix_rates>mix_rate_vo,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否下一页
	HasNextPage int64 `json:"has_next_page,omitempty" xml:"has_next_page,omitempty"`
	// 评价统计信息
	ItemStatistic *ItemStatisticVo `json:"item_statistic,omitempty" xml:"item_statistic,omitempty"`
	// 总数量
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 是否成功
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
	v.ErrMsg = ""
	v.HasNextPage = 0
	v.ItemStatistic = nil
	v.TotalNum = 0
	v.Success = false
	poolGetMixRateListResult.Put(v)
}
