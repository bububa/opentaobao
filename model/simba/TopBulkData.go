package simba

import (
	"sync"
)

// TopBulkData 结构体
type TopBulkData struct {
	// 批量结果集
	TopMarketSceneVOList []TopMarketSceneVo `json:"top_market_scene_v_o_list,omitempty" xml:"top_market_scene_v_o_list>top_market_scene_vo,omitempty"`
	// 批量结果集
	AdgroupVOList []AdgroupVo `json:"adgroup_v_o_list,omitempty" xml:"adgroup_v_o_list>adgroup_vo,omitempty"`
	// 批量结果集
	AdzoneConfigVOList []AdzoneConfigVo `json:"adzone_config_v_o_list,omitempty" xml:"adzone_config_v_o_list>adzone_config_vo,omitempty"`
	// 批量结果集
	AdzoneRefVOList []AdzoneRefVo `json:"adzone_ref_v_o_list,omitempty" xml:"adzone_ref_v_o_list>adzone_ref_vo,omitempty"`
	// 批量结果集
	WordVOList []WordVo `json:"word_v_o_list,omitempty" xml:"word_v_o_list>word_vo,omitempty"`
	// 批量结果集
	BidwordSuggestItemVOList []BidwordSuggestItemVo `json:"bidword_suggest_item_v_o_list,omitempty" xml:"bidword_suggest_item_v_o_list>bidword_suggest_item_vo,omitempty"`
	// 批量结果集
	SuggestBidwordVOList []SuggestBidwordVo `json:"suggest_bidword_v_o_list,omitempty" xml:"suggest_bidword_v_o_list>suggest_bidword_vo,omitempty"`
	// 批量结果集
	CampaignVOList []CampaignVo `json:"campaign_v_o_list,omitempty" xml:"campaign_v_o_list>campaign_vo,omitempty"`
	// 批量结果集
	CampaignGroupVOList []CampaignGroupVo `json:"campaign_group_v_o_list,omitempty" xml:"campaign_group_v_o_list>campaign_group_vo,omitempty"`
	// 批量结果集
	CreativeRefVOList []CreativeRefVo `json:"creative_ref_v_o_list,omitempty" xml:"creative_ref_v_o_list>creative_ref_vo,omitempty"`
	// 批量结果集
	CreativeVOList []CreativeVo `json:"creative_v_o_list,omitempty" xml:"creative_v_o_list>creative_vo,omitempty"`
	// 批量结果集
	CrowdBindResultVOList []CrowdBindResultVo `json:"crowd_bind_result_v_o_list,omitempty" xml:"crowd_bind_result_v_o_list>crowd_bind_result_vo,omitempty"`
	// 批量结果集
	CrowdRefVOList []CrowdRefVo `json:"crowd_ref_v_o_list,omitempty" xml:"crowd_ref_v_o_list>crowd_ref_vo,omitempty"`
	// 批量结果集
	LabelConfigVOList []LabelConfigVo `json:"label_config_v_o_list,omitempty" xml:"label_config_v_o_list>label_config_vo,omitempty"`
	// 批量结果集
	MaterialAccessAllowVOList []MaterialAccessAllowVo `json:"material_access_allow_v_o_list,omitempty" xml:"material_access_allow_v_o_list>material_access_allow_vo,omitempty"`
	// 批量结果集
	ItemVOList []ItemVo `json:"item_v_o_list,omitempty" xml:"item_v_o_list>item_vo,omitempty"`
	// 批量结果集
	BrandInfoVOList []BrandInfoVo `json:"brand_info_v_o_list,omitempty" xml:"brand_info_v_o_list>brand_info_vo,omitempty"`
	// 批量结果集
	TopReportVOList []TopReportVo `json:"top_report_v_o_list,omitempty" xml:"top_report_v_o_list>top_report_vo,omitempty"`
	// 批量结果集
	ShopCategoryVOList []ShopCategoryVo `json:"shop_category_v_o_list,omitempty" xml:"shop_category_v_o_list>shop_category_vo,omitempty"`
	// 批量结果集
	StdCategoryVOList []StdCategoryVo `json:"std_category_v_o_list,omitempty" xml:"std_category_v_o_list>std_category_vo,omitempty"`
	// 批量结果集
	WordPackageVOList []WordPackageVo `json:"word_package_v_o_list,omitempty" xml:"word_package_v_o_list>word_package_vo,omitempty"`
	// 批量结果集
	WordPackageSuggestItemVOList []WordPackageSuggestItemVo `json:"word_package_suggest_item_v_o_list,omitempty" xml:"word_package_suggest_item_v_o_list>word_package_suggest_item_vo,omitempty"`
	// 批量结果集
	SuggestWordPackageVOList []SuggestWordPackageVo `json:"suggest_word_package_v_o_list,omitempty" xml:"suggest_word_package_v_o_list>suggest_word_package_vo,omitempty"`
	// 数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}

var poolTopBulkData = sync.Pool{
	New: func() any {
		return new(TopBulkData)
	},
}

// GetTopBulkData() 从对象池中获取TopBulkData
func GetTopBulkData() *TopBulkData {
	return poolTopBulkData.Get().(*TopBulkData)
}

// ReleaseTopBulkData 释放TopBulkData
func ReleaseTopBulkData(v *TopBulkData) {
	v.TopMarketSceneVOList = v.TopMarketSceneVOList[:0]
	v.AdgroupVOList = v.AdgroupVOList[:0]
	v.AdzoneConfigVOList = v.AdzoneConfigVOList[:0]
	v.AdzoneRefVOList = v.AdzoneRefVOList[:0]
	v.WordVOList = v.WordVOList[:0]
	v.BidwordSuggestItemVOList = v.BidwordSuggestItemVOList[:0]
	v.SuggestBidwordVOList = v.SuggestBidwordVOList[:0]
	v.CampaignVOList = v.CampaignVOList[:0]
	v.CampaignGroupVOList = v.CampaignGroupVOList[:0]
	v.CreativeRefVOList = v.CreativeRefVOList[:0]
	v.CreativeVOList = v.CreativeVOList[:0]
	v.CrowdBindResultVOList = v.CrowdBindResultVOList[:0]
	v.CrowdRefVOList = v.CrowdRefVOList[:0]
	v.LabelConfigVOList = v.LabelConfigVOList[:0]
	v.MaterialAccessAllowVOList = v.MaterialAccessAllowVOList[:0]
	v.ItemVOList = v.ItemVOList[:0]
	v.BrandInfoVOList = v.BrandInfoVOList[:0]
	v.TopReportVOList = v.TopReportVOList[:0]
	v.ShopCategoryVOList = v.ShopCategoryVOList[:0]
	v.StdCategoryVOList = v.StdCategoryVOList[:0]
	v.WordPackageVOList = v.WordPackageVOList[:0]
	v.WordPackageSuggestItemVOList = v.WordPackageSuggestItemVOList[:0]
	v.SuggestWordPackageVOList = v.SuggestWordPackageVOList[:0]
	v.Count = 0
	poolTopBulkData.Put(v)
}
