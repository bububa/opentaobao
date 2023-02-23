package alihouse

// StoreBailDto 结构体
type StoreBailDto struct {
	// 外部门店id
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 测试标
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
	// 版本号
	EtcVersion int64 `json:"etc_version,omitempty" xml:"etc_version,omitempty"`
	// 保证金余额
	Bail int64 `json:"bail,omitempty" xml:"bail,omitempty"`
	// 行业类型
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
}
