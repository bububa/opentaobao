package drugtrace

// AlibabaAlihealthDrugDownloadGetentauthentResult 结构体
type AlibabaAlihealthDrugDownloadGetentauthentResult struct {
	// list
	AuthList []AlibabaAlihealthDrugDownloadGetentauthentModel `json:"auth_list,omitempty" xml:"auth_list>alibaba_alihealth_drug_download_getentauthent_model,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	Issuccess bool `json:"issuccess,omitempty" xml:"issuccess,omitempty"`
}
