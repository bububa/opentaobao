package alihouse

// GoodsDto 结构体
type GoodsDto struct {
	// 服务者id列表
	OuterConsultantIds []string `json:"outer_consultant_ids,omitempty" xml:"outer_consultant_ids>string,omitempty"`
	// 外部委托房源列表
	OuterHouses []GoodsHouseDto `json:"outer_houses,omitempty" xml:"outer_houses>goods_house_dto,omitempty"`
	// 外部私域楼盘id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 品外部唯一码
	OuterTid string `json:"outer_tid,omitempty" xml:"outer_tid,omitempty"`
	// 渠道电话
	ChannelTelephone string `json:"channel_telephone,omitempty" xml:"channel_telephone,omitempty"`
	// 外部项目店id
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 核销有效期(绝对时间)
	VerifyDeadline string `json:"verify_deadline,omitempty" xml:"verify_deadline,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 品类型 1-安心置业 2-特价房 3-0元购 4-大额电商券 5-认购商品 6-楼栋 7-户型 8-房源 9-预存金 10-购房优惠
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 优惠类型
	DiscountType int64 `json:"discount_type,omitempty" xml:"discount_type,omitempty"`
	// 审核状态
	AuditStatus int64 `json:"audit_status,omitempty" xml:"audit_status,omitempty"`
	// 是否电子签约
	IsElectricSign int64 `json:"is_electric_sign,omitempty" xml:"is_electric_sign,omitempty"`
	// 是否商家确客
	IsVerifyConfirm int64 `json:"is_verify_confirm,omitempty" xml:"is_verify_confirm,omitempty"`
	// 合同主体id
	MerchantOpenId int64 `json:"merchant_open_id,omitempty" xml:"merchant_open_id,omitempty"`
	// 电子签章id
	ElecSignatureId int64 `json:"elec_signature_id,omitempty" xml:"elec_signature_id,omitempty"`
	// 电子协议id
	ElecAgreementId int64 `json:"elec_agreement_id,omitempty" xml:"elec_agreement_id,omitempty"`
	// 核销有效期(相对时间)
	VerifyValidDate int64 `json:"verify_valid_date,omitempty" xml:"verify_valid_date,omitempty"`
	// 是否异步销控 0-否 1-是
	IsAsyncSell int64 `json:"is_async_sell,omitempty" xml:"is_async_sell,omitempty"`
	// 限购信息
	LimitExtend *GoodsLimitDto `json:"limit_extend,omitempty" xml:"limit_extend,omitempty"`
	// 结算账户ID
	BankCardId int64 `json:"bank_card_id,omitempty" xml:"bank_card_id,omitempty"`
	// 0-非测试 1-测试
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
	// 版本号(时间戳)
	EtcVersion int64 `json:"etc_version,omitempty" xml:"etc_version,omitempty"`
	// 是否用户确客
	IsUserVerifyConfirm int64 `json:"is_user_verify_confirm,omitempty" xml:"is_user_verify_confirm,omitempty"`
	// 业务类型，1新房，2二手房 3租房
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 是否房源 -1无关联 0部分关联 1全部
	IsAllHouse int64 `json:"is_all_house,omitempty" xml:"is_all_house,omitempty"`
}
