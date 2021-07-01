package happytrip

// ResourceMainDto 
type ResourceMainDto struct {
    // 成人成本价
    AdultCost   string `json:"adult_cost,omitempty" xml:"adult_cost,omitempty"`
    // 成人数
    AdultCount   int64 `json:"adult_count,omitempty" xml:"adult_count,omitempty"`
    // 成人单价
    AdultPrice   string `json:"adult_price,omitempty" xml:"adult_price,omitempty"`
    // 代理商ID
    AgencyId   string `json:"agency_id,omitempty" xml:"agency_id,omitempty"`
    // 代理商名称
    AgencyName   string `json:"agency_name,omitempty" xml:"agency_name,omitempty"`
    // 代理商订单号
    AgencyOrderId   string `json:"agency_order_id,omitempty" xml:"agency_order_id,omitempty"`
    // 交易号
    AlipayTradeNo   string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
    // 是否b2g 0不是 1是
    B2gFlag   int64 `json:"b2g_flag,omitempty" xml:"b2g_flag,omitempty"`
    // 返回城市code
    BackCityCode   string `json:"back_city_code,omitempty" xml:"back_city_code,omitempty"`
    // 返回城市名
    BackCityName   string `json:"back_city_name,omitempty" xml:"back_city_name,omitempty"`
    // 返回时间
    BackTime   string `json:"back_time,omitempty" xml:"back_time,omitempty"`
    // 基础价格
    BaseAmount   string `json:"base_amount,omitempty" xml:"base_amount,omitempty"`
    // 基础价格货币代码
    BaseAmountCurrencyCode   string `json:"base_amount_currency_code,omitempty" xml:"base_amount_currency_code,omitempty"`
    // 基础价格小数点位数
    BaseAmountDecimalPlaces   string `json:"base_amount_decimal_places,omitempty" xml:"base_amount_decimal_places,omitempty"`
    // 资源预定城市code
    BookCityCode   string `json:"book_city_code,omitempty" xml:"book_city_code,omitempty"`
    // 资源预定城市名
    BookCityName   string `json:"book_city_name,omitempty" xml:"book_city_name,omitempty"`
    // 资源预定时间
    BookTime   string `json:"book_time,omitempty" xml:"book_time,omitempty"`
    // 是否合约 0非合约 1合约
    ContractFlag   int64 `json:"contract_flag,omitempty" xml:"contract_flag,omitempty"`
    // 合约价格
    ContractPrice   string `json:"contract_price,omitempty" xml:"contract_price,omitempty"`
    // 结束地址
    EndAddr   string `json:"end_addr,omitempty" xml:"end_addr,omitempty"`
    // 结束城市code
    EndCityCode   string `json:"end_city_code,omitempty" xml:"end_city_code,omitempty"`
    // 结束城市名
    EndCityName   string `json:"end_city_name,omitempty" xml:"end_city_name,omitempty"`
    // 结束时间
    EndDate   string `json:"end_date,omitempty" xml:"end_date,omitempty"`
    // 结束地纬度
    EndLatitude   string `json:"end_latitude,omitempty" xml:"end_latitude,omitempty"`
    // 结束地经度
    EndLongitude   string `json:"end_longitude,omitempty" xml:"end_longitude,omitempty"`
    // 资源失效时间
    ExpireTime   string `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    // 主键
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 最低价
    LowerAmount   string `json:"lower_amount,omitempty" xml:"lower_amount,omitempty"`
    // 最低价货币代码
    LowerAmountCurrencyCode   string `json:"lower_amount_currency_code,omitempty" xml:"lower_amount_currency_code,omitempty"`
    // 最低价小数点位数
    LowerAmountDecimalPlaces   string `json:"lower_amount_decimal_places,omitempty" xml:"lower_amount_decimal_places,omitempty"`
    // 订单号
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    // 计划时间
    PlanTime   string `json:"plan_time,omitempty" xml:"plan_time,omitempty"`
    // 飞猪询价单PNR
    Pnr   string `json:"pnr,omitempty" xml:"pnr,omitempty"`
    // PNR失效时间
    PnrExpireTime   string `json:"pnr_expire_time,omitempty" xml:"pnr_expire_time,omitempty"`
    // 说明描述
    Reason   string `json:"reason,omitempty" xml:"reason,omitempty"`
    // 说明Code
    ReasonCode   string `json:"reason_code,omitempty" xml:"reason_code,omitempty"`
    // 退款原因
    RefundReason   string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
    // 退款状态
    RefundStatus   string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
    // 退款类型
    RefundType   string `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
    // 资源ID，只有打包类订单才有resourceId，例如GT机票订单，内部还有火车，酒店等，单品类可以存外系统订单号或者不存
    ResourceId   string `json:"resource_id,omitempty" xml:"resource_id,omitempty"`
    // 资源名称
    ResourceName   string `json:"resource_name,omitempty" xml:"resource_name,omitempty"`
    // 资源状态
    ResourceStatus   string `json:"resource_status,omitempty" xml:"resource_status,omitempty"`
    // 资源状态描述
    ResourceStatusDesc   string `json:"resource_status_desc,omitempty" xml:"resource_status_desc,omitempty"`
    // 资源类型  -1, "未知资源类型"  100110, "飞猪机票国内直订"  100210, "飞猪机票国际直订"  100211, "飞猪机票国际询价单"  101110, "飞猪酒店国内直订"  110210, "GT机票国际直订"  110211, "GT机票国际酒店"  110212, "GT机票国际火车"  110213, "GT机票国际用车"  121210, "HRS酒店国际酒店直订"  132110, "滴滴国内预约用车"  142110, "曹操国内预约用车"  152110, "快滴国内预约用车"
    ResourceType   int64 `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
    // 通用规则字段，比如机票退改签等
    Rule   string `json:"rule,omitempty" xml:"rule,omitempty"`
    // 座位金额
    SeatPrice   string `json:"seat_price,omitempty" xml:"seat_price,omitempty"`
    // 座位金额货币代码
    SeatPriceCurrencyCode   string `json:"seat_price_currency_code,omitempty" xml:"seat_price_currency_code,omitempty"`
    // 座位金额小数点位数
    SeatPriceDecimalPlaces   string `json:"seat_price_decimal_places,omitempty" xml:"seat_price_decimal_places,omitempty"`
    // 起始地址
    StartAddr   string `json:"start_addr,omitempty" xml:"start_addr,omitempty"`
    // 起始城市code
    StartCityCode   string `json:"start_city_code,omitempty" xml:"start_city_code,omitempty"`
    // 起始城市名
    StartCityName   string `json:"start_city_name,omitempty" xml:"start_city_name,omitempty"`
    // 起始时间
    StartDate   string `json:"start_date,omitempty" xml:"start_date,omitempty"`
    // 起始地纬度
    StartLatitude   string `json:"start_latitude,omitempty" xml:"start_latitude,omitempty"`
    // 起始地经度
    StartLongitude   string `json:"start_longitude,omitempty" xml:"start_longitude,omitempty"`
    // 资源提交时间
    SubmitTime   string `json:"submit_time,omitempty" xml:"submit_time,omitempty"`
    // 税收金额（机票国内基建费）
    TaxAmount   string `json:"tax_amount,omitempty" xml:"tax_amount,omitempty"`
    // 税收金额货币代码
    TaxAmountCurrencyCode   string `json:"tax_amount_currency_code,omitempty" xml:"tax_amount_currency_code,omitempty"`
    // 税收金额小数点位数
    TaxAmountDecimalPlaces   string `json:"tax_amount_decimal_places,omitempty" xml:"tax_amount_decimal_places,omitempty"`
    // 资源总价（是否包含税看业务）
    TotalAmount   string `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
    // 总金额货币代码
    TotalAmountCurrencyCode   string `json:"total_amount_currency_code,omitempty" xml:"total_amount_currency_code,omitempty"`
    // 总金额小数点位数
    TotalAmountDecimalPlaces   string `json:"total_amount_decimal_places,omitempty" xml:"total_amount_decimal_places,omitempty"`
    // 资源更新时间
    UpdateTime   string `json:"update_time,omitempty" xml:"update_time,omitempty"`
}
