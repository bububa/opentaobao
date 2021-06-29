package icbuseller

// OpenTradeDetailDTO 
type OpenTradeDetailDTO struct {
    // 阿里id
    BuyerAliId   int64 `json:"buyer_ali_id,omitempty" xml:"buyer_ali_id,omitempty"`
    // 购买人
    BuyerName   string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
    // 公司名称
    CompanyName   string `json:"company_name,omitempty" xml:"company_name,omitempty"`
    // 联系电话
    ContactMobile   string `json:"contact_mobile,omitempty" xml:"contact_mobile,omitempty"`
    // 创建时间
    CreateTime   string `json:"create_time,omitempty" xml:"create_time,omitempty"`
    // 订单货币类型
    Currency   string `json:"currency,omitempty" xml:"currency,omitempty"`
    // 当前状态
    CurrentStatus   string `json:"current_status,omitempty" xml:"current_status,omitempty"`
    // 状态更新时间
    FireTime   string `json:"fire_time,omitempty" xml:"fire_time,omitempty"`
    // 订单编号
    OrderNo   string `json:"order_no,omitempty" xml:"order_no,omitempty"`
    // 订单标题
    OrderTitle   string `json:"order_title,omitempty" xml:"order_title,omitempty"`
    // 订单类型
    OrderType   string `json:"order_type,omitempty" xml:"order_type,omitempty"`
    // 支付通道
    PayChannel   string `json:"pay_channel,omitempty" xml:"pay_channel,omitempty"`
    // 购买数量
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // 退款金额
    RefundPrice   string `json:"refund_price,omitempty" xml:"refund_price,omitempty"`
    // 服务编码
    ServiceCode   string `json:"service_code,omitempty" xml:"service_code,omitempty"`
    // 服务名称
    ServiceName   string `json:"service_name,omitempty" xml:"service_name,omitempty"`
    // 服务评分
    ServiceScore   int64 `json:"service_score,omitempty" xml:"service_score,omitempty"`
    // 规格编码
    SkuCode   string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
    // 规格
    SkuValue   string `json:"sku_value,omitempty" xml:"sku_value,omitempty"`
    // 成交价格
    TransactionPrice   string `json:"transaction_price,omitempty" xml:"transaction_price,omitempty"`
    // 成交单价
    TransactionUnitPrice   string `json:"transaction_unit_price,omitempty" xml:"transaction_unit_price,omitempty"`
    // 购买者登陆账号
    BuyerLoginId   string `json:"buyer_login_id,omitempty" xml:"buyer_login_id,omitempty"`
    // 服务分类id
    ServiceCategoryId   int64 `json:"service_category_id,omitempty" xml:"service_category_id,omitempty"`
    // 服务分类名称
    ServiceCategory   string `json:"service_category,omitempty" xml:"service_category,omitempty"`
    // 用户购买sku
    ServiceSkuLabel   string `json:"service_sku_label,omitempty" xml:"service_sku_label,omitempty"`
}
