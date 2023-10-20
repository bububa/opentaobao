package retail

import (
	"sync"
)

// AlibabaRetailElectronicCertificatePreConfirmResult 结构体
type AlibabaRetailElectronicCertificatePreConfirmResult struct {
	// warningInfos
	WarningInfos []string `json:"warning_infos,omitempty" xml:"warning_infos>string,omitempty"`
	// errorInfos
	ErrorInfos []string `json:"error_infos,omitempty" xml:"error_infos>string,omitempty"`
	// module
	Module *ElectronicCertificateDto `json:"module,omitempty" xml:"module,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaRetailElectronicCertificatePreConfirmResult = sync.Pool{
	New: func() any {
		return new(AlibabaRetailElectronicCertificatePreConfirmResult)
	},
}

// GetAlibabaRetailElectronicCertificatePreConfirmResult() 从对象池中获取AlibabaRetailElectronicCertificatePreConfirmResult
func GetAlibabaRetailElectronicCertificatePreConfirmResult() *AlibabaRetailElectronicCertificatePreConfirmResult {
	return poolAlibabaRetailElectronicCertificatePreConfirmResult.Get().(*AlibabaRetailElectronicCertificatePreConfirmResult)
}

// ReleaseAlibabaRetailElectronicCertificatePreConfirmResult 释放AlibabaRetailElectronicCertificatePreConfirmResult
func ReleaseAlibabaRetailElectronicCertificatePreConfirmResult(v *AlibabaRetailElectronicCertificatePreConfirmResult) {
	v.WarningInfos = v.WarningInfos[:0]
	v.ErrorInfos = v.ErrorInfos[:0]
	v.Module = nil
	v.Success = false
	poolAlibabaRetailElectronicCertificatePreConfirmResult.Put(v)
}
