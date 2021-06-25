package logistic

// Param 
type Param struct {

    // appId
    AppId   string `json:"app_id,omitempty"`

    // 外部订单号
    PartnerOrderCode   string `json:"partner_order_code,omitempty"`

    // 渠道code
    ChannelId   string `json:"channel_id,omitempty"`

    // 取消原因描述
    OrderCancelDesc   string `json:"order_cancel_desc,omitempty"`

    // 取消code
    OrderCancelCode   int64 `json:"order_cancel_code,omitempty"`

    // 取消原因code  1:用户取消
    OrderCancelReasonCode   int64 `json:"order_cancel_reason_code,omitempty"`

    // 门店编码，对应大润发deliveryDockCode
    StoreCodes   []String `json:"store_codes,omitempty"`

    // 服务包code
    ServicePackageCode   string `json:"service_package_code,omitempty"`

    // 商户code
    MerchantCode   string `json:"merchant_code,omitempty"`

    // 门店code
    ChainstoreCode   string `json:"chainstore_code,omitempty"`

    // 经纬度来源。1-腾讯，2-百度，3-高德
    PositionSource   int64 `json:"position_source,omitempty"`

    // 文本地址
    Address   string `json:"address,omitempty"`

    // 门店名字
    ChainStoreName   string `json:"chain_store_name,omitempty"`

    // 门店code
    ChainStoreCode   string `json:"chain_store_code,omitempty"`

    // 联系手机
    ContactPhone   string `json:"contact_phone,omitempty"`

    // 经度
    Longitude   string `json:"longitude,omitempty"`

    // 纬度
    Latitude   string `json:"latitude,omitempty"`

    // 门店列表
    Chainstores   []ChainStore `json:"chainstores,omitempty"`

    // 交易唯一id，汇金outBizId
    TradeId   int64 `json:"trade_id,omitempty"`

    // 汇金扣拥状态
    Success   bool `json:"success,omitempty"`

    // 汇金唯一id
    HjTradeNo   int64 `json:"hj_trade_no,omitempty"`

    // 错误详情
    ErrorMsg   string `json:"error_msg,omitempty"`

}
