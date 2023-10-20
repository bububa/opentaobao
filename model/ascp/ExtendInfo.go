package ascp

import (
	"sync"
)

// ExtendInfo 结构体
type ExtendInfo struct {
	// 合作快递公司编码列表；最多50个，创建时必填
	PlatformDeliveryCodes []string `json:"platform_delivery_codes,omitempty" xml:"platform_delivery_codes>string,omitempty"`
	// 优势品类：创建时必填 10=医药健康；11=宠物用品；12=家装家居；13=家清用品；14=母婴亲子；15=美妆个护；16=图书音像；17=服装服饰；18=箱包配件；19=汽车用品；20=电子产品；21=玩具潮玩；22=生鲜；23=运动户外；24=食品
	GoodsCategories []string `json:"goods_categories,omitempty" xml:"goods_categories>string,omitempty"`
	// 增值服务：创建时必填 1=商品有效期管理；2=冷链储运；3=包裹内物品整理；4=礼盒打包；5=商品定制服务；6=送货上门；7=订单仓内拦截；8=订单仓外拦截
	AdditionalServices []string `json:"additional_services,omitempty" xml:"additional_services>string,omitempty"`
	// 消防验收凭证链接（仅支持png、jpg、jpge格式，最多6张，每张大小不超过10M）：创建时必填
	FireSafetyAcceptances []string `json:"fire_safety_acceptances,omitempty" xml:"fire_safety_acceptances>string,omitempty"`
	// 租赁/产权证明链接（仅支持png、jpg、jpge格式，最多6张，每张大小不超过10M）：创建时必填
	OwnershipProofs []string `json:"ownership_proofs,omitempty" xml:"ownership_proofs>string,omitempty"`
	// 仓库/货品保险单链接（仅支持png、jpg、jpge格式，最多6张，每张大小不超过10M）
	WhInsurancePolicies []string `json:"wh_insurance_policies,omitempty" xml:"wh_insurance_policies>string,omitempty"`
	// 仓内图片链接（仅支持png、jpg、jpge格式，最多6张，每张大小不超过10M）：创建时必填
	WhInteriorPictures []string `json:"wh_interior_pictures,omitempty" xml:"wh_interior_pictures>string,omitempty"`
	// 营业执照链接（仅支持png、jpg、jpge格式，最多6张，每张大小不超过10M）：创建时必填
	BusinessLicenses []string `json:"business_licenses,omitempty" xml:"business_licenses>string,omitempty"`
	// 服务产品列表
	ServiceProducts []string `json:"service_products,omitempty" xml:"service_products>string,omitempty"`
	// 介绍视频链接（仅支持mov、mp4、avi格式，大小不超过20M建议上传仓内作业视频。且不能露出人脸）
	IntroductionVideo string `json:"introduction_video,omitempty" xml:"introduction_video,omitempty"`
	// 日均单量，创建时必填
	DailyOrderQuantity int64 `json:"daily_order_quantity,omitempty" xml:"daily_order_quantity,omitempty"`
	// 峰值出库能力，创建时必填 1：小于日常20倍 2：大于等于日常20倍小于50倍 3：大于等于日常50倍
	MaxOutboundCapacity int64 `json:"max_outbound_capacity,omitempty" xml:"max_outbound_capacity,omitempty"`
	// 仓储面积（m³），创建时必填
	StorageArea int64 `json:"storage_area,omitempty" xml:"storage_area,omitempty"`
	// B2C经验：创建时必填 1：小于1年 2：大于等于1，小于2年 3：大于等于2，小于5年 4：大于等于5年
	B2cExperience int64 `json:"b2c_experience,omitempty" xml:"b2c_experience,omitempty"`
	// 仓库类型：创建时必填 1=平库；2=楼库；3=高台库；4=立体仓库
	WarehouseType int64 `json:"warehouse_type,omitempty" xml:"warehouse_type,omitempty"`
	// 消防资质：创建时必填 1=甲类；2=乙一类；3=乙二类；4=丙一类；5=丙二类；6=丁类；7=戊类
	FireSafetyQualification int64 `json:"fire_safety_qualification,omitempty" xml:"fire_safety_qualification,omitempty"`
	// 租赁资质：创建时必填 1=租赁物业；2=自建物业
	LeaseQualification int64 `json:"lease_qualification,omitempty" xml:"lease_qualification,omitempty"`
	// 保险合同 1=基本险；2=综合险；3=一切险
	InsuranceContract int64 `json:"insurance_contract,omitempty" xml:"insurance_contract,omitempty"`
}

var poolExtendInfo = sync.Pool{
	New: func() any {
		return new(ExtendInfo)
	},
}

// GetExtendInfo() 从对象池中获取ExtendInfo
func GetExtendInfo() *ExtendInfo {
	return poolExtendInfo.Get().(*ExtendInfo)
}

// ReleaseExtendInfo 释放ExtendInfo
func ReleaseExtendInfo(v *ExtendInfo) {
	v.PlatformDeliveryCodes = v.PlatformDeliveryCodes[:0]
	v.GoodsCategories = v.GoodsCategories[:0]
	v.AdditionalServices = v.AdditionalServices[:0]
	v.FireSafetyAcceptances = v.FireSafetyAcceptances[:0]
	v.OwnershipProofs = v.OwnershipProofs[:0]
	v.WhInsurancePolicies = v.WhInsurancePolicies[:0]
	v.WhInteriorPictures = v.WhInteriorPictures[:0]
	v.BusinessLicenses = v.BusinessLicenses[:0]
	v.ServiceProducts = v.ServiceProducts[:0]
	v.IntroductionVideo = ""
	v.DailyOrderQuantity = 0
	v.MaxOutboundCapacity = 0
	v.StorageArea = 0
	v.B2cExperience = 0
	v.WarehouseType = 0
	v.FireSafetyQualification = 0
	v.LeaseQualification = 0
	v.InsuranceContract = 0
	poolExtendInfo.Put(v)
}
