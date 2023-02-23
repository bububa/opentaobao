package alihouse

// CompanyStoreDto 结构体
type CompanyStoreDto struct {
	// 公司外部ID
	OuterCompanyId string `json:"outer_company_id,omitempty" xml:"outer_company_id,omitempty"`
	// 外部城市品牌ID
	OuterCompanyBrandId string `json:"outer_company_brand_id,omitempty" xml:"outer_company_brand_id,omitempty"`
	// 店铺名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 店铺简称
	StoreNameShort string `json:"store_name_short,omitempty" xml:"store_name_short,omitempty"`
	// 门店描述
	StoreInfo string `json:"store_info,omitempty" xml:"store_info,omitempty"`
	// 经纪人昵称
	MainUserNick string `json:"main_user_nick,omitempty" xml:"main_user_nick,omitempty"`
	// 联系人
	ContactMan string `json:"contact_man,omitempty" xml:"contact_man,omitempty"`
	// 联系电话
	ContactPhone string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// 400主号
	MainPhone string `json:"main_phone,omitempty" xml:"main_phone,omitempty"`
	// 400子号
	SubPhone string `json:"sub_phone,omitempty" xml:"sub_phone,omitempty"`
	// 高德经度
	GaodeLongitude string `json:"gaode_longitude,omitempty" xml:"gaode_longitude,omitempty"`
	// 高德纬度
	GaodeLatitude string `json:"gaode_latitude,omitempty" xml:"gaode_latitude,omitempty"`
	// 地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 门店开通业务（1-新房 2-二手房 3-租房,5-交易服务 以英文逗号分隔）
	StoreBizType string `json:"store_biz_type,omitempty" xml:"store_biz_type,omitempty"`
	// 营业执照编号
	CompanyLicenseNo string `json:"company_license_no,omitempty" xml:"company_license_no,omitempty"`
	// 企业法定代表人
	CompanyLegalPerson string `json:"company_legal_person,omitempty" xml:"company_legal_person,omitempty"`
	// 营业执照图片
	CompanyLicensePhoto string `json:"company_license_photo,omitempty" xml:"company_license_photo,omitempty"`
	// 营业执照有效时间
	CompanyLicenseExpireTime string `json:"company_license_expire_time,omitempty" xml:"company_license_expire_time,omitempty"`
	// 外部门店ID
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 公司名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 营业执照地址
	CompanyLicenseUrl string `json:"company_license_url,omitempty" xml:"company_license_url,omitempty"`
	// 门店负责业务（0-二手房 1-长租公寓 2-居间模式 3-新房 5-交易服务）
	PartakeBusiness string `json:"partake_business,omitempty" xml:"partake_business,omitempty"`
	// 标准门店标签
	TagCodes string `json:"tag_codes,omitempty" xml:"tag_codes,omitempty"`
	// 批次号
	BatchNumber string `json:"batch_number,omitempty" xml:"batch_number,omitempty"`
	// 门店图标，有就传，没有会展示默认的兜底图标
	StoreIcon string `json:"store_icon,omitempty" xml:"store_icon,omitempty"`
	// 来源渠道(麦滴、全房通、寓盟、寓小二、趣房、泛慈)
	SourceChannel string `json:"source_channel,omitempty" xml:"source_channel,omitempty"`
	// 签约公司名称
	SigningCompanyName string `json:"signing_company_name,omitempty" xml:"signing_company_name,omitempty"`
	// 外部合作品牌id
	OuterCooperateBrandIds string `json:"outer_cooperate_brand_ids,omitempty" xml:"outer_cooperate_brand_ids,omitempty"`
	// 签约公司id
	OuterSignCompanyId string `json:"outer_sign_company_id,omitempty" xml:"outer_sign_company_id,omitempty"`
	// 扩展信息
	ExtendsInfo string `json:"extends_info,omitempty" xml:"extends_info,omitempty"`
	// 城市ID
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 门店状态
	StoreStatus int64 `json:"store_status,omitempty" xml:"store_status,omitempty"`
	// 是否删除 0 - 否 1 - 是
	IsDeleted int64 `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
	// 是否长期有效 1-是 0-否
	CompanyLicenseStatus int64 `json:"company_license_status,omitempty" xml:"company_license_status,omitempty"`
	// 是否是小ka数据 1-是 0-否
	IsSmallKa int64 `json:"is_small_ka,omitempty" xml:"is_small_ka,omitempty"`
	// 是否为虚拟门店（0-否，在前端展示    1-是，在前端不展示）
	VirtualType int64 `json:"virtual_type,omitempty" xml:"virtual_type,omitempty"`
	// 是否为测试 0-否 1-是
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
	// 0-同步，1-异步
	IsAsync int64 `json:"is_async,omitempty" xml:"is_async,omitempty"`
	// 门店类型：0-二租标准店（默认）   1-新房项目店 2-渠道标准店   5-交易中心店
	StoreType int64 `json:"store_type,omitempty" xml:"store_type,omitempty"`
	// 商务等级（0-普通  1-公寓KA）
	BusinessLevel int64 `json:"business_level,omitempty" xml:"business_level,omitempty"`
	// ms级时间戳
	EtcVersion int64 `json:"etc_version,omitempty" xml:"etc_version,omitempty"`
	// 主营类目 渠道标准店必填 主营类目 1-新房 2-二手房 3-租房 5-交易服务
	MainCategory int64 `json:"main_category,omitempty" xml:"main_category,omitempty"`
	// 子类型
	SubType int64 `json:"sub_type,omitempty" xml:"sub_type,omitempty"`
}
