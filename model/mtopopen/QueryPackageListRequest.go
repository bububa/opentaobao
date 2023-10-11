package mtopopen

// QueryPackageListRequest 结构体
type QueryPackageListRequest struct {
	// 快递公司编码
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 用户的唯一标识
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 页数，从1开始
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
