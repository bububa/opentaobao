package logistic

// AlibabaEleFengniaoServicePackageQueryResult 
type AlibabaEleFengniaoServicePackageQueryResult struct {
    // 门店code
    ChainstoreCode   string `json:"chainstore_code,omitempty" xml:"chainstore_code,omitempty"`
    // servicePackageCodes
    ServicePackageCodes   []ServicePackage `json:"service_package_codes,omitempty" xml:"service_package_codes>service_package,omitempty"`
}
