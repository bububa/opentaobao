package einvoice

// InvoiceApplyItemsDTO 
type InvoiceApplyItemsDTO struct {
    // 交易金额（含税）=?单价*数量。单位：元，格式为2位小数，精度2位小数。开红票时传正数。
    Amount   string `json:"amount,omitempty" xml:"amount,omitempty"`
    // 优惠金额（含税）可为0,交易金额-优惠金额=实付金额。单位：元，格式为2位小数，精度2位小数。开红票时传正数。
    Discount   string `json:"discount,omitempty" xml:"discount,omitempty"`
    // 商品ID.当商户配置了按照商品ID匹配税编的税编规则时，该字段可用于税编规则匹配
    ItemId   string `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 商品名称/货物名称
    ItemName   string `json:"item_name,omitempty" xml:"item_name,omitempty"`
    // 发票项目编号（或商品编号）业务前台未传值中台会做税编规则匹配。业务前台有传值则优先取前台传入的值。
    ItemNo   string `json:"item_no,omitempty" xml:"item_no,omitempty"`
    // 商品类型名称，如新零售、餐饮等等。当商户配置了按照类型匹配税编的税编规则时，该字段可用于税编规则匹配
    ItemType   string `json:"item_type,omitempty" xml:"item_type,omitempty"`
    // 数量，最多6位小数
    Quantity   string `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // 规格型号
    Specification   string `json:"specification,omitempty" xml:"specification,omitempty"`
    // 单价（含税），格式为2位小数。最大支持6位小数，不足2位小数时需转化为2位小数格式。
    TaxPrice   string `json:"tax_price,omitempty" xml:"tax_price,omitempty"`
    // 税率。格式为2位小数，业务前台未传值中台会做税编规则匹配。业务前台有传值则优先取前台传入的值。
    TaxRate   string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
    // 单位
    Unit   string `json:"unit,omitempty" xml:"unit,omitempty"`
    // 0税率标识，只有税率为0的情况才有值，0=出口零税率，1=免税，2=不征收，3=普通零税率
    ZeroRateFlag   string `json:"zero_rate_flag,omitempty" xml:"zero_rate_flag,omitempty"`
    // 开票明细备注
    BizMemo   string `json:"biz_memo,omitempty" xml:"biz_memo,omitempty"`
}
