package simba

import (
	"sync"
)

// TaobaoUniversalbpAdzoneFindconfiglistTopResult 结构体
type TaobaoUniversalbpAdzoneFindconfiglistTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	AdzoneConfigVOTopBulkData *TopBulkData `json:"adzone_config_v_o_top_bulk_data,omitempty" xml:"adzone_config_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpAdzoneFindconfiglistTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpAdzoneFindconfiglistTopResult)
	},
}

// GetTaobaoUniversalbpAdzoneFindconfiglistTopResult() 从对象池中获取TaobaoUniversalbpAdzoneFindconfiglistTopResult
func GetTaobaoUniversalbpAdzoneFindconfiglistTopResult() *TaobaoUniversalbpAdzoneFindconfiglistTopResult {
	return poolTaobaoUniversalbpAdzoneFindconfiglistTopResult.Get().(*TaobaoUniversalbpAdzoneFindconfiglistTopResult)
}

// ReleaseTaobaoUniversalbpAdzoneFindconfiglistTopResult 释放TaobaoUniversalbpAdzoneFindconfiglistTopResult
func ReleaseTaobaoUniversalbpAdzoneFindconfiglistTopResult(v *TaobaoUniversalbpAdzoneFindconfiglistTopResult) {
	v.Info = nil
	v.AdzoneConfigVOTopBulkData = nil
	poolTaobaoUniversalbpAdzoneFindconfiglistTopResult.Put(v)
}
