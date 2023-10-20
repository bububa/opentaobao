package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugCodeKytWesGetlicenseResultModel 结构体
type AlibabaAlihealthDrugCodeKytWesGetlicenseResultModel struct {
	// licenseToken
	LicenseToken string `json:"license_token,omitempty" xml:"license_token,omitempty"`
	// 状态码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 状态值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否响应成功
	ResponseSuccess string `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

var poolAlibabaAlihealthDrugCodeKytWesGetlicenseResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugCodeKytWesGetlicenseResultModel)
	},
}

// GetAlibabaAlihealthDrugCodeKytWesGetlicenseResultModel() 从对象池中获取AlibabaAlihealthDrugCodeKytWesGetlicenseResultModel
func GetAlibabaAlihealthDrugCodeKytWesGetlicenseResultModel() *AlibabaAlihealthDrugCodeKytWesGetlicenseResultModel {
	return poolAlibabaAlihealthDrugCodeKytWesGetlicenseResultModel.Get().(*AlibabaAlihealthDrugCodeKytWesGetlicenseResultModel)
}

// ReleaseAlibabaAlihealthDrugCodeKytWesGetlicenseResultModel 释放AlibabaAlihealthDrugCodeKytWesGetlicenseResultModel
func ReleaseAlibabaAlihealthDrugCodeKytWesGetlicenseResultModel(v *AlibabaAlihealthDrugCodeKytWesGetlicenseResultModel) {
	v.LicenseToken = ""
	v.MsgCode = ""
	v.MsgInfo = ""
	v.ResponseSuccess = ""
	poolAlibabaAlihealthDrugCodeKytWesGetlicenseResultModel.Put(v)
}
