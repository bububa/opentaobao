package xhotelonlineorder

// TopOrderPackageDO 
type TopOrderPackageDO struct {
    // 套餐详情
    OrderPackageDetails   []TopOrderPackageDetail `json:"order_package_details,omitempty" xml:"order_package_details>top_order_package_detail,omitempty"`
    // 描述
    OrderPackageDesc   string `json:"order_package_desc,omitempty" xml:"order_package_desc,omitempty"`
}
