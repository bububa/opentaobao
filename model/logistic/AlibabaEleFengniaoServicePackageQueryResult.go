package logistic

// AlibabaelefengniaoservicepackagequeryResult 结构体
type AlibabaelefengniaoservicepackagequeryResult struct {
	// servicePackageCodes
	ServicePackageCodes []ServicePackage `json:"service_package_codes,omitempty" xml:"service_package_codes>service_package,omitempty"`
	// 门店code
	ChainstoreCode string `json:"chainstore_code,omitempty" xml:"chainstore_code,omitempty"`
}
