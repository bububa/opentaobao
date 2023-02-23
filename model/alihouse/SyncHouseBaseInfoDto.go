package alihouse

// SyncHouseBaseInfoDto 结构体
type SyncHouseBaseInfoDto struct {
	// 1
	CommunityOuterId string `json:"community_outer_id,omitempty" xml:"community_outer_id,omitempty"`
	// 1
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 1
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 1
	ECode string `json:"e_code,omitempty" xml:"e_code,omitempty"`
	// 1
	OuterCompanyId string `json:"outer_company_id,omitempty" xml:"outer_company_id,omitempty"`
	// 1
	OuterAgentId string `json:"outer_agent_id,omitempty" xml:"outer_agent_id,omitempty"`
	// 1
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 1
	OuterHouseBaseId string `json:"outer_house_base_id,omitempty" xml:"outer_house_base_id,omitempty"`
	// 1
	OuterLayoutsId string `json:"outer_layouts_id,omitempty" xml:"outer_layouts_id,omitempty"`
	// 1
	ShowPrice string `json:"show_price,omitempty" xml:"show_price,omitempty"`
	// 1
	InsideArea string `json:"inside_area,omitempty" xml:"inside_area,omitempty"`
	// 1
	BuildingArea string `json:"building_area,omitempty" xml:"building_area,omitempty"`
	// 房源封面图列表
	CoverPicture string `json:"cover_picture,omitempty" xml:"cover_picture,omitempty"`
	// 1
	OnlineStatus int64 `json:"online_status,omitempty" xml:"online_status,omitempty"`
	// 1
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 1
	Scene int64 `json:"scene,omitempty" xml:"scene,omitempty"`
	// 1
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
	// 1
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 1
	MerchantOpenId int64 `json:"merchant_open_id,omitempty" xml:"merchant_open_id,omitempty"`
	// 1
	HouseCategory int64 `json:"house_category,omitempty" xml:"house_category,omitempty"`
	// 1
	HouseModel int64 `json:"house_model,omitempty" xml:"house_model,omitempty"`
	// 1
	RentPrice int64 `json:"rent_price,omitempty" xml:"rent_price,omitempty"`
	// 1
	Room int64 `json:"room,omitempty" xml:"room,omitempty"`
	// 1
	Hall int64 `json:"hall,omitempty" xml:"hall,omitempty"`
	// 1
	Toilet int64 `json:"toilet,omitempty" xml:"toilet,omitempty"`
	// 1
	Kitchen int64 `json:"kitchen,omitempty" xml:"kitchen,omitempty"`
	// 1
	Balcony int64 `json:"balcony,omitempty" xml:"balcony,omitempty"`
	// IC商品id
	NewItemId int64 `json:"new_item_id,omitempty" xml:"new_item_id,omitempty"`
	// 委托状态
	EntrustStatus int64 `json:"entrust_status,omitempty" xml:"entrust_status,omitempty"`
	// 模版id
	AgreementId int64 `json:"agreement_id,omitempty" xml:"agreement_id,omitempty"`
	// 寄租-房源房东类型 交易类型不为空，房源房东类型 1、企业B房东 2、自然人主体
	LandlordType int64 `json:"landlord_type,omitempty" xml:"landlord_type,omitempty"`
	// 房东自然人主体id
	NaturalMerchantOpenId int64 `json:"natural_merchant_open_id,omitempty" xml:"natural_merchant_open_id,omitempty"`
	// 收款模式1-直收    2-代收
	CollectionType int64 `json:"collection_type,omitempty" xml:"collection_type,omitempty"`
	// 代收主体id 当收款模式为代收是必填
	InsteadMerchantOpenId int64 `json:"instead_merchant_open_id,omitempty" xml:"instead_merchant_open_id,omitempty"`
	// 1-静默 2-授权代理人
	SignType int64 `json:"sign_type,omitempty" xml:"sign_type,omitempty"`
	// 寄租模式下-企业B房东主体ID
	EnterpriseMerchantOpenId int64 `json:"enterprise_merchant_open_id,omitempty" xml:"enterprise_merchant_open_id,omitempty"`
	// 授权的主体ID  说明： 当签约类型是授权代理人时，授权代理人ID需要必填
	SigningMerchantOpenId int64 `json:"signing_merchant_open_id,omitempty" xml:"signing_merchant_open_id,omitempty"`
}
