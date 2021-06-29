package smartstore

// AlipayCheckResult 
type AlipayCheckResult struct {
    // 检查结果描述
    CheckMessage   string `json:"check_message,omitempty" xml:"check_message,omitempty"`
    // 检查CaseID
    CaseId   string `json:"case_id,omitempty" xml:"case_id,omitempty"`
    // 检查结果，0未执行，1测试通过，2不通过，3执行中？
    CheckResult   int64 `json:"check_result,omitempty" xml:"check_result,omitempty"`
}
