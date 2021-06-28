package wdk

// StockShiftDTO 
/* model for simplify = false
type StockShiftDTO struct {

    // itemList
    
    ItemList  struct {
        StockShiftDetailDTO  []StockShiftDetailDTO `json:"stock_shift_detail_dto,omitempty"`
    } `json:"item_list,omitempty"`
    

    // remark
    
    Remark   string `json:"remark,omitempty"`
    

    // occurrenceDate
    
    OccurrenceDate   string `json:"occurrence_date,omitempty"`
    

    // warehouseCode
    
    WarehouseCode   string `json:"warehouse_code,omitempty"`
    

    // documentNo
    
    DocumentNo   string `json:"document_no,omitempty"`
    

    // uuid
    
    Uuid   string `json:"uuid,omitempty"`
    

    // shiftDescribe
    
    ShiftDescribe   string `json:"shift_describe,omitempty"`
    

}
*/

// StockShiftDTO 
type StockShiftDTO struct {

    // itemList
    ItemList   []StockShiftDetailDTO `json:"item_list,omitempty"`

    // remark
    Remark   string `json:"remark,omitempty"`

    // occurrenceDate
    OccurrenceDate   string `json:"occurrence_date,omitempty"`

    // warehouseCode
    WarehouseCode   string `json:"warehouse_code,omitempty"`

    // documentNo
    DocumentNo   string `json:"document_no,omitempty"`

    // uuid
    Uuid   string `json:"uuid,omitempty"`

    // shiftDescribe
    ShiftDescribe   string `json:"shift_describe,omitempty"`

}
