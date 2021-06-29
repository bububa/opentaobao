package nlife

// ProcurementDetailResponseDo 
type ProcurementDetailResponseDo struct {

    // 采购单号
    
    ProcurementNo   string `json:"procurement_no,omitempty" xml:"procurement_no,omitempty"`
    

    // 收货门店id
    
    StoreId   int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
    

    // 收货门店名称
    
    StoreName   string `json:"store_name,omitempty" xml:"store_name,omitempty"`
    

    // 供应商ID
    
    SupplierId   int64 `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
    

    // 供应商名称
    
    SupplierName   string `json:"supplier_name,omitempty" xml:"supplier_name,omitempty"`
    

    // 合同编号
    
    AgreementNo   string `json:"agreement_no,omitempty" xml:"agreement_no,omitempty"`
    

    // 合同名称
    
    AgreementName   string `json:"agreement_name,omitempty" xml:"agreement_name,omitempty"`
    

    // 采购单状态:1-待提交; 2-待审核; 3-已驳回; 4-待发货; 5-已发货; 6-提交确认; 7-已签收; 8-已完结; 9-已删除; 10-已到货; 11-待供应商确认差异; 12-采购单已生效; 13-已失效
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    

    // 创建时间
    
    CreateTime   string `json:"create_time,omitempty" xml:"create_time,omitempty"`
    

    // 结算计价模式:1-扣点模式;2-采购价模式
    
    PricingModel   int64 `json:"pricing_model,omitempty" xml:"pricing_model,omitempty"`
    

    // poItemList
    
    PoItemList   []Poitemlist `json:"po_item_list,omitempty" xml:"po_item_list,omitempty"`
    

}
