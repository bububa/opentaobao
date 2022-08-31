package alihouse

// TradeMerchantOpenDto 结构体
type TradeMerchantOpenDto struct {
	// 授权协议与证明列表
	ProofList []TradeMerchantProofDto `json:"proof_list,omitempty" xml:"proof_list>trade_merchant_proof_dto,omitempty"`
	// 法人证件背面照
	LegalCertBackUrl string `json:"legal_cert_back_url,omitempty" xml:"legal_cert_back_url,omitempty"`
	// 法人证件正面照
	LegalCertFrontUrl string `json:"legal_cert_front_url,omitempty" xml:"legal_cert_front_url,omitempty"`
	// 营业执照有效期时间
	BusinessLicenseExpireTime string `json:"business_license_expire_time,omitempty" xml:"business_license_expire_time,omitempty"`
	// 营业执照图片
	BusinessLicenseUrl string `json:"business_license_url,omitempty" xml:"business_license_url,omitempty"`
	// 营业执照编号
	BusinessLicenseNo string `json:"business_license_no,omitempty" xml:"business_license_no,omitempty"`
	// 法人编号
	LegalCertNo string `json:"legal_cert_no,omitempty" xml:"legal_cert_no,omitempty"`
	// 法人姓名
	LegalName string `json:"legal_name,omitempty" xml:"legal_name,omitempty"`
	// 商家联系方式
	BusinessContactPhone string `json:"business_contact_phone,omitempty" xml:"business_contact_phone,omitempty"`
	// 商家联系人
	BusinessContactName string `json:"business_contact_name,omitempty" xml:"business_contact_name,omitempty"`
	// 商家别名
	AliasName string `json:"alias_name,omitempty" xml:"alias_name,omitempty"`
	// 商家名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 商家类型  01：企业 07：个体工商户
	MerchantType string `json:"merchant_type,omitempty" xml:"merchant_type,omitempty"`
	// 外部id
	OuterCompanyId string `json:"outer_company_id,omitempty" xml:"outer_company_id,omitempty"`
	// 自然人名称
	PersonalName string `json:"personal_name,omitempty" xml:"personal_name,omitempty"`
	// 自然人电话
	PersonalPhone string `json:"personal_phone,omitempty" xml:"personal_phone,omitempty"`
	// 自然人证件编号
	PersonalCertNo string `json:"personal_cert_no,omitempty" xml:"personal_cert_no,omitempty"`
	// 有效期
	PersonalCertExpireTime string `json:"personal_cert_expire_time,omitempty" xml:"personal_cert_expire_time,omitempty"`
	// 证件正面照
	PersonalCertFrontUrl string `json:"personal_cert_front_url,omitempty" xml:"personal_cert_front_url,omitempty"`
	// 证件反面照
	PersonalCertBackUrl string `json:"personal_cert_back_url,omitempty" xml:"personal_cert_back_url,omitempty"`
	// 详细地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 过期时间
	LegalCertExpireTime string `json:"legal_cert_expire_time,omitempty" xml:"legal_cert_expire_time,omitempty"`
	// 外部绑定楼盘ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 外部项目店ID
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 合同快照
	ContractSnapshot string `json:"contract_snapshot,omitempty" xml:"contract_snapshot,omitempty"`
	// 法人证件类型  1身份证 2 港澳居民往来内地通行证；3 台湾同胞往来大陆通行证；4 外国人居留证
	LegalCertType int64 `json:"legal_cert_type,omitempty" xml:"legal_cert_type,omitempty"`
	// 营业执照有效期类型
	BusinessLicenseStatus int64 `json:"business_license_status,omitempty" xml:"business_license_status,omitempty"`
	// 进件类型 1-公司 2-城市品牌店 3-标准门店
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 自然人证件类型
	PersonalCertType int64 `json:"personal_cert_type,omitempty" xml:"personal_cert_type,omitempty"`
	// 自然人有效期类型
	PersonalCertExpireStatus int64 `json:"personal_cert_expire_status,omitempty" xml:"personal_cert_expire_status,omitempty"`
	// 场景类型1-二手房交易 2-租房交易
	SceneType int64 `json:"scene_type,omitempty" xml:"scene_type,omitempty"`
	// 省份ID
	ProvinceId int64 `json:"province_id,omitempty" xml:"province_id,omitempty"`
	// 城市ID
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 区域ID
	AreaId int64 `json:"area_id,omitempty" xml:"area_id,omitempty"`
	// 有效期类型 1-永久 0-非永久
	LegalCertStatus int64 `json:"legal_cert_status,omitempty" xml:"legal_cert_status,omitempty"`
	// 是否支持贷记 1- 支持 0-不支持
	IsSupportCredit int64 `json:"is_support_credit,omitempty" xml:"is_support_credit,omitempty"`
	// 收款类型 1-直收 2-代收
	CollectionType int64 `json:"collection_type,omitempty" xml:"collection_type,omitempty"`
	// 代收主体ID
	InsteadMerchantOpenId int64 `json:"instead_merchant_open_id,omitempty" xml:"instead_merchant_open_id,omitempty"`
	// 签约类型
	SignType int64 `json:"sign_type,omitempty" xml:"sign_type,omitempty"`
	// 代签主体ID
	SigningMerchantOpenId int64 `json:"signing_merchant_open_id,omitempty" xml:"signing_merchant_open_id,omitempty"`
}
