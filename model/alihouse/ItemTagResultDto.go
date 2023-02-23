package alihouse

// ItemTagResultDto 结构体
type ItemTagResultDto struct {
	// 失败的外部ID列表
	FailedOuterIds []string `json:"failed_outer_ids,omitempty" xml:"failed_outer_ids>string,omitempty"`
	// 是否全部成功
	HasSuccess bool `json:"has_success,omitempty" xml:"has_success,omitempty"`
}
