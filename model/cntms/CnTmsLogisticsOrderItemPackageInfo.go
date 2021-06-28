package cntms

// CnTmsLogisticsOrderItemPackageInfo 
/* model for simplify = false
type CnTmsLogisticsOrderItemPackageInfo struct {

    // 发货商品信息，最大50条记录
    
    Items  struct {
        CnTmsLogisticsOrderItem  []CnTmsLogisticsOrderItem `json:"cn_tms_logistics_order_item,omitempty"`
    } `json:"items,omitempty"`
    

    // 运单号
    
    MailNo   string `json:"mail_no,omitempty"`
    

    // 包裹重量（克）
    
    PackageWeight   string `json:"package_weight,omitempty"`
    

    // 包裹长度（厘米）
    
    PackageLength   string `json:"package_length,omitempty"`
    

    // 包裹宽度（厘米）
    
    PackageWidth   string `json:"package_width,omitempty"`
    

    // 包裹高度（厘米）
    
    PackageHeight   string `json:"package_height,omitempty"`
    

    // 包裹体积（立方厘米）
    
    PackageVolume   string `json:"package_volume,omitempty"`
    

}
*/

// CnTmsLogisticsOrderItemPackageInfo 
type CnTmsLogisticsOrderItemPackageInfo struct {

    // 发货商品信息，最大50条记录
    Items   []CnTmsLogisticsOrderItem `json:"items,omitempty"`

    // 运单号
    MailNo   string `json:"mail_no,omitempty"`

    // 包裹重量（克）
    PackageWeight   string `json:"package_weight,omitempty"`

    // 包裹长度（厘米）
    PackageLength   string `json:"package_length,omitempty"`

    // 包裹宽度（厘米）
    PackageWidth   string `json:"package_width,omitempty"`

    // 包裹高度（厘米）
    PackageHeight   string `json:"package_height,omitempty"`

    // 包裹体积（立方厘米）
    PackageVolume   string `json:"package_volume,omitempty"`

}
