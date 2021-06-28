package wdk

// WdkWmsPickMedicineQueryResult 
/* model for simplify = false
type WdkWmsPickMedicineQueryResult struct {

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 系统自动生成
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

    // 拣货单维度药品信息list
    
    MedicineItems  struct {
        MedicineItemDo  []MedicineItemDo `json:"medicine_item_do,omitempty"`
    } `json:"medicine_items,omitempty"`
    

    // 履约单维度药品明细
    
    SourceOrderMedicineDTOS  struct {
        SourceOrderMedicineDto  []SourceOrderMedicineDto `json:"source_order_medicine_dto,omitempty"`
    } `json:"source_order_medicine_d_t_o_s,omitempty"`
    

}
*/

// WdkWmsPickMedicineQueryResult 
type WdkWmsPickMedicineQueryResult struct {

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 系统自动生成
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 拣货单维度药品信息list
    MedicineItems   []MedicineItemDo `json:"medicine_items,omitempty"`

    // 履约单维度药品明细
    SourceOrderMedicineDTOS   []SourceOrderMedicineDto `json:"source_order_medicine_d_t_o_s,omitempty"`

}
