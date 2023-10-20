package tbk

import (
	"sync"
)

// TaobaoTbkRtaConsumerMatchData 结构体
type TaobaoTbkRtaConsumerMatchData struct {
	// 返回结果列表
	ResultList []Resultlist `json:"result_list,omitempty" xml:"result_list>resultlist,omitempty"`
	// 策略ID的匹配结果，仅在入参strategy_id_list字段非空时返回
	StrategyResultList []StrategyResultList `json:"strategy_result_list,omitempty" xml:"strategy_result_list>strategy_result_list,omitempty"`
}

var poolTaobaoTbkRtaConsumerMatchData = sync.Pool{
	New: func() any {
		return new(TaobaoTbkRtaConsumerMatchData)
	},
}

// GetTaobaoTbkRtaConsumerMatchData() 从对象池中获取TaobaoTbkRtaConsumerMatchData
func GetTaobaoTbkRtaConsumerMatchData() *TaobaoTbkRtaConsumerMatchData {
	return poolTaobaoTbkRtaConsumerMatchData.Get().(*TaobaoTbkRtaConsumerMatchData)
}

// ReleaseTaobaoTbkRtaConsumerMatchData 释放TaobaoTbkRtaConsumerMatchData
func ReleaseTaobaoTbkRtaConsumerMatchData(v *TaobaoTbkRtaConsumerMatchData) {
	v.ResultList = v.ResultList[:0]
	v.StrategyResultList = v.StrategyResultList[:0]
	poolTaobaoTbkRtaConsumerMatchData.Put(v)
}
