package wdk

// SkuSupplierDo 结构体
type SkuSupplierDo struct {
	// 商品针对该供应商是否可以退货；因为淘鲜达商家不使用平台的采配功能，建议传默认值， 1：是  0：否（默认为0）
	ReturnFlag int64 `json:"return_flag,omitempty" xml:"return_flag,omitempty"`
	// 最小起订量
	Minimum string `json:"minimum,omitempty" xml:"minimum,omitempty"`
	// 采购单价；淘鲜达合作商家填默认值0；单位:元
	PurchasePrice string `json:"purchase_price,omitempty" xml:"purchase_price,omitempty"`
	// 供应商编码
	SupplierNo string `json:"supplier_no,omitempty" xml:"supplier_no,omitempty"`
	// 供应商编码名称
	SupplierName string `json:"supplier_name,omitempty" xml:"supplier_name,omitempty"`
	// 是否主供应商 1是主供应商 0非主供
	MainFlag int64 `json:"main_flag,omitempty" xml:"main_flag,omitempty"`
}
