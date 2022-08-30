package pur

// MaterialInformationTopDto 结构体
type MaterialInformationTopDto struct {
	// 验收材料
	AcceptanceMaterialList []string `json:"acceptance_material_list,omitempty" xml:"acceptance_material_list>string,omitempty"`
	// 电子发票
	InvoiceFileList []string `json:"invoice_file_list,omitempty" xml:"invoice_file_list>string,omitempty"`
	// gst
	TaxInformationGstList []string `json:"tax_information_gst_list,omitempty" xml:"tax_information_gst_list>string,omitempty"`
	// wht
	TaxInformationWhtList []string `json:"tax_information_wht_list,omitempty" xml:"tax_information_wht_list>string,omitempty"`
	// 订单单号
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
}
