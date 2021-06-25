package icbu

// SourcingTrade 
type SourcingTrade struct {

    // FOB货币价格，枚举值
    FobCurrency   string `json:"fob_currency,omitempty"`

    // FOB最小价格
    FobMinPrice   string `json:"fob_min_price,omitempty"`

    // FOB最大价格
    FobMaxPrice   string `json:"fob_max_price,omitempty"`

    // FOB计量单位，枚举值
    FobUnitType   string `json:"fob_unit_type,omitempty"`

    // 付款方式，枚举值
    PaymentMethods   []String `json:"payment_methods,omitempty"`

    // 最小起订量
    MinOrderQuantity   string `json:"min_order_quantity,omitempty"`

    // 最小起订量计量单位，枚举值
    MinOrderUnitType   string `json:"min_order_unit_type,omitempty"`

    // 供货能力
    SupplyQuantity   string `json:"supply_quantity,omitempty"`

    // 供货能力计量单位，枚举值
    SupplyUnitType   string `json:"supply_unit_type,omitempty"`

    // 供货能力周期，枚举值
    SupplyPeriodType   string `json:"supply_period_type,omitempty"`

    // 发货港口
    DeliveryPort   string `json:"delivery_port,omitempty"`

    // 发货期限
    DeliveryTime   string `json:"delivery_time,omitempty"`

    // 包装信息
    PackagingDesc   string `json:"packaging_desc,omitempty"`

    // 发货周期，发货时间相关建议使用此项
    DeliverPeriods   []DeliverPeriod `json:"deliver_periods,omitempty"`

}
