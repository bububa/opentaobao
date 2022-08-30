package simba

// ADGroup 结构体
type ADGroup struct {
	// 推广组主人昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 商品类目id，从根类目到子类目，用空格分割
	CategoryIds string `json:"category_ids,omitempty" xml:"category_ids,omitempty"`
	// 用户设置的上下线状态，offline-下线(暂停竞价)；online-上线；默认为online
	OnlineStatus string `json:"online_status,omitempty" xml:"online_status,omitempty"`
	// online-上线；audit_offline-审核下线；crm_offline-CRM下线；默认为online
	OfflineType string `json:"offline_type,omitempty" xml:"offline_type,omitempty"`
	// 审核下线原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 最后修改时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// title
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// imgUrl
	ImgUrl string `json:"img_url,omitempty" xml:"img_url,omitempty"`
	// outsideItemPrice
	ItemPrice string `json:"item_price,omitempty" xml:"item_price,omitempty"`
	// 推广计划Id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 推广组id
	AdgroupId int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
	// 商品数字id
	NumIid int64 `json:"num_iid,omitempty" xml:"num_iid,omitempty"`
	// 默认出价，单位为分，不能小于5
	DefaultPrice int64 `json:"default_price,omitempty" xml:"default_price,omitempty"`
	// 非搜索出价，单位为分，不能小于5
	NonsearchMaxPrice int64 `json:"nonsearch_max_price,omitempty" xml:"nonsearch_max_price,omitempty"`
	// 已经废弃
	MobileDiscount int64 `json:"mobile_discount,omitempty" xml:"mobile_discount,omitempty"`
	// 非搜索是否使用默认出价，false-不用；true-使用；默认为true;
	IsNonsearchDefaultPrice bool `json:"is_nonsearch_default_price,omitempty" xml:"is_nonsearch_default_price,omitempty"`
}
