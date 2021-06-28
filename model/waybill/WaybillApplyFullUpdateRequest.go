package waybill

// WaybillApplyFullUpdateRequest 
/* model for simplify = false
type WaybillApplyFullUpdateRequest struct {

    // 发件人联系方式
    
    SendPhone   string `json:"send_phone,omitempty"`
    

    // 收件人姓名
    
    ConsigneeName   string `json:"consignee_name,omitempty"`
    

    // 包裹重量 单位为G(克)
    
    Weight   int64 `json:"weight,omitempty"`
    

    // 发件人姓名
    
    SendName   string `json:"send_name,omitempty"`
    

    // 交易订单号（组合表示合并订单）
    
    TradeOrderList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"trade_order_list,omitempty"`
    

    // 快递服务商CODE
    
    CpCode   string `json:"cp_code,omitempty"`
    

    // 电子面单单号
    
    WaybillCode   string `json:"waybill_code,omitempty"`
    

    // 快递服务产品类型编码
    
    ProductType   string `json:"product_type,omitempty"`
    

    // 订单渠道类型
    
    OrderChannelsType   string `json:"order_channels_type,omitempty"`
    

    // 使用者ID
    
    RealUserId   int64 `json:"real_user_id,omitempty"`
    

    // 包裹体积 单位为ML(毫升)或立方厘米
    
    Volume   int64 `json:"volume,omitempty"`
    

    // 包裹里面的商品名称
    
    PackageItems  struct {
        PackageItem  []PackageItem `json:"package_item,omitempty"`
    } `json:"package_items,omitempty"`
    

    // 物流服务能力集合
    
    LogisticsServiceList  struct {
        LogisticsService  []LogisticsService `json:"logistics_service,omitempty"`
    } `json:"logistics_service_list,omitempty"`
    

    // 收\发货地址
    
    ConsigneeAddress  *struct {
        WaybillAddress  *WaybillAddress `json:"waybill_address,omitempty"`
    } `json:"consignee_address,omitempty"`
    

    // 收件人电话
    
    ConsigneePhone   string `json:"consignee_phone,omitempty"`
    

    // ERP 订单号或包裹号
    
    PackageId   string `json:"package_id,omitempty"`
    

}
*/

// WaybillApplyFullUpdateRequest 
type WaybillApplyFullUpdateRequest struct {

    // 发件人联系方式
    SendPhone   string `json:"send_phone,omitempty"`

    // 收件人姓名
    ConsigneeName   string `json:"consignee_name,omitempty"`

    // 包裹重量 单位为G(克)
    Weight   int64 `json:"weight,omitempty"`

    // 发件人姓名
    SendName   string `json:"send_name,omitempty"`

    // 交易订单号（组合表示合并订单）
    TradeOrderList   []string `json:"trade_order_list,omitempty"`

    // 快递服务商CODE
    CpCode   string `json:"cp_code,omitempty"`

    // 电子面单单号
    WaybillCode   string `json:"waybill_code,omitempty"`

    // 快递服务产品类型编码
    ProductType   string `json:"product_type,omitempty"`

    // 订单渠道类型
    OrderChannelsType   string `json:"order_channels_type,omitempty"`

    // 使用者ID
    RealUserId   int64 `json:"real_user_id,omitempty"`

    // 包裹体积 单位为ML(毫升)或立方厘米
    Volume   int64 `json:"volume,omitempty"`

    // 包裹里面的商品名称
    PackageItems   []PackageItem `json:"package_items,omitempty"`

    // 物流服务能力集合
    LogisticsServiceList   []LogisticsService `json:"logistics_service_list,omitempty"`

    // 收\发货地址
    ConsigneeAddress   *WaybillAddress `json:"consignee_address,omitempty"`

    // 收件人电话
    ConsigneePhone   string `json:"consignee_phone,omitempty"`

    // ERP 订单号或包裹号
    PackageId   string `json:"package_id,omitempty"`

}
