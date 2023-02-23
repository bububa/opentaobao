package tmallservice

// WorkerDto 结构体
type WorkerDto struct {
	// 1111
	ServiceAreas []DivisionDto `json:"service_areas,omitempty" xml:"service_areas>division_dto,omitempty"`
	// 11
	ServiceTypes []string `json:"service_types,omitempty" xml:"service_types>string,omitempty"`
	// 工人姓名
	RealName string `json:"real_name,omitempty" xml:"real_name,omitempty"`
	// 昵称
	NickName string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	// 手机号码
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 区
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 镇
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// 详细地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 头像图片url
	ProfilePictureUrl string `json:"profile_picture_url,omitempty" xml:"profile_picture_url,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 111
	IdentityId string `json:"identity_id,omitempty" xml:"identity_id,omitempty"`
	// 11
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 1111
	ProviderName string `json:"provider_name,omitempty" xml:"provider_name,omitempty"`
	// 111
	RegisterTime string `json:"register_time,omitempty" xml:"register_time,omitempty"`
	// 1111
	WorkType string `json:"work_type,omitempty" xml:"work_type,omitempty"`
	// 111
	HandheldCardPic string `json:"handheld_card_pic,omitempty" xml:"handheld_card_pic,omitempty"`
	// 111
	Photo string `json:"photo,omitempty" xml:"photo,omitempty"`
	// 11
	IdCardPicBack string `json:"id_card_pic_back,omitempty" xml:"id_card_pic_back,omitempty"`
	// 11
	IdCardPic string `json:"id_card_pic,omitempty" xml:"id_card_pic,omitempty"`
	// 工人所属行业类型
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 覆盖的service_code列表，|隔开
	ServiceCodes string `json:"service_codes,omitempty" xml:"service_codes,omitempty"`
	// 工人支持的商品类目，格式：类目id1|类目id2
	CoverCategoryIds string `json:"cover_category_ids,omitempty" xml:"cover_category_ids,omitempty"`
	// 网点编码
	ServiceStoreCode string `json:"service_store_code,omitempty" xml:"service_store_code,omitempty"`
	// 工人ID
	WorkerId int64 `json:"worker_id,omitempty" xml:"worker_id,omitempty"`
	// 111
	ProviderId int64 `json:"provider_id,omitempty" xml:"provider_id,omitempty"`
}
