package legalsuit

// CourtEntrustModel 结构体
type CourtEntrustModel struct {
	// 委托订单号
	EntrustOrderNumber int64 `json:"entrust_order_number,omitempty" xml:"entrust_order_number,omitempty"`
	// 委托时间
	EntrustTime string `json:"entrust_time,omitempty" xml:"entrust_time,omitempty"`
	// 委托名字
	EntrustName string `json:"entrust_name,omitempty" xml:"entrust_name,omitempty"`
	// 原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 委托类型
	EntrustType string `json:"entrust_type,omitempty" xml:"entrust_type,omitempty"`
	// 供应商code
	SupplierCode string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 委托方名称
	EntrustingParty string `json:"entrusting_party,omitempty" xml:"entrusting_party,omitempty"`
	// 附件数量
	AttachmentCount int64 `json:"attachment_count,omitempty" xml:"attachment_count,omitempty"`
	// 律所执业编号
	FirmBusinessLicNum string `json:"firm_business_lic_num,omitempty" xml:"firm_business_lic_num,omitempty"`
	// 备注
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 附件列表
	AttachmentList []FileModel `json:"attachment_list,omitempty" xml:"attachment_list>file_model,omitempty"`
	// 目标
	Objective string `json:"objective,omitempty" xml:"objective,omitempty"`
	// 委托建议
	Suggest string `json:"suggest,omitempty" xml:"suggest,omitempty"`
}
