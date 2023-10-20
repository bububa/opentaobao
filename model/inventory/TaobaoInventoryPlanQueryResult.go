package inventory

// TaobaoinventoryplanqueryResult 结构体
type TaobaoinventoryplanqueryResult struct {
	// 返回的对象
	Data *PlanInvTopDto `json:"data,omitempty" xml:"data,omitempty"`
	// 返回结果码
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
