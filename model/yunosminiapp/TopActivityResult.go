package yunosminiapp

// TopActivityResult 结构体
type TopActivityResult struct {
	// 详细信息
	Detail string `json:"detail,omitempty" xml:"detail,omitempty"`
	// 成功与否
	Success string `json:"success,omitempty" xml:"success,omitempty"`
	// 活动状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}
