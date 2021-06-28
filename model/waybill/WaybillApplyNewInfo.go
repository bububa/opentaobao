package waybill

// WaybillApplyNewInfo 
type WaybillApplyNewInfo struct {

    // 根据收货地址返回大头笔信息
    
    ShortAddress   string `json:"short_address,omitempty" xml:"short_address,omitempty"`
    

    // 面单对应的订单列
    
    TradeOrderInfo   *TradeOrderInfo `json:"trade_order_info,omitempty" xml:"trade_order_info,omitempty"`
    

    // 返回的面单号
    
    WaybillCode   string `json:"waybill_code,omitempty" xml:"waybill_code,omitempty"`
    

    // 集包地代码
    
    PackageCenterCode   string `json:"package_center_code,omitempty" xml:"package_center_code,omitempty"`
    

    // 集包地名称
    
    PackageCenterName   string `json:"package_center_name,omitempty" xml:"package_center_name,omitempty"`
    

    // 打印配置项，传给ali-print组件
    
    PrintConfig   string `json:"print_config,omitempty" xml:"print_config,omitempty"`
    

    // 面单号对应的物流服务商网点（分支机构）代码
    
    ShippingBranchCode   string `json:"shipping_branch_code,omitempty" xml:"shipping_branch_code,omitempty"`
    

    // 包裹对应的派件（收件）物流服务商网点（分支机构）名称
    
    ConsigneeBranchName   string `json:"consignee_branch_name,omitempty" xml:"consignee_branch_name,omitempty"`
    

    // 面单号对于的物流服务商网点（分支机构）名称
    
    ShippingBranchName   string `json:"shipping_branch_name,omitempty" xml:"shipping_branch_name,omitempty"`
    

    // 包裹对应的派件（收件）物流服务商网点（分支机构）代码
    
    ConsigneeBranchCode   string `json:"consignee_branch_code,omitempty" xml:"consignee_branch_code,omitempty"`
    

}
