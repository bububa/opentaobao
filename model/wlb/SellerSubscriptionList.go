package wlb

// SellerSubscriptionList 结构体
type SellerSubscriptionList struct {
	// 定购用户NICK
	SubscriberUserNick string `json:"subscriber_user_nick,omitempty" xml:"subscriber_user_nick,omitempty"`
	// 定购有效开始日期
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 定购有效结束日期
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 服务编码
	ServiceCode string `json:"service_code,omitempty" xml:"service_code,omitempty"`
	// 服务名
	ServiceName string `json:"service_name,omitempty" xml:"service_name,omitempty"`
	// 服务类型，&lt;br/&gt;STORE 1-仓储、&lt;br/&gt;TMS 2-TMS、&lt;br/&gt;PACKAGE 3-包装服务&lt;br/&gt;SUPPLIER 4-供货&lt;br/&gt;INSTALL 5-安装&lt;br/&gt;COMPLEX_SERVICE 100-综合服务
	ServiceType string `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// 状态&lt;br/&gt;AUDITING 1-待审核&lt;br/&gt;CANCEL 2-撤销&lt;br/&gt;CHECKED 3-审核通过&lt;br/&gt;FAILED 4-审核未通过&lt;br/&gt;SYNCHRONIZING 5-同步中
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 自有仓的别名，当当前订购记录为自有仓时才会有值
	ServiceAlias string `json:"service_alias,omitempty" xml:"service_alias,omitempty"`
	// openuid
	Openuid string `json:"openuid,omitempty" xml:"openuid,omitempty"`
	// 定购ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 定购用户ID
	SubscriberUserId int64 `json:"subscriber_user_id,omitempty" xml:"subscriber_user_id,omitempty"`
	// 服务商ID
	ProviderUserId int64 `json:"provider_user_id,omitempty" xml:"provider_user_id,omitempty"`
	// 服务ID
	ServiceId int64 `json:"service_id,omitempty" xml:"service_id,omitempty"`
	// 父定购ID&lt;br/&gt;可通过该字段来得之服务上下级关系。&lt;br/&gt;例定购了仓储服务，下又有TMS服务。&lt;br/&gt;该字段保存仓储服务ID。
	ParentId int64 `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
	// 联系人地址信息
	WlbPartnerAddress *AddressInfo `json:"wlb_partner_address,omitempty" xml:"wlb_partner_address,omitempty"`
	// 联系人联系详情
	WlbPartnerContact *ContactInfo `json:"wlb_partner_contact,omitempty" xml:"wlb_partner_contact,omitempty"`
	// 判断该仓库是否是实体仓，还是虚拟仓，null是实体仓，10:代表虚拟仓
	IsOwnService int64 `json:"is_own_service,omitempty" xml:"is_own_service,omitempty"`
}
