package nlife

// DeliverDetailDO 
type DeliverDetailDO struct {
    // 门店id
    StoreId   int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
    // 发货单批次
    BatchNo   int64 `json:"batch_no,omitempty" xml:"batch_no,omitempty"`
    // 发货单号
    ConsignNo   string `json:"consign_no,omitempty" xml:"consign_no,omitempty"`
    // 物流公司
    LogisticsCompany   string `json:"logistics_company,omitempty" xml:"logistics_company,omitempty"`
    // 物流单号
    LogisticsNo   string `json:"logistics_no,omitempty" xml:"logistics_no,omitempty"`
    // 关联的门店采购订单号
    TradeNo   string `json:"trade_no,omitempty" xml:"trade_no,omitempty"`
    // 供应商id
    SupplierId   int64 `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
    // 发货时间
    GmtConsign   string `json:"gmt_consign,omitempty" xml:"gmt_consign,omitempty"`
    // 发货单状态: 5-已发货; 10-已到货(入库中); 11-待供应商确认差异; 7-已签收
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 发货商品列表
    DeliverItemList   []Deliveritemlist `json:"deliver_item_list,omitempty" xml:"deliver_item_list>deliveritemlist,omitempty"`
}
