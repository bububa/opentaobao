package wdk

// StockChangeDto 结构体
type StockChangeDto struct {
	// itemList
	ItemList []StockChangeDetailDto `json:"item_list,omitempty" xml:"item_list>stock_change_detail_dto,omitempty"`
	// 部门编码
	DeptCode string `json:"dept_code,omitempty" xml:"dept_code,omitempty"`
	// remark
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// adjustDescribe
	AdjustDescribe string `json:"adjust_describe,omitempty" xml:"adjust_describe,omitempty"`
	// occurrenceDate
	OccurrenceDate string `json:"occurrence_date,omitempty" xml:"occurrence_date,omitempty"`
	// warehouseCode
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// supplierCode
	SupplierCode string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
	// documentNo
	DocumentNo string `json:"document_no,omitempty" xml:"document_no,omitempty"`
	// uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 费用承担部门（按需取）
	CostDutyDeptCode string `json:"cost_duty_dept_code,omitempty" xml:"cost_duty_dept_code,omitempty"`
	// 单据类型
	DocumentType int64 `json:"document_type,omitempty" xml:"document_type,omitempty"`
}
