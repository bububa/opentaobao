package tmallcar

// SignContractReq 结构体
type SignContractReq struct {
	// 乙方手机号
	SecondPartyMobileNum string `json:"second_party_mobile_num,omitempty" xml:"second_party_mobile_num,omitempty"`
	// 合同类型。购车合同：PURCHASE_VEHICLE；提车合同：GET_VEHICLE
	ContractType string `json:"contract_type,omitempty" xml:"contract_type,omitempty"`
	// 甲方名称
	FirstPartyName string `json:"first_party_name,omitempty" xml:"first_party_name,omitempty"`
	// 合同文件url
	FileUrl string `json:"file_url,omitempty" xml:"file_url,omitempty"`
	// 合同签署日期
	IssueDate string `json:"issue_date,omitempty" xml:"issue_date,omitempty"`
	// 乙方名称
	SecondPartyName string `json:"second_party_name,omitempty" xml:"second_party_name,omitempty"`
	// 乙方身份证号
	SecondPartyIdCard string `json:"second_party_id_card,omitempty" xml:"second_party_id_card,omitempty"`
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}
