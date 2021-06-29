package ascpchannel

// Request 
type Request struct {
    // 二级供应商id
    SupplierId   string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
    // 价格类型1-基准价，2-区间价
    PriceType   int64 `json:"price_type,omitempty" xml:"price_type,omitempty"`
    // 失效时间（区间采购价必须）
    PriceDeadlineDate   string `json:"price_deadline_date,omitempty" xml:"price_deadline_date,omitempty"`
    // 供应链仓code
    StoreCodeSets   []string `json:"store_code_sets,omitempty" xml:"store_code_sets>string,omitempty"`
    // 货品id
    ScItemId   int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
    // 操作人
    OperatorName   string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
    // 生效时间（区间采购价必须）
    PriceStartDate   string `json:"price_start_date,omitempty" xml:"price_start_date,omitempty"`
    // 租户ID，请填写商家编码
    TenantId   string `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
    // 默认根据合同；币种，USD-美元，CNY-人民币，RUB-卢布，JPY-日元，EUR-欧元，GBP-英镑，HKD-港币，NZD-新西兰元，SGD-新加坡元，AUD-澳元，KRW-韩元，THB-泰铢，
    Currency   string `json:"currency,omitempty" xml:"currency,omitempty"`
    // 运输方式，1-水路运输，2-航空运输，3-铁路运输，4-公路运输，0-其他
    TransportType   int64 `json:"transport_type,omitempty" xml:"transport_type,omitempty"`
    // 贸易类型,1-EXW(跨境工厂交货价交易),2-CIF(跨境到岸价交易),3-FCA(跨境货交承运人交易),4-FOB(跨境离岸价交易),5-CFR(跨境目的港交货价交易),6-DDU(跨境到仓不含税价格交易),7-日本境内贸易,9-国内一般贸易交易,11-FAS,12-CNF,13-CPT,14-CIP,15-DDP,16-DAT,17-DAP
    TradeType   int64 `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
    // 目的港代码
    PortOfDestination   string `json:"port_of_destination,omitempty" xml:"port_of_destination,omitempty"`
    // 起运港代码
    PortOfLoading   string `json:"port_of_loading,omitempty" xml:"port_of_loading,omitempty"`
    // 采购价（含税）（单位：元)
    PurchasePrice   string `json:"purchase_price,omitempty" xml:"purchase_price,omitempty"`
}
