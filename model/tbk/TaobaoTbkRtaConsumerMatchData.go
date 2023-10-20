package tbk

// TaobaotbkrtaconsumermatchData 结构体
type TaobaotbkrtaconsumermatchData struct {
	// 返回结果列表
	Resultlist []Resultlist `json:"result_list,omitempty" xml:"result_list>resultlist,omitempty"`
	// 策略ID的匹配结果，仅在入参strategy_id_list字段非空时返回
	Strategyresultlist []StrategyResultList `json:"strategy_result_list,omitempty" xml:"strategy_result_list>strategy_result_list,omitempty"`
}
