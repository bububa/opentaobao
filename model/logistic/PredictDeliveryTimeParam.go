package logistic

// PredictDeliveryTimeParam 
type PredictDeliveryTimeParam struct {

    // 发件人地址
    TransportLocation   *LocationParam `json:"transport_location,omitempty"`

    // 定位来源
    PositionSource   string `json:"position_source,omitempty"`

    // 餐品id
    IndexId   string `json:"index_id,omitempty"`

    // 收件人地址
    ReceiverLocation   *LocationParam `json:"receiver_location,omitempty"`

    // 配送费
    DeliveryTotalPrice   int64 `json:"delivery_total_price,omitempty"`

    // 餐品详情
    ItemInfos   []ItemInfoParam `json:"item_infos,omitempty"`

    // appid
    AppId   string `json:"app_id,omitempty"`

    // 商户code
    MerchantCode   string `json:"merchant_code,omitempty"`

    // 门店code
    ChainStoreCode   string `json:"chain_store_code,omitempty"`

    // 总数量
    OrderAmount   int64 `json:"order_amount,omitempty"`

    // 用户收件地址
    CustomerAddress   string `json:"customer_address,omitempty"`

    // 预计送达时间
    RequireReceiveTime   int64 `json:"require_receive_time,omitempty"`

}
