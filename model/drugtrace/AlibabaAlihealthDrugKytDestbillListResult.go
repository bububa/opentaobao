package drugtrace

// AlibabaAlihealthDrugKytDestbillListResult 结构体
type AlibabaAlihealthDrugKytDestbillListResult struct {
	// 接口返回model
	ModelList []AlibabaAlihealthDrugKytDestbillListModel `json:"model_list,omitempty" xml:"model_list>alibaba_alihealth_drug_kyt_destbill_list_model,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}
