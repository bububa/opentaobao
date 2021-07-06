package pur

// SupplierPoDetailDto 结构体
type SupplierPoDetailDto struct {
	// 订单行列表
	PoLineList []Polinelist `json:"po_line_list,omitempty" xml:"po_line_list>polinelist,omitempty"`
	// 供应商名称
	SupplierName string `json:"supplier_name,omitempty" xml:"supplier_name,omitempty"`
	// 订单编号
	PoNo string `json:"po_no,omitempty" xml:"po_no,omitempty"`
	// 联系人
	SupplierContact string `json:"supplier_contact,omitempty" xml:"supplier_contact,omitempty"`
	// 联系人电话
	SupplierMobile string `json:"supplier_mobile,omitempty" xml:"supplier_mobile,omitempty"`
	// 订单生效日期
	EffectTime string `json:"effect_time,omitempty" xml:"effect_time,omitempty"`
	// 需方
	OuName string `json:"ou_name,omitempty" xml:"ou_name,omitempty"`
	// 需方(代号)
	OuCode string `json:"ou_code,omitempty" xml:"ou_code,omitempty"`
	// 合同编号
	ContractNo string `json:"contract_no,omitempty" xml:"contract_no,omitempty"`
	// 采购员姓名
	BuyerName string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// 采购员电话
	BuyMobile string `json:"buy_mobile,omitempty" xml:"buy_mobile,omitempty"`
	// PR单号
	PrNo string `json:"pr_no,omitempty" xml:"pr_no,omitempty"`
	// 注意事项
	Comments string `json:"comments,omitempty" xml:"comments,omitempty"`
}
