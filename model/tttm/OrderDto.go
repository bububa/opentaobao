package tttm

// OrderDto 
type OrderDto struct {
    // 订单id
    OrderId   string `json:"order_id,omitempty" xml:"order_id,omitempty"`
    // 订单总金额
    TotalPrice   string `json:"total_price,omitempty" xml:"total_price,omitempty"`
    // 合同类型
    ContractType   int64 `json:"contract_type,omitempty" xml:"contract_type,omitempty"`
    // 公司名称
    CompanyName   string `json:"company_name,omitempty" xml:"company_name,omitempty"`
    // 备注
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    // 下单货品
    OrderProductList   []OrderProductDto `json:"order_product_list,omitempty" xml:"order_product_list>order_product_dto,omitempty"`
    // 下单时间
    OrderTime   string `json:"order_time,omitempty" xml:"order_time,omitempty"`
    // 生产状态
    ProduceStatus   int64 `json:"produce_status,omitempty" xml:"produce_status,omitempty"`
    // 订单商品列表
    OrderProducts   string `json:"order_products,omitempty" xml:"order_products,omitempty"`
    // 采购单号
    ExternalId   string `json:"external_id,omitempty" xml:"external_id,omitempty"`
    // 交付时间
    DueTime   string `json:"due_time,omitempty" xml:"due_time,omitempty"`
    // 附件列表
    Annexes   string `json:"annexes,omitempty" xml:"annexes,omitempty"`
    // 订单总数
    TotalAmount   string `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
    // 业务来源
    BizSource   int64 `json:"biz_source,omitempty" xml:"biz_source,omitempty"`
    // 附件列表
    AnnexesJson   string `json:"annexes_json,omitempty" xml:"annexes_json,omitempty"`
    // 加工类型
    WorkingType   int64 `json:"working_type,omitempty" xml:"working_type,omitempty"`
}
