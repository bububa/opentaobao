package tbk

// StepRateDto 结构体
type StepRateDto struct {
	// 前N件佣金结束时间， 当前N件佣金 失效，本字段置空
	Topnendtime int64 `json:"topn_end_time,omitempty" xml:"topn_end_time,omitempty"`
	// 前N件佣金开始时间，当前N件佣金失效，本字段置空
	Topnstarttime int64 `json:"topn_start_time,omitempty" xml:"topn_start_time,omitempty"`
	// 前N件剩余库存，当前N件佣金失效，本字段置空
	Topnquantity int64 `json:"topn_quantity,omitempty" xml:"topn_quantity,omitempty"`
	// 前N件初始总库存，当前N件佣金失效，本字段置空（失效：任务完成、时间结束、商品下架）
	Topntotalcount int64 `json:"topn_total_count,omitempty" xml:"topn_total_count,omitempty"`
}
