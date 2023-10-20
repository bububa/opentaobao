package alihouse

import (
	"sync"
)

// SyncExistingHouseEntrustDto 结构体
type SyncExistingHouseEntrustDto struct {
	// 委托人身份正号
	RightsIdCards []string `json:"rights_id_cards,omitempty" xml:"rights_id_cards>string,omitempty"`
	// 委托人姓名
	RightsNames []string `json:"rights_names,omitempty" xml:"rights_names>string,omitempty"`
	// 委托证书图片链接
	EntrustPictures []string `json:"entrust_pictures,omitempty" xml:"entrust_pictures>string,omitempty"`
	// 租房凭证
	RentVouchers []string `json:"rent_vouchers,omitempty" xml:"rent_vouchers>string,omitempty"`
	// 核验码图片
	HouseCheckPictureUrl string `json:"house_check_picture_url,omitempty" xml:"house_check_picture_url,omitempty"`
	// 核验码图片
	HouseCheckCode string `json:"house_check_code,omitempty" xml:"house_check_code,omitempty"`
	// 委托结束时间
	EntrustEndTime string `json:"entrust_end_time,omitempty" xml:"entrust_end_time,omitempty"`
	// 委托开始时间
	EntrustBeginTime string `json:"entrust_begin_time,omitempty" xml:"entrust_begin_time,omitempty"`
	// 佣金费率 百分比小数
	SaleRatioFee string `json:"sale_ratio_fee,omitempty" xml:"sale_ratio_fee,omitempty"`
	// 委托证书图片
	EntrustPicture string `json:"entrust_picture,omitempty" xml:"entrust_picture,omitempty"`
	// 产权证书图片
	CertificatePicture string `json:"certificate_picture,omitempty" xml:"certificate_picture,omitempty"`
	// 委托状态变更原因
	EntrustStatusReason string `json:"entrust_status_reason,omitempty" xml:"entrust_status_reason,omitempty"`
	// 委托价格 单位万元
	HousePrice string `json:"house_price,omitempty" xml:"house_price,omitempty"`
	// 外部经纪人id
	OuterAgentId string `json:"outer_agent_id,omitempty" xml:"outer_agent_id,omitempty"`
	// 外部门店id
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 外部公司id
	OuterCompanyId string `json:"outer_company_id,omitempty" xml:"outer_company_id,omitempty"`
	// 委托人电话
	ConsignerPhone string `json:"consigner_phone,omitempty" xml:"consigner_phone,omitempty"`
	// 委托人姓名
	ConsignerName string `json:"consigner_name,omitempty" xml:"consigner_name,omitempty"`
	// 委托证书编号
	RightsNo string `json:"rights_no,omitempty" xml:"rights_no,omitempty"`
	// 外部委托id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 核验描述
	HouseCheckDesc string `json:"house_check_desc,omitempty" xml:"house_check_desc,omitempty"`
	// 房产证地址
	OwnershipAddress string `json:"ownership_address,omitempty" xml:"ownership_address,omitempty"`
	// 付款方式json
	PayTypeJson string `json:"pay_type_json,omitempty" xml:"pay_type_json,omitempty"`
	// 服务内容
	ServiceContent string `json:"service_content,omitempty" xml:"service_content,omitempty"`
	// 租客要求
	RentRequirement string `json:"rent_requirement,omitempty" xml:"rent_requirement,omitempty"`
	// 可看房时间
	ViewingTime string `json:"viewing_time,omitempty" xml:"viewing_time,omitempty"`
	// 可入住时间
	CheckInTime string `json:"check_in_time,omitempty" xml:"check_in_time,omitempty"`
	// 外部品牌ID
	OuterBrandId string `json:"outer_brand_id,omitempty" xml:"outer_brand_id,omitempty"`
	// 租金
	RentPrice string `json:"rent_price,omitempty" xml:"rent_price,omitempty"`
	// 最新房源维护时间
	MaintenanceTime string `json:"maintenance_time,omitempty" xml:"maintenance_time,omitempty"`
	// 核验信息
	HouseCheckInformation string `json:"house_check_information,omitempty" xml:"house_check_information,omitempty"`
	// 经营主体id
	MerchantOpenId string `json:"merchant_open_id,omitempty" xml:"merchant_open_id,omitempty"`
	// 寄租收费模式	1、房东承担寄租服务费 2、租客承担门店服务费 3、租客承担门店服务费+房东承担寄租服务费
	SendRentChargeModel string `json:"send_rent_charge_model,omitempty" xml:"send_rent_charge_model,omitempty"`
	// 角色列表，英文逗号分隔，委托角色 1首付 2 实勘 3 委托
	RoleList string `json:"role_list,omitempty" xml:"role_list,omitempty"`
	// 核验码状态 0未通过 1通过
	HouseCheckStatus int64 `json:"house_check_status,omitempty" xml:"house_check_status,omitempty"`
	// 委托状态
	EntrustStatus int64 `json:"entrust_status,omitempty" xml:"entrust_status,omitempty"`
	// 委托角色 1首付 2 实勘 3 委托
	Role int64 `json:"role,omitempty" xml:"role,omitempty"`
	// 委托售后id
	EntrustSellingId int64 `json:"entrust_selling_id,omitempty" xml:"entrust_selling_id,omitempty"`
	// 最长租期
	MaxRentTime int64 `json:"max_rent_time,omitempty" xml:"max_rent_time,omitempty"`
	// 最短租期
	MinRentTime int64 `json:"min_rent_time,omitempty" xml:"min_rent_time,omitempty"`
	// 产权证证件类型,1-业务件号，2-证件号，3-商品房备案号
	RightsNoType int64 `json:"rights_no_type,omitempty" xml:"rights_no_type,omitempty"`
	// 交易业务类型，1-B2C，2-C2C，3-寄租
	TradeBizType int64 `json:"trade_biz_type,omitempty" xml:"trade_biz_type,omitempty"`
	// 模版id
	AgreementId int64 `json:"agreement_id,omitempty" xml:"agreement_id,omitempty"`
	// 房东自然人主体id
	NaturalMerchantOpenId int64 `json:"natural_merchant_open_id,omitempty" xml:"natural_merchant_open_id,omitempty"`
	// 寄租模式下-企业B房东主体ID
	EnterpriseMerchantOpenId int64 `json:"enterprise_merchant_open_id,omitempty" xml:"enterprise_merchant_open_id,omitempty"`
	// 收款模式
	CollectionType int64 `json:"collection_type,omitempty" xml:"collection_type,omitempty"`
	// 寄租-房源房东类型 交易类型不为空，房源房东类型 1、企业B房东 2、自然人主体
	LandlordType int64 `json:"landlord_type,omitempty" xml:"landlord_type,omitempty"`
	// 代收主体id 当收款模式为代收是必填
	InsteadMerchantOpenId int64 `json:"instead_merchant_open_id,omitempty" xml:"instead_merchant_open_id,omitempty"`
}

var poolSyncExistingHouseEntrustDto = sync.Pool{
	New: func() any {
		return new(SyncExistingHouseEntrustDto)
	},
}

// GetSyncExistingHouseEntrustDto() 从对象池中获取SyncExistingHouseEntrustDto
func GetSyncExistingHouseEntrustDto() *SyncExistingHouseEntrustDto {
	return poolSyncExistingHouseEntrustDto.Get().(*SyncExistingHouseEntrustDto)
}

// ReleaseSyncExistingHouseEntrustDto 释放SyncExistingHouseEntrustDto
func ReleaseSyncExistingHouseEntrustDto(v *SyncExistingHouseEntrustDto) {
	v.RightsIdCards = v.RightsIdCards[:0]
	v.RightsNames = v.RightsNames[:0]
	v.EntrustPictures = v.EntrustPictures[:0]
	v.RentVouchers = v.RentVouchers[:0]
	v.HouseCheckPictureUrl = ""
	v.HouseCheckCode = ""
	v.EntrustEndTime = ""
	v.EntrustBeginTime = ""
	v.SaleRatioFee = ""
	v.EntrustPicture = ""
	v.CertificatePicture = ""
	v.EntrustStatusReason = ""
	v.HousePrice = ""
	v.OuterAgentId = ""
	v.OuterStoreId = ""
	v.OuterCompanyId = ""
	v.ConsignerPhone = ""
	v.ConsignerName = ""
	v.RightsNo = ""
	v.OuterId = ""
	v.HouseCheckDesc = ""
	v.OwnershipAddress = ""
	v.PayTypeJson = ""
	v.ServiceContent = ""
	v.RentRequirement = ""
	v.ViewingTime = ""
	v.CheckInTime = ""
	v.OuterBrandId = ""
	v.RentPrice = ""
	v.MaintenanceTime = ""
	v.HouseCheckInformation = ""
	v.MerchantOpenId = ""
	v.SendRentChargeModel = ""
	v.RoleList = ""
	v.HouseCheckStatus = 0
	v.EntrustStatus = 0
	v.Role = 0
	v.EntrustSellingId = 0
	v.MaxRentTime = 0
	v.MinRentTime = 0
	v.RightsNoType = 0
	v.TradeBizType = 0
	v.AgreementId = 0
	v.NaturalMerchantOpenId = 0
	v.EnterpriseMerchantOpenId = 0
	v.CollectionType = 0
	v.LandlordType = 0
	v.InsteadMerchantOpenId = 0
	poolSyncExistingHouseEntrustDto.Put(v)
}
