package wdk

// WdkErpArrivalNoticeDTO 
type WdkErpArrivalNoticeDTO struct {
    // 如部分商品不存在，是否允许其他商品入库(1：允许，其他：不允许)
    Bypass   int64 `json:"bypass,omitempty" xml:"bypass,omitempty"`
    // 入库时间，商家系统中记录的本批次商品的实际入库时间
    ArrivalDate   string `json:"arrival_date,omitempty" xml:"arrival_date,omitempty"`
    // 1
    ItemList   []ErpArrivalNoticeDetailDTO `json:"item_list,omitempty" xml:"item_list>erp_arrival_notice_detail_dto,omitempty"`
    // 联系方式，门店联系电话，可以是移动电话
    ContactInfo   string `json:"contact_info,omitempty" xml:"contact_info,omitempty"`
    // 店仓code，指的是入库对象，对应一个物理店或仓编码
    WarehouseCode   string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
    // 供应商code，可选值：当是供应商供货时，提供供应商编码；当是大仓供货时，提供大仓编码；当是调拨入库时，提供对方门店编码
    SupplierCode   string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
    // 失效日期，到货通知单中规定的收货入库时限
    InvalidDate   string `json:"invalid_date,omitempty" xml:"invalid_date,omitempty"`
    // 选填（单据类型为采配单时为必填）单据子类型，入库单据类型为采配单时，需要进一步区分子类型为采购单（供应商）和送货单（DC）
    SubOrderType   int64 `json:"sub_order_type,omitempty" xml:"sub_order_type,omitempty"`
    // 选填(单据子类型为送货单时为必填) 原始单据号，单据类型为采配单，子类型为送货单时，需要提供原始配货申请单据号
    OriginalBillCode   string `json:"original_bill_code,omitempty" xml:"original_bill_code,omitempty"`
    // 单据类型，入库接口的单据类型包括采配单和调拨入单(1 采购单(供应商)； 2 送货单(DC) ； 其他情况默认为0)
    BizOrderType   int64 `json:"biz_order_type,omitempty" xml:"biz_order_type,omitempty"`
    // 单据号
    BizOrderCode   string `json:"biz_order_code,omitempty" xml:"biz_order_code,omitempty"`
}
