package waybill

// WaybillDetailQueryInfo 
type WaybillDetailQueryInfo struct {

    // 发货网点编码
    ShippingBranchCode   string `json:"shipping_branch_code,omitempty"`

    // 创建时间
    CreateTime   string `json:"create_time,omitempty"`

    // 发件人联系方式
    SendPhone   string `json:"send_phone,omitempty"`

    // 收件人姓名
    ConsigneeName   string `json:"consignee_name,omitempty"`

    // 包裹体积 单位为ML(毫升)或立方厘米
    Weight   int64 `json:"weight,omitempty"`

    // 发件人姓名
    SendName   string `json:"send_name,omitempty"`

    // 面单状态
    Status   int64 `json:"status,omitempty"`

    // 打印次数
    PrintCount   int64 `json:"print_count,omitempty"`

    // 包裹里面的商品类型
    PackageItems   []PackageItem `json:"package_items,omitempty"`

    // 揽收时间
    PickupTime   string `json:"pickup_time,omitempty"`

    // 物流商编码CODE
    CpCode   string `json:"cp_code,omitempty"`

    // 最后一次打印时间
    LastPrintTime   string `json:"last_print_time,omitempty"`

    // 电子面单信息
    WaybillCode   string `json:"waybill_code,omitempty"`

    // 大头笔信息
    ShortAddress   string `json:"short_address,omitempty"`

    // 签收时间
    SignTime   string `json:"sign_time,omitempty"`

    // 使用者ID
    RealUserId   int64 `json:"real_user_id,omitempty"`

    // 包裹重量 单位为G(克)
    Volume   int64 `json:"volume,omitempty"`

    // 发货网点信息
    ShippingBranchName   string `json:"shipping_branch_name,omitempty"`

    // 包裹对应的派件（收件）物流服务商网点（分支机构）代码
    ConsigneeBranchCode   string `json:"consignee_branch_code,omitempty"`

    // 交易订单列表
    TradeOrderList   []String `json:"trade_order_list,omitempty"`

    // 包裹对应的派件（收件）物流服务商网点（分支机构）名称
    ConsigneeBranchName   string `json:"consignee_branch_name,omitempty"`

    // 收件人联系方式
    ConsigneePhone   string `json:"consignee_phone,omitempty"`

    // 收货人地址
    ConsigneeAddress   *WaybillAddress `json:"consignee_address,omitempty"`

    // 发货地址
    ShippingAddress   *WaybillAddress `json:"shipping_address,omitempty"`

    // 物流服务能力集合
    LogisticsServiceList   []LogisticsService `json:"logistics_service_list,omitempty"`

    // 快递服务产品类型编码
    ProductType   string `json:"product_type,omitempty"`

    // ERP订单号或包裹号
    PackageId   string `json:"package_id,omitempty"`

    // 集包地、目的地中心代码。打印时根据该 code 生成目的地中心的条码，条码生成的算法与对应的电子面单条码一致
    PackageCenterCode   string `json:"package_center_code,omitempty"`

    // 集包地、目的地中心名称
    PackageCenterName   string `json:"package_center_name,omitempty"`

    // 打印配置项
    PrintConfig   string `json:"print_config,omitempty"`

}
