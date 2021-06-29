package wdk

// WdkWmsPickMedicineQueryResult 
type WdkWmsPickMedicineQueryResult struct {
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 系统自动生成
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // 拣货单维度药品信息list
    MedicineItems   []MedicineItemDO `json:"medicine_items,omitempty" xml:"medicine_items>medicine_item_do,omitempty"`
    // 履约单维度药品明细
    SourceOrderMedicineDTOS   []SourceOrderMedicineDTO `json:"source_order_medicine_d_t_o_s,omitempty" xml:"source_order_medicine_d_t_o_s>source_order_medicine_dto,omitempty"`
}
