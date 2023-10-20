package alihouse

import (
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CompanyStoreForCheckDto 结构体
type CompanyStoreForCheckDto struct {
	// 外部门店id
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 外部公司id
	OuterCompanyId string `json:"outer_company_id,omitempty" xml:"outer_company_id,omitempty"`
	// 外部城市品牌店id
	OuterCompanyBrandId string `json:"outer_company_brand_id,omitempty" xml:"outer_company_brand_id,omitempty"`
	// 外部店铺id
	OuterShopId string `json:"outer_shop_id,omitempty" xml:"outer_shop_id,omitempty"`
	// 门店名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 门店简称
	StoreNameShort string `json:"store_name_short,omitempty" xml:"store_name_short,omitempty"`
	// 门店简介
	StoreInfo string `json:"store_info,omitempty" xml:"store_info,omitempty"`
	// 联系人
	ContactMan string `json:"contact_man,omitempty" xml:"contact_man,omitempty"`
	// 联系人电话
	ContactPhone string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// 400主号
	MainPhone string `json:"main_phone,omitempty" xml:"main_phone,omitempty"`
	// 子号
	SubPhone string `json:"sub_phone,omitempty" xml:"sub_phone,omitempty"`
	// 高德经度
	GaodeLongitude string `json:"gaode_longitude,omitempty" xml:"gaode_longitude,omitempty"`
	// 高德纬度
	GaodeLatitude string `json:"gaode_latitude,omitempty" xml:"gaode_latitude,omitempty"`
	// 地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 门店业务类型
	StoreBizType string `json:"store_biz_type,omitempty" xml:"store_biz_type,omitempty"`
	// 营业执照编号
	CompanyLicenseNo string `json:"company_license_no,omitempty" xml:"company_license_no,omitempty"`
	// 法人
	CompanyLegalPerson string `json:"company_legal_person,omitempty" xml:"company_legal_person,omitempty"`
	// 营业执照图片
	CompanyLicensePhoto string `json:"company_license_photo,omitempty" xml:"company_license_photo,omitempty"`
	// 营业执照过期时间
	CompanyLicenseExpireTime string `json:"company_license_expire_time,omitempty" xml:"company_license_expire_time,omitempty"`
	// 公司名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 营业执照备案地址
	CompanyLicenseUrl string `json:"company_license_url,omitempty" xml:"company_license_url,omitempty"`
	// 负责业务
	PartakeBusiness string `json:"partake_business,omitempty" xml:"partake_business,omitempty"`
	// 门店标签
	TagCodes string `json:"tag_codes,omitempty" xml:"tag_codes,omitempty"`
	// 城市id
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 门店状态
	StoreStatus int64 `json:"store_status,omitempty" xml:"store_status,omitempty"`
	// 营业执照状态
	CompanyLicenseStatus *model.File `json:"company_license_status,omitempty" xml:"company_license_status,omitempty"`
	// 是否为大ka
	IsSmallKa *model.File `json:"is_small_ka,omitempty" xml:"is_small_ka,omitempty"`
	// 是否为虚拟店铺
	VirtualType int64 `json:"virtual_type,omitempty" xml:"virtual_type,omitempty"`
	// 测试标
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
}

var poolCompanyStoreForCheckDto = sync.Pool{
	New: func() any {
		return new(CompanyStoreForCheckDto)
	},
}

// GetCompanyStoreForCheckDto() 从对象池中获取CompanyStoreForCheckDto
func GetCompanyStoreForCheckDto() *CompanyStoreForCheckDto {
	return poolCompanyStoreForCheckDto.Get().(*CompanyStoreForCheckDto)
}

// ReleaseCompanyStoreForCheckDto 释放CompanyStoreForCheckDto
func ReleaseCompanyStoreForCheckDto(v *CompanyStoreForCheckDto) {
	v.OuterStoreId = ""
	v.OuterCompanyId = ""
	v.OuterCompanyBrandId = ""
	v.OuterShopId = ""
	v.StoreName = ""
	v.StoreNameShort = ""
	v.StoreInfo = ""
	v.ContactMan = ""
	v.ContactPhone = ""
	v.MainPhone = ""
	v.SubPhone = ""
	v.GaodeLongitude = ""
	v.GaodeLatitude = ""
	v.Address = ""
	v.StoreBizType = ""
	v.CompanyLicenseNo = ""
	v.CompanyLegalPerson = ""
	v.CompanyLicensePhoto = ""
	v.CompanyLicenseExpireTime = ""
	v.CompanyName = ""
	v.CompanyLicenseUrl = ""
	v.PartakeBusiness = ""
	v.TagCodes = ""
	v.CityId = 0
	v.StoreStatus = 0
	v.CompanyLicenseStatus = nil
	v.IsSmallKa = nil
	v.VirtualType = 0
	v.IsTest = 0
	poolCompanyStoreForCheckDto.Put(v)
}
