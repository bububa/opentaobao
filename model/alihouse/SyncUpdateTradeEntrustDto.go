package alihouse

// SyncUpdateTradeEntrustDto 结构体
type SyncUpdateTradeEntrustDto struct {
	// 租房凭证
	RentVouchers []string `json:"rent_vouchers,omitempty" xml:"rent_vouchers>string,omitempty"`
	// 外部小区ID
	OuterCommunityId string `json:"outer_community_id,omitempty" xml:"outer_community_id,omitempty"`
	// 外部委托ID
	OuterEntrustId string `json:"outer_entrust_id,omitempty" xml:"outer_entrust_id,omitempty"`
	// 外部门店ID
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 外部经纪人ID
	OuterBrokerId string `json:"outer_broker_id,omitempty" xml:"outer_broker_id,omitempty"`
	// 寄租收费模式 1、房东承担寄租服务费 2、租客承担门店服务费 3、租客承担门店服务费+房东承担寄租服务费
	SendRentChargeModel string `json:"send_rent_charge_model,omitempty" xml:"send_rent_charge_model,omitempty"`
	// 房源地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 房产证地址
	OwnershipAddress string `json:"ownership_address,omitempty" xml:"ownership_address,omitempty"`
	// 业务类型，1-新房，2-二手房 3-租房
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 经营主体ID
	MerchantOpenId int64 `json:"merchant_open_id,omitempty" xml:"merchant_open_id,omitempty"`
	// 交易业务类型，1-B2C，2-C2C，3-寄租
	TradeBizType int64 `json:"trade_biz_type,omitempty" xml:"trade_biz_type,omitempty"`
	// 模版id
	AgreementId int64 `json:"agreement_id,omitempty" xml:"agreement_id,omitempty"`
	// 房东自然人主体id
	NaturalMerchantOpenId int64 `json:"natural_merchant_open_id,omitempty" xml:"natural_merchant_open_id,omitempty"`
	// 委托状态
	EntrustStatus int64 `json:"entrust_status,omitempty" xml:"entrust_status,omitempty"`
	// 寄租模式下-企业B房东主体ID
	EnterpriseMerchantOpenId int64 `json:"enterprise_merchant_open_id,omitempty" xml:"enterprise_merchant_open_id,omitempty"`
	// 收款模式 是	交易类型不为空，必填 1-直收   （默认值） 2-代收
	CollectionType int64 `json:"collection_type,omitempty" xml:"collection_type,omitempty"`
	// 寄租-房源房东类型  交易类型不为空，房源房东类型 1、企业B房东 2、自然人主体
	LandlordType int64 `json:"landlord_type,omitempty" xml:"landlord_type,omitempty"`
	// 代收主体id 当收款模式为代收是必填
	InsteadMerchantOpenId int64 `json:"instead_merchant_open_id,omitempty" xml:"instead_merchant_open_id,omitempty"`
	// ms级时间戳
	EtcVersion int64 `json:"etc_version,omitempty" xml:"etc_version,omitempty"`
}
