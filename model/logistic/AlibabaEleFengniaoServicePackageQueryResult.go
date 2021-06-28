package logistic

// AlibabaEleFengniaoServicePackageQueryResult 
/* model for simplify = false
type AlibabaEleFengniaoServicePackageQueryResult struct {

    // 门店code
    
    ChainstoreCode   string `json:"chainstore_code,omitempty"`
    

    // servicePackageCodes
    
    ServicePackageCodes  struct {
        ServicePackage  []ServicePackage `json:"service_package,omitempty"`
    } `json:"service_package_codes,omitempty"`
    

}
*/

// AlibabaEleFengniaoServicePackageQueryResult 
type AlibabaEleFengniaoServicePackageQueryResult struct {

    // 门店code
    ChainstoreCode   string `json:"chainstore_code,omitempty"`

    // servicePackageCodes
    ServicePackageCodes   []ServicePackage `json:"service_package_codes,omitempty"`

}
