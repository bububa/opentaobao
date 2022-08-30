package ascp

// Integer 结构体
type Integer struct {
	// 货品模型
	ItemInfoList []HiErpItemInfo `json:"item_info_list,omitempty" xml:"item_info_list>hi_erp_item_info,omitempty"`
	// 分销商公司/用户名
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 运单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 分销商用户id
	OwnerId int64 `json:"owner_id,omitempty" xml:"owner_id,omitempty"`
}
