package wdk

// StockShiftDto 结构体
type StockShiftDto struct {
	// itemList
	ItemList []StockShiftDetailDto `json:"item_list,omitempty" xml:"item_list>stock_shift_detail_dto,omitempty"`
	// remark
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// occurrenceDate
	OccurrenceDate string `json:"occurrence_date,omitempty" xml:"occurrence_date,omitempty"`
	// warehouseCode
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// documentNo
	DocumentNo string `json:"document_no,omitempty" xml:"document_no,omitempty"`
	// uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// shiftDescribe
	ShiftDescribe string `json:"shift_describe,omitempty" xml:"shift_describe,omitempty"`
}
