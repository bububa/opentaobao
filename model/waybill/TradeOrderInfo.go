package waybill

// TradeOrderInfo 
type TradeOrderInfo struct {
    // 商品名称
    ItemName   string `json:"item_name,omitempty" xml:"item_name,omitempty"`
    // 收货人
    ConsigneeName   string `json:"consignee_name,omitempty" xml:"consignee_name,omitempty"`
    // 是否阿里系订单
    AliOrder   bool `json:"ali_order,omitempty" xml:"ali_order,omitempty"`
    // 大头笔
    ShortAddress   string `json:"short_address,omitempty" xml:"short_address,omitempty"`
    // 订单渠道
    OrderChannelsType   string `json:"order_channels_type,omitempty" xml:"order_channels_type,omitempty"`
    // 交易订单列表
    TradeOrderList   []string `json:"trade_order_list,omitempty" xml:"trade_order_list>string,omitempty"`
    // 面单号
    WaybillCode   string `json:"waybill_code,omitempty" xml:"waybill_code,omitempty"`
    // 收货人联系方式
    ConsigneePhone   string `json:"consignee_phone,omitempty" xml:"consignee_phone,omitempty"`
    // 收货人地址
    ConsigneeAddress   *WaybillAddress `json:"consignee_address,omitempty" xml:"consignee_address,omitempty"`
    // 发货人联系方式
    SendPhone   string `json:"send_phone,omitempty" xml:"send_phone,omitempty"`
    // 包裹重量（克）
    Weight   int64 `json:"weight,omitempty" xml:"weight,omitempty"`
    // 发货人姓名
    SendName   string `json:"send_name,omitempty" xml:"send_name,omitempty"`
    // 订单渠道来源
    OrderType   int64 `json:"order_type,omitempty" xml:"order_type,omitempty"`
    // 包裹中的商品类型
    PackageItems   []PackageItem `json:"package_items,omitempty" xml:"package_items>package_item,omitempty"`
    // 物流服务能力集合
    LogisticsServiceList   []LogisticsService `json:"logistics_service_list,omitempty" xml:"logistics_service_list>logistics_service,omitempty"`
    // 快递服务产品类型编码
    ProductType   string `json:"product_type,omitempty" xml:"product_type,omitempty"`
    // 使用者ID
    RealUserId   int64 `json:"real_user_id,omitempty" xml:"real_user_id,omitempty"`
    // 包裹体积（立方厘米）
    Volume   int64 `json:"volume,omitempty" xml:"volume,omitempty"`
    // 包裹号(或者ERP订单号)
    PackageId   string `json:"package_id,omitempty" xml:"package_id,omitempty"`
}
