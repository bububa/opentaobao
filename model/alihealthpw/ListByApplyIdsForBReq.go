package alihealthpw

// ListByApplyIdsForBReq 结构体
type ListByApplyIdsForBReq struct {
	// 编码列表
	ApplyUniqueCodes []string `json:"apply_unique_codes,omitempty" xml:"apply_unique_codes>string,omitempty"`
}
