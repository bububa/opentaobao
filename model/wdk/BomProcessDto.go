package wdk

// BomProcessDto 结构体
type BomProcessDto struct {
	// productItemInfos
	ProductItemInfos []BomItemInfos `json:"product_item_infos,omitempty" xml:"product_item_infos>bom_item_infos,omitempty"`
	// materialItemInfos
	MaterialItemInfos []BomItemInfos `json:"material_item_infos,omitempty" xml:"material_item_infos>bom_item_infos,omitempty"`
	// 部门编码
	DeptCode string `json:"dept_code,omitempty" xml:"dept_code,omitempty"`
	// 加工日期
	OccurrenceDate string `json:"occurrence_date,omitempty" xml:"occurrence_date,omitempty"`
	// 加工类型
	OccurrenceType string `json:"occurrence_type,omitempty" xml:"occurrence_type,omitempty"`
	// 店仓code，指的是库调对象，对应一个物理店或仓编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 唯一识别码
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 单据编码
	BomProcessCode string `json:"bom_process_code,omitempty" xml:"bom_process_code,omitempty"`
}
