package btrip

// OpenEmployeeQueryRequest 结构体
type OpenEmployeeQueryRequest struct {
	// 每页的最大数据记录数量；默认10，该值要求大于0且小于等于1000。
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 游标分页的游标值，用来标记当前分页的开始位置，第一次请求不填表示从头开始遍历；分页查询还有更多数据项时会同时返回下一页起始游标值page_token，遍历下一页时用该page_token获取查询结果。
	PageToken string `json:"page_token,omitempty" xml:"page_token,omitempty"`
	// 第三方企业id。
	ThirdPartCorpId string `json:"third_part_corp_id,omitempty" xml:"third_part_corp_id,omitempty"`
	// 第三方员工工号。
	ThirdPartJobNo string `json:"third_part_job_no,omitempty" xml:"third_part_job_no,omitempty"`
	// 员工更新时间大于等于值，该值格式yyyy-MM-dd HH:mm:ss。
	ModifiedTimeGreaterOrEqualThan string `json:"modified_time_greater_or_equal_than,omitempty" xml:"modified_time_greater_or_equal_than,omitempty"`
}
