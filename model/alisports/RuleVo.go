package alisports

// RuleVo 结构体
type RuleVo struct {
	// 补助方式
	SubsidyMethod int64 `json:"subsidy_method,omitempty" xml:"subsidy_method,omitempty"`
	// 补助上限
	SubsidyUpper int64 `json:"subsidy_upper,omitempty" xml:"subsidy_upper,omitempty"`
	// 每人每次补助限制类型
	OnceSubsidyLimitType int64 `json:"once_subsidy_limit_type,omitempty" xml:"once_subsidy_limit_type,omitempty"`
	// 每人每次补助上限
	OnceSubsidyUpper int64 `json:"once_subsidy_upper,omitempty" xml:"once_subsidy_upper,omitempty"`
	// 公司补助限制类型
	WholeSubsidyLimitType int64 `json:"whole_subsidy_limit_type,omitempty" xml:"whole_subsidy_limit_type,omitempty"`
	// 公司补助上限
	WholeSubsidyUpper int64 `json:"whole_subsidy_upper,omitempty" xml:"whole_subsidy_upper,omitempty"`
	// 每人每月补助限制类型
	SubsidyLimitType int64 `json:"subsidy_limit_type,omitempty" xml:"subsidy_limit_type,omitempty"`
}
