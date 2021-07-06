package drugtrace

// PurchaseInfoDto 结构体
type PurchaseInfoDto struct {
	// 采购管理图片（上传图片）
	PurchasePictures []string `json:"purchase_pictures,omitempty" xml:"purchase_pictures>string,omitempty"`
	// 采购日期yyyy-MM-dd
	PurchaseDate string `json:"purchase_date,omitempty" xml:"purchase_date,omitempty"`
	// 供应商名称
	SupplierName string `json:"supplier_name,omitempty" xml:"supplier_name,omitempty"`
	// 采购数量
	PurchaseNum string `json:"purchase_num,omitempty" xml:"purchase_num,omitempty"`
	// 药材批号
	MaterialsBatchNo string `json:"materials_batch_no,omitempty" xml:"materials_batch_no,omitempty"`
}
