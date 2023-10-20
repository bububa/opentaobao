package retail

import (
	"sync"
)

// AlibabaRetailElectronicCertificateConfirmResult 结构体
type AlibabaRetailElectronicCertificateConfirmResult struct {
	// warningInfos
	WarningInfos []string `json:"warning_infos,omitempty" xml:"warning_infos>string,omitempty"`
	// errorInfos
	ErrorInfos []string `json:"error_infos,omitempty" xml:"error_infos>string,omitempty"`
	// module
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaRetailElectronicCertificateConfirmResult = sync.Pool{
	New: func() any {
		return new(AlibabaRetailElectronicCertificateConfirmResult)
	},
}

// GetAlibabaRetailElectronicCertificateConfirmResult() 从对象池中获取AlibabaRetailElectronicCertificateConfirmResult
func GetAlibabaRetailElectronicCertificateConfirmResult() *AlibabaRetailElectronicCertificateConfirmResult {
	return poolAlibabaRetailElectronicCertificateConfirmResult.Get().(*AlibabaRetailElectronicCertificateConfirmResult)
}

// ReleaseAlibabaRetailElectronicCertificateConfirmResult 释放AlibabaRetailElectronicCertificateConfirmResult
func ReleaseAlibabaRetailElectronicCertificateConfirmResult(v *AlibabaRetailElectronicCertificateConfirmResult) {
	v.WarningInfos = v.WarningInfos[:0]
	v.ErrorInfos = v.ErrorInfos[:0]
	v.Module = false
	v.Success = false
	poolAlibabaRetailElectronicCertificateConfirmResult.Put(v)
}
