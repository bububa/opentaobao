package alitrippoi

// TripPoiRawSaveParam 结构体
type TripPoiRawSaveParam struct {
	// 国家名称
	CountryName string `json:"country_name,omitempty" xml:"country_name,omitempty"`
	// 当地poi名称
	LocalName string `json:"local_name,omitempty" xml:"local_name,omitempty"`
	// 备用号码
	AlternativePhone string `json:"alternative_phone,omitempty" xml:"alternative_phone,omitempty"`
	// poi英文名称
	NameEn string `json:"name_en,omitempty" xml:"name_en,omitempty"`
	// 视频url
	VideoUrl string `json:"video_url,omitempty" xml:"video_url,omitempty"`
	// 国家码
	CountryCode string `json:"country_code,omitempty" xml:"country_code,omitempty"`
	// 经度
	Lng string `json:"lng,omitempty" xml:"lng,omitempty"`
	// 简介
	Bios string `json:"bios,omitempty" xml:"bios,omitempty"`
	// 城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 营业时间
	BusinessHour string `json:"business_hour,omitempty" xml:"business_hour,omitempty"`
	// 类型
	Category string `json:"category,omitempty" xml:"category,omitempty"`
	// 外部源唯一id
	SourceBizId string `json:"source_biz_id,omitempty" xml:"source_biz_id,omitempty"`
	// 扩展字段
	ExtendMap string `json:"extend_map,omitempty" xml:"extend_map,omitempty"`
	// 邮政编码
	PostalCode string `json:"postal_code,omitempty" xml:"postal_code,omitempty"`
	// 主要电话
	MainPhone string `json:"main_phone,omitempty" xml:"main_phone,omitempty"`
	// 来源
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 当地语言
	LocalLanguage string `json:"local_language,omitempty" xml:"local_language,omitempty"`
	// 描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// poi名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// poi状态
	OpenStatus int64 `json:"open_status,omitempty" xml:"open_status,omitempty"`
	// 图片urls
	PhotoUrls []string `json:"photo_urls,omitempty" xml:"photo_urls>string,omitempty"`
	// 外部网站url
	WebSiteUrl string `json:"web_site_url,omitempty" xml:"web_site_url,omitempty"`
	// 纬度
	Lat string `json:"lat,omitempty" xml:"lat,omitempty"`
}
