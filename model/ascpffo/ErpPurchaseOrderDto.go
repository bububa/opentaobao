package ascpffo

// ErpPurchaseOrderDTO 
type ErpPurchaseOrderDTO struct {
    // 入库完结时间
    ActualInboundTime   int64 `json:"actual_inbound_time,omitempty" xml:"actual_inbound_time,omitempty"`
    // 预约单号
    AppointOrderNo   string `json:"appoint_order_no,omitempty" xml:"appoint_order_no,omitempty"`
    // 创建人
    Creator   string `json:"creator,omitempty" xml:"creator,omitempty"`
    // 扩展字段
    ExtendFields   string `json:"extend_fields,omitempty" xml:"extend_fields,omitempty"`
    // 采购单创建时间
    GmtCreate   int64 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 采购单失效时间
    GmtExpiration   int64 `json:"gmt_expiration,omitempty" xml:"gmt_expiration,omitempty"`
    // 物流单号
    LbxNo   string `json:"lbx_no,omitempty" xml:"lbx_no,omitempty"`
    // 预计到仓时间
    PreArriveTime   int64 `json:"pre_arrive_time,omitempty" xml:"pre_arrive_time,omitempty"`
    // 采购单号
    PurchaseOrderNo   string `json:"purchase_order_no,omitempty" xml:"purchase_order_no,omitempty"`
    // 收货数量
    ReceivedQuantity   int64 `json:"received_quantity,omitempty" xml:"received_quantity,omitempty"`
    // 单据状态
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 状态描述
    StatusDesc   string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
    // 仓编码
    StoreCode   string `json:"store_code,omitempty" xml:"store_code,omitempty"`
    // 仓名称
    StoreName   string `json:"store_name,omitempty" xml:"store_name,omitempty"`
    // 供应商Id
    SupplierId   int64 `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
    // 供应商名称
    SupplierName   string `json:"supplier_name,omitempty" xml:"supplier_name,omitempty"`
    // 采购单总金额
    TotalAmount   string `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
    // 采购单总数量
    TotalQuantity   int64 `json:"total_quantity,omitempty" xml:"total_quantity,omitempty"`
    // 采购单SKU数量
    TotalSkuCount   int64 `json:"total_sku_count,omitempty" xml:"total_sku_count,omitempty"`
}
