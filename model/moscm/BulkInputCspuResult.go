package moscm

// BulkInputCspuResult 结构体
type BulkInputCspuResult struct {
	// 录入结果对象
	CspuResultList []InputCspuResult `json:"cspu_result_list,omitempty" xml:"cspu_result_list>input_cspu_result,omitempty"`
}
