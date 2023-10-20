package logistic

import (
	"sync"
)

// AlibabaEleFengniaoServicePackageQueryResult 结构体
type AlibabaEleFengniaoServicePackageQueryResult struct {
	// servicePackageCodes
	ServicePackageCodes []ServicePackage `json:"service_package_codes,omitempty" xml:"service_package_codes>service_package,omitempty"`
	// 门店code
	ChainstoreCode string `json:"chainstore_code,omitempty" xml:"chainstore_code,omitempty"`
}

var poolAlibabaEleFengniaoServicePackageQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaEleFengniaoServicePackageQueryResult)
	},
}

// GetAlibabaEleFengniaoServicePackageQueryResult() 从对象池中获取AlibabaEleFengniaoServicePackageQueryResult
func GetAlibabaEleFengniaoServicePackageQueryResult() *AlibabaEleFengniaoServicePackageQueryResult {
	return poolAlibabaEleFengniaoServicePackageQueryResult.Get().(*AlibabaEleFengniaoServicePackageQueryResult)
}

// ReleaseAlibabaEleFengniaoServicePackageQueryResult 释放AlibabaEleFengniaoServicePackageQueryResult
func ReleaseAlibabaEleFengniaoServicePackageQueryResult(v *AlibabaEleFengniaoServicePackageQueryResult) {
	v.ServicePackageCodes = v.ServicePackageCodes[:0]
	v.ChainstoreCode = ""
	poolAlibabaEleFengniaoServicePackageQueryResult.Put(v)
}
