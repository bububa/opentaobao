package wms

// CainiaoStockOutBillPackageitemlist 
/* model for simplify = false
type CainiaoStockOutBillPackageitemlist struct {

    // 1
    
    PackageItem  *struct {
        CainiaoStockOutBillPackageitem  *CainiaoStockOutBillPackageitem `json:"cainiao_stock_out_bill_packageitem,omitempty"`
    } `json:"package_item,omitempty"`
    

}
*/

// CainiaoStockOutBillPackageitemlist 
type CainiaoStockOutBillPackageitemlist struct {

    // 1
    PackageItem   *CainiaoStockOutBillPackageitem `json:"package_item,omitempty"`

}
