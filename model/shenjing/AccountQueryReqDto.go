package shenjing

// AccountQueryReqDto 结构体
type AccountQueryReqDto struct {
	// 当前页码
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 每一页的条数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 账号列表（多个账号必须是同币种的）
	Accounts []string `json:"accounts,omitempty" xml:"accounts>string,omitempty"`
	// 结束日期
	Date2 int64 `json:"date2,omitempty" xml:"date2,omitempty"`
	// 起始日期
	Date1 int64 `json:"date1,omitempty" xml:"date1,omitempty"`
}
