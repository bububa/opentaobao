package drugtrace

import (
	"sync"
)

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

var poolAlibabaAlihealthDrugDownloadGetentauthentResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugDownloadGetentauthentResult)
	},
}

// GetAlibabaAlihealthDrugDownloadGetentauthentResult() 从对象池中获取AlibabaAlihealthDrugDownloadGetentauthentResult
func GetAlibabaAlihealthDrugDownloadGetentauthentResult() *AlibabaAlihealthDrugDownloadGetentauthentResult {
	return poolAlibabaAlihealthDrugDownloadGetentauthentResult.Get().(*AlibabaAlihealthDrugDownloadGetentauthentResult)
}

// ReleaseAlibabaAlihealthDrugDownloadGetentauthentResult 释放AlibabaAlihealthDrugDownloadGetentauthentResult
func ReleaseAlibabaAlihealthDrugDownloadGetentauthentResult(v *AlibabaAlihealthDrugDownloadGetentauthentResult) {
	v.AuthList = v.AuthList[:0]
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Issuccess = false
	poolAlibabaAlihealthDrugDownloadGetentauthentResult.Put(v)
}
