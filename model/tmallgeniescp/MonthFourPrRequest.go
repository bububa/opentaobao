package tmallgeniescp

// MonthFourPrRequest 
type MonthFourPrRequest struct {
    // 请求参数
    MonthFourPrParamDTOS   []MonthFourPrParamDTO `json:"month_four_pr_param_d_t_o_s,omitempty" xml:"month_four_pr_param_d_t_o_s>month_four_pr_param_dto,omitempty"`
    // 扩展参数
    RequestExtendJson   string `json:"request_extend_json,omitempty" xml:"request_extend_json,omitempty"`
}
