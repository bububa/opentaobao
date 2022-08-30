package pur

// DeliveryTopDto 结构体
type DeliveryTopDto struct {
	// 发货明细
	List []DeliveryItemTopDto `json:"list,omitempty" xml:"list>delivery_item_top_dto,omitempty"`
	// 验收材料/税务信息
	MaterialInformation []MaterialInformationTopDto `json:"material_information,omitempty" xml:"material_information>material_information_top_dto,omitempty"`
	// 业务类型
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 发货联系人
	ContactPerson string `json:"contact_person,omitempty" xml:"contact_person,omitempty"`
	// 创建人
	Creator string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 发货时间
	DeliveryDate string `json:"delivery_date,omitempty" xml:"delivery_date,omitempty"`
	// 配送类型
	DeliveryType string `json:"delivery_type,omitempty" xml:"delivery_type,omitempty"`
	// 拓展字段
	ExtStr string `json:"ext_str,omitempty" xml:"ext_str,omitempty"`
	// 物流公司
	LogisticsCompany string `json:"logistics_company,omitempty" xml:"logistics_company,omitempty"`
	// 物流编号
	LogisticsNumber string `json:"logistics_number,omitempty" xml:"logistics_number,omitempty"`
	// 联系方式
	PhoneNo string `json:"phone_no,omitempty" xml:"phone_no,omitempty"`
	// 描述
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 饿了么头id
	SourceId string `json:"source_id,omitempty" xml:"source_id,omitempty"`
	// 来源
	SourceType string `json:"source_type,omitempty" xml:"source_type,omitempty"`
	// 租户id
	TenantId int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
	// 是否涉及配送
	Distribution bool `json:"distribution,omitempty" xml:"distribution,omitempty"`
	// 是否免审
	NoApprovalRequired bool `json:"no_approval_required,omitempty" xml:"no_approval_required,omitempty"`
}
