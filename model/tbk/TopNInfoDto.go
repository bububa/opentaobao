package tbk

// TopNInfoDto 结构体
type TopNInfoDto struct {
	// 前N件佣金结束时间
	Topnendtime string `json:"topn_end_time,omitempty" xml:"topn_end_time,omitempty"`
	// 前N件佣金开始时间
	Topnstarttime string `json:"topn_start_time,omitempty" xml:"topn_start_time,omitempty"`
	// 前N件佣金率
	Topnrate string `json:"topn_rate,omitempty" xml:"topn_rate,omitempty"`
	// 前N件剩余库存
	Topnquantity int64 `json:"topn_quantity,omitempty" xml:"topn_quantity,omitempty"`
	// 前N件初始总库存
	Topntotalcount int64 `json:"topn_total_count,omitempty" xml:"topn_total_count,omitempty"`
}
