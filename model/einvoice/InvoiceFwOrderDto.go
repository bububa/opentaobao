package einvoice

// InvoiceFwOrderDTO 
type InvoiceFwOrderDTO struct {
    // 联系人对象
    InvoiceContact   *InvoiceContactDTO `json:"invoice_contact,omitempty" xml:"invoice_contact,omitempty"`
    // 订购时间，格式yyyy-MM-dd HH:mm:ss
    BuyDate   string `json:"buy_date,omitempty" xml:"buy_date,omitempty"`
    // 实付总金额，单位元，最多2位小数
    FactTotalFee   string `json:"fact_total_fee,omitempty" xml:"fact_total_fee,omitempty"`
    // 商品规格ID
    ItemId   string `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 订购商品ID
    ArticleId   string `json:"article_id,omitempty" xml:"article_id,omitempty"`
    // 服务市场订单号
    OrderId   string `json:"order_id,omitempty" xml:"order_id,omitempty"`
    // 产品Code,中台定义的税控产品Code
    ProductCode   string `json:"product_code,omitempty" xml:"product_code,omitempty"`
    // 纳税人识别号
    PayeeRegisterNo   string `json:"payee_register_no,omitempty" xml:"payee_register_no,omitempty"`
    // 税控设备订购单ID
    FlowId   string `json:"flow_id,omitempty" xml:"flow_id,omitempty"`
    // 入驻类型，可选值： 新订购：new，已有税控初始化：online
    RegisterType   string `json:"register_type,omitempty" xml:"register_type,omitempty"`
    // 服务结束时间，格式yyyy-MM-dd HH:mm:ss
    ServEndTime   string `json:"serv_end_time,omitempty" xml:"serv_end_time,omitempty"`
    // 服务起始时间，格式yyyy-MM-dd HH:mm:ss
    ServStartTime   string `json:"serv_start_time,omitempty" xml:"serv_start_time,omitempty"`
    // 商品名称
    ArticleName   string `json:"article_name,omitempty" xml:"article_name,omitempty"`
}
