package tbk

// SubsidyDetailDto 结构体
type SubsidyDetailDto struct {
	// 该笔订单包含的补贴类型
	Subsidytype string `json:"subsidy_type,omitempty" xml:"subsidy_type,omitempty"`
	// 补贴比率
	Subsidyrate float64 `json:"subsidy_rate,omitempty" xml:"subsidy_rate,omitempty"`
	// 对应补贴类型的补贴金额
	Subsidyfee float64 `json:"subsidy_fee,omitempty" xml:"subsidy_fee,omitempty"`
	// 单笔订单补贴上限。对应补贴类型的单笔订单补贴上限
	Subsidyupperlimit float64 `json:"subsidy_upper_limit,omitempty" xml:"subsidy_upper_limit,omitempty"`
	// 补贴分成比率
	Subsidysharerate float64 `json:"subsidy_share_rate,omitempty" xml:"subsidy_share_rate,omitempty"`
}
