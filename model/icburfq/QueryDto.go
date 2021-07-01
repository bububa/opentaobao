package icburfq

// QueryDto 结构体
type QueryDto struct {
	// 推荐数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 当前页面数
	Current int64 `json:"current,omitempty" xml:"current,omitempty"`
	// 页面大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 系统参数qn-homepage
	Site string `json:"site,omitempty" xml:"site,omitempty"`
	// 系统参数U_P_I
	Type string `json:"type,omitempty" xml:"type,omitempty"`
}
