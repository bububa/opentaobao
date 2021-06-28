package wdk

// SourceOrderMedicineDto 
/* model for simplify = false
type SourceOrderMedicineDto struct {

    // 药品明细list
    
    MedicineItemDOS  struct {
        MedicineItemDo  []MedicineItemDo `json:"medicine_item_do,omitempty"`
    } `json:"medicine_item_d_o_s,omitempty"`
    

    // 盒马履约单号
    
    SourceOrderCode   string `json:"source_order_code,omitempty"`
    

    // 外部单号(比如饿了么)
    
    OriginalOrderId   string `json:"original_order_id,omitempty"`
    

    // 盒马app订单号
    
    OutMainOrderId   string `json:"out_main_order_id,omitempty"`
    

}
*/

// SourceOrderMedicineDto 
type SourceOrderMedicineDto struct {

    // 药品明细list
    MedicineItemDOS   []MedicineItemDo `json:"medicine_item_d_o_s,omitempty"`

    // 盒马履约单号
    SourceOrderCode   string `json:"source_order_code,omitempty"`

    // 外部单号(比如饿了么)
    OriginalOrderId   string `json:"original_order_id,omitempty"`

    // 盒马app订单号
    OutMainOrderId   string `json:"out_main_order_id,omitempty"`

}
