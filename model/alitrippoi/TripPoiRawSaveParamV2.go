package alitrippoi

import (
	"sync"
)

// TripPoiRawSaveParamV2 结构体
type TripPoiRawSaveParamV2 struct {
	// 图片地址(,分隔多张)
	PhotoUrls []string `json:"photo_urls,omitempty" xml:"photo_urls>string,omitempty"`
	// 下线原因
	OfflineReason string `json:"offline_reason,omitempty" xml:"offline_reason,omitempty"`
	// 更新时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 主要电话
	MainPhone string `json:"main_phone,omitempty" xml:"main_phone,omitempty"`
	// 简介
	Bios string `json:"bios,omitempty" xml:"bios,omitempty"`
	// 邮政编码
	PostalCode string `json:"postal_code,omitempty" xml:"postal_code,omitempty"`
	// 描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 外部网站url
	WebSiteUrl string `json:"web_site_url,omitempty" xml:"web_site_url,omitempty"`
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 视频url
	VideoUrl string `json:"video_url,omitempty" xml:"video_url,omitempty"`
	// 商圈
	CommercialCircle string `json:"commercial_circle,omitempty" xml:"commercial_circle,omitempty"`
	// 国家码
	CountryCode string `json:"country_code,omitempty" xml:"country_code,omitempty"`
	// 扩展字段
	ExtendMap string `json:"extend_map,omitempty" xml:"extend_map,omitempty"`
	// 门店Id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 开放时间
	OpenTime string `json:"open_time,omitempty" xml:"open_time,omitempty"`
	// 英文地址
	AddressEn string `json:"address_en,omitempty" xml:"address_en,omitempty"`
	// 纬度
	Lat string `json:"lat,omitempty" xml:"lat,omitempty"`
	// 地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 外部源唯一id
	SourceBizId string `json:"source_biz_id,omitempty" xml:"source_biz_id,omitempty"`
	// 经度
	Lng string `json:"lng,omitempty" xml:"lng,omitempty"`
	// 别名
	NameAlias string `json:"name_alias,omitempty" xml:"name_alias,omitempty"`
	// 服务时间
	BusinessHour string `json:"business_hour,omitempty" xml:"business_hour,omitempty"`
	// 当地地址
	AddressLocal string `json:"address_local,omitempty" xml:"address_local,omitempty"`
	// 备用号码
	AlternativePhone string `json:"alternative_phone,omitempty" xml:"alternative_phone,omitempty"`
	// 手机号
	Telephone string `json:"telephone,omitempty" xml:"telephone,omitempty"`
	// 人均消费
	Consumption string `json:"consumption,omitempty" xml:"consumption,omitempty"`
	// 英文名
	NameEn string `json:"name_en,omitempty" xml:"name_en,omitempty"`
	// 交通
	Transport string `json:"transport,omitempty" xml:"transport,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 本地名
	LocalName string `json:"local_name,omitempty" xml:"local_name,omitempty"`
	// 推荐信息
	RecommendInfos string `json:"recommend_infos,omitempty" xml:"recommend_infos,omitempty"`
	// 名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 国家名
	CountryName string `json:"country_name,omitempty" xml:"country_name,omitempty"`
	// 店铺类型
	ShopType string `json:"shop_type,omitempty" xml:"shop_type,omitempty"`
	// 类型
	Category string `json:"category,omitempty" xml:"category,omitempty"`
	// 当地语言
	LocalLanguage string `json:"local_language,omitempty" xml:"local_language,omitempty"`
	// 下线详细原因
	OfflineReasonDetail string `json:"offline_reason_detail,omitempty" xml:"offline_reason_detail,omitempty"`
	// 服务详情
	ServiceInfo *StructureServiceInfo `json:"service_info,omitempty" xml:"service_info,omitempty"`
	// 类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 认领主账号id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 子账号id
	SubSellerId int64 `json:"sub_seller_id,omitempty" xml:"sub_seller_id,omitempty"`
	// poiId(飞猪端poiId)
	PoiId int64 `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	// 品牌信息
	BrandInfo *StructureBrandInfo `json:"brand_info,omitempty" xml:"brand_info,omitempty"`
	// poi状态
	OpenStatus int64 `json:"open_status,omitempty" xml:"open_status,omitempty"`
	// 中台门店id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 操作类型(0:新增 1:更新)
	OperatorType int64 `json:"operator_type,omitempty" xml:"operator_type,omitempty"`
}

var poolTripPoiRawSaveParamV2 = sync.Pool{
	New: func() any {
		return new(TripPoiRawSaveParamV2)
	},
}

// GetTripPoiRawSaveParamV2() 从对象池中获取TripPoiRawSaveParamV2
func GetTripPoiRawSaveParamV2() *TripPoiRawSaveParamV2 {
	return poolTripPoiRawSaveParamV2.Get().(*TripPoiRawSaveParamV2)
}

// ReleaseTripPoiRawSaveParamV2 释放TripPoiRawSaveParamV2
func ReleaseTripPoiRawSaveParamV2(v *TripPoiRawSaveParamV2) {
	v.PhotoUrls = v.PhotoUrls[:0]
	v.OfflineReason = ""
	v.GmtModified = ""
	v.City = ""
	v.MainPhone = ""
	v.Bios = ""
	v.PostalCode = ""
	v.Description = ""
	v.WebSiteUrl = ""
	v.Province = ""
	v.VideoUrl = ""
	v.CommercialCircle = ""
	v.CountryCode = ""
	v.ExtendMap = ""
	v.ShopId = ""
	v.OpenTime = ""
	v.AddressEn = ""
	v.Lat = ""
	v.Address = ""
	v.SourceBizId = ""
	v.Lng = ""
	v.NameAlias = ""
	v.BusinessHour = ""
	v.AddressLocal = ""
	v.AlternativePhone = ""
	v.Telephone = ""
	v.Consumption = ""
	v.NameEn = ""
	v.Transport = ""
	v.GmtCreate = ""
	v.LocalName = ""
	v.RecommendInfos = ""
	v.Name = ""
	v.CountryName = ""
	v.ShopType = ""
	v.Category = ""
	v.LocalLanguage = ""
	v.OfflineReasonDetail = ""
	v.ServiceInfo = nil
	v.Type = 0
	v.SellerId = 0
	v.SubSellerId = 0
	v.PoiId = 0
	v.BrandInfo = nil
	v.OpenStatus = 0
	v.StoreId = 0
	v.OperatorType = 0
	poolTripPoiRawSaveParamV2.Put(v)
}
