package pur

// AccessGoodsDto 结构体
type AccessGoodsDto struct {
	// 报价明细
	QuotationList []AccessQuotationDto `json:"quotation_list,omitempty" xml:"quotation_list>access_quotation_dto,omitempty"`
	// 阿里侧合同编号
	ContractCode string `json:"contract_code,omitempty" xml:"contract_code,omitempty"`
	// 外部商家标识
	DataSource string `json:"data_source,omitempty" xml:"data_source,omitempty"`
	// 是否上架商城
	IsApplyDirectoryMall string `json:"is_apply_directory_mall,omitempty" xml:"is_apply_directory_mall,omitempty"`
	// 外部商家商品标记值
	SourceValue string `json:"source_value,omitempty" xml:"source_value,omitempty"`
	// 最小采购量
	MinimumPurchaseQuantity float64 `json:"minimum_purchase_quantity,omitempty" xml:"minimum_purchase_quantity,omitempty"`
	// 阿里侧供应商ID
	SupplierId int64 `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 税率，如6%则为6
	TaxRate float64 `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
}
