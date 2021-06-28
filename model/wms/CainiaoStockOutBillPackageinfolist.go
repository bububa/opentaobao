package wms

// CainiaoStockOutBillPackageinfolist 
/* model for simplify = false
type CainiaoStockOutBillPackageinfolist struct {

    // 包裹信息
    
    PackageInfo  *struct {
        CainiaoStockOutBillPackageinfo  *CainiaoStockOutBillPackageinfo `json:"cainiao_stock_out_bill_packageinfo,omitempty"`
    } `json:"package_info,omitempty"`
    

}
*/

// CainiaoStockOutBillPackageinfolist 
type CainiaoStockOutBillPackageinfolist struct {

    // 包裹信息
    PackageInfo   *CainiaoStockOutBillPackageinfo `json:"package_info,omitempty"`

}
