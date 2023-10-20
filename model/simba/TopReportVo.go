package simba

import (
	"sync"
)

// TopReportVo 结构体
type TopReportVo struct {
	// 场景code
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 场景名称
	Scene1Name string `json:"scene1_name,omitempty" xml:"scene1_name,omitempty"`
	// 日期
	Thedate string `json:"thedate,omitempty" xml:"thedate,omitempty"`
	// 计划id
	CampaignId string `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 计划名称
	CampaignName string `json:"campaign_name,omitempty" xml:"campaign_name,omitempty"`
	// 单元id
	AdgroupId string `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
	// 单元名称
	AdgroupName string `json:"adgroup_name,omitempty" xml:"adgroup_name,omitempty"`
	// 关键词id
	BidwordId string `json:"bidword_id,omitempty" xml:"bidword_id,omitempty"`
	// 关键词/词包名称
	OriginalWord string `json:"original_word,omitempty" xml:"original_word,omitempty"`
	// 关键词包id
	BidwordPkgId string `json:"bidword_pkg_id,omitempty" xml:"bidword_pkg_id,omitempty"`
	// 词的类型
	BidWordType string `json:"bid_word_type,omitempty" xml:"bid_word_type,omitempty"`
	// 推广宝贝id
	PromotionId string `json:"promotion_id,omitempty" xml:"promotion_id,omitempty"`
	// 推广宝贝名称
	PromotionName string `json:"promotion_name,omitempty" xml:"promotion_name,omitempty"`
	// 推广宝贝图片地址
	ImgUrl string `json:"img_url,omitempty" xml:"img_url,omitempty"`
	// 推广宝贝跳转地址
	LinkUrl string `json:"link_url,omitempty" xml:"link_url,omitempty"`
	// 人群名称
	CrowdName string `json:"crowd_name,omitempty" xml:"crowd_name,omitempty"`
	// 创意id
	CreativeId string `json:"creative_id,omitempty" xml:"creative_id,omitempty"`
	// 创意名称
	CreativeName string `json:"creative_name,omitempty" xml:"creative_name,omitempty"`
	// 创意图片地址
	CreativeImagePath string `json:"creative_image_path,omitempty" xml:"creative_image_path,omitempty"`
	// 创意视频地址
	CreativeVedioPath string `json:"creative_vedio_path,omitempty" xml:"creative_vedio_path,omitempty"`
	// 创意尺寸
	CreativeSize string `json:"creative_size,omitempty" xml:"creative_size,omitempty"`
	// 黑盒创意的 宝贝id，当创意id为 -1 时使用
	BlackCreativePromotionId string `json:"black_creative_promotion_id,omitempty" xml:"black_creative_promotion_id,omitempty"`
	// 黑盒创意 宝贝主体类型，当创意id为 -1 时使用
	BlackCreativePromotionType string `json:"black_creative_promotion_type,omitempty" xml:"black_creative_promotion_type,omitempty"`
	// 省份id
	ProvinceId string `json:"province_id,omitempty" xml:"province_id,omitempty"`
	// 省份名称
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
	// 城市id
	CityId string `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 城市名称
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 人群id
	BpCrowdId string `json:"bp_crowd_id,omitempty" xml:"bp_crowd_id,omitempty"`
	// 小时ID
	HourId string `json:"hour_id,omitempty" xml:"hour_id,omitempty"`
	// 报表指标集合对象
	ReportIndex *TopReportIndexVo `json:"report_index,omitempty" xml:"report_index,omitempty"`
}

var poolTopReportVo = sync.Pool{
	New: func() any {
		return new(TopReportVo)
	},
}

// GetTopReportVo() 从对象池中获取TopReportVo
func GetTopReportVo() *TopReportVo {
	return poolTopReportVo.Get().(*TopReportVo)
}

// ReleaseTopReportVo 释放TopReportVo
func ReleaseTopReportVo(v *TopReportVo) {
	v.BizCode = ""
	v.Scene1Name = ""
	v.Thedate = ""
	v.CampaignId = ""
	v.CampaignName = ""
	v.AdgroupId = ""
	v.AdgroupName = ""
	v.BidwordId = ""
	v.OriginalWord = ""
	v.BidwordPkgId = ""
	v.BidWordType = ""
	v.PromotionId = ""
	v.PromotionName = ""
	v.ImgUrl = ""
	v.LinkUrl = ""
	v.CrowdName = ""
	v.CreativeId = ""
	v.CreativeName = ""
	v.CreativeImagePath = ""
	v.CreativeVedioPath = ""
	v.CreativeSize = ""
	v.BlackCreativePromotionId = ""
	v.BlackCreativePromotionType = ""
	v.ProvinceId = ""
	v.ProvinceName = ""
	v.CityId = ""
	v.CityName = ""
	v.BpCrowdId = ""
	v.HourId = ""
	v.ReportIndex = nil
	poolTopReportVo.Put(v)
}
