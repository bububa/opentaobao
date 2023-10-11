package alihouse

// CategoryControlResultDto 结构体
type CategoryControlResultDto struct {
	// 原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 类目id
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
}
