package mos

// BrandCoProductGroupUserQueryParam 结构体
type BrandCoProductGroupUserQueryParam struct {
	// 业务时间
	BusinessTime string `json:"business_time,omitempty" xml:"business_time,omitempty"`
	// 上一页最大id
	PreMaxId int64 `json:"pre_max_id,omitempty" xml:"pre_max_id,omitempty"`
	// 当前页数
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 分页条数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
