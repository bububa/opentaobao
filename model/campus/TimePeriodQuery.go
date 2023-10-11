package campus

// TimePeriodQuery 结构体
type TimePeriodQuery struct {
	// 时间规则ID
	TimePeriodIdList []int64 `json:"time_period_id_list,omitempty" xml:"time_period_id_list>int64,omitempty"`
	// 序列号
	SnNo string `json:"sn_no,omitempty" xml:"sn_no,omitempty"`
}
