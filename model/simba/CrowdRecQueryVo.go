package simba

import (
	"sync"
)

// CrowdRecQueryVo 结构体
type CrowdRecQueryVo struct {
	// 推荐场景,bd_sys:系统推荐人群,bd_ctr:行业榜单-点击率,bd_col_cart_rate:行业榜单-收藏加购,bd_cvr:行业榜单-转化率,bd_assist:行业榜单-助攻率,recSceneNew:关键词推广-新建流程,recSceneEffect:效果榜单,recSceneHeat:热度榜单,recSceneActivity:大促人群推荐,recSceneHomePage:首页榜单,recSceneAdgroupQualityOptions:单元列表智能拉新人群-优质组合推荐
	RecSceneList []string `json:"rec_scene_list,omitempty" xml:"rec_scene_list>string,omitempty"`
	// 宝贝id集合
	MaterialIdList []int64 `json:"material_id_list,omitempty" xml:"material_id_list>int64,omitempty"`
	// 选品方式,shop:全店优选,user_define:用户自定义选品,scope:范围圈选,trend:趋势选品
	ItemSelectedMode string `json:"item_selected_mode,omitempty" xml:"item_selected_mode,omitempty"`
	// 优选集合,effect:效果优选,growing:全店成长,activity:活动选品,industry:行业尖货,scope:定义选品
	ShopItemType string `json:"shop_item_type,omitempty" xml:"shop_item_type,omitempty"`
	// 推广主体类型
	PromotionType string `json:"promotion_type,omitempty" xml:"promotion_type,omitempty"`
	// 主体类型，支持一个计划下存在多种类型
	SubPromotionType string `json:"sub_promotion_type,omitempty" xml:"sub_promotion_type,omitempty"`
	// 优化目标，conv:促进成交,coll_cart:促进加购,click:促进点击,deal_count:促进成交笔数,exposure_pv:促进曝光,mini_view:促进橱窗宝贝浏览个数,mini_interactive:促进橱窗宝贝互动,ad_strategy_wangwang:策略中心旺旺咨询,shop_potential:提升潜客人数,shop_interest_new:提升兴趣新客人数,shop_purchase_new:提升首购新客人数,shop_visit_new:提升访问新客人数,shop_repurchase:提升复购人数,high_cvr:提升高转换人群成交人数,deeplink_d:提升品牌发现人数-D,deeplink_e1:提升品牌种草人数-E,deeplink_e2:提升品牌互动人数-E,deeplink_p:提升品牌行动人数-P,deeplink_i:提升品牌首购人数-I,deeplink_n:提升品牌复购人数-N,deeplink_k:提升品牌挚爱人数-K,hf_grass_plant:预热种草,hf_impoundment:预售蓄水,hf_harvest:爆发收割,nd_click:提高互动量,nd_cart:提高加购量,nd_deal:提高成交量,ad_strategy_ruhui_count:策略中心入会快新会员场景入会量目标,ad_strategy_lzl:策略中心留资快留资量目标,ad_strategy_try:策略中心派样量,ad_strategy_view:策略中心优化展现,ad_strategy_auto_trans:策略中心自动流转,ad_strategy_trial_order:策略中心表单获客成本,ad_strategy_liuzi_cost:策略中心广义留资目标,ad_strategy_cool_start:策略中心自动冷启动,ad_strategy_cool_start_mini_aim:策略中心冷启动分阶段目标流转,wxt_agency_ai:A转I人群流转,wxt_agency_smart:自定义场景,form_submit:表单提交,trial_order:试用下单,wangwang_liuzi:旺旺留资
	OptimizeTarget string `json:"optimize_target,omitempty" xml:"optimize_target,omitempty"`
	// 推广场景,promotion_scene_crowd:人群方舟,promotion_scene_item:店铺宝贝运营,promotion_scene_traffic:宝藏场景,AD_STRATEGY_LAXIN:策略中心拉新场景,AD_STRATEGY_SHANGXINLI:策略中心上新礼场景,AD_STRATEGY_FRESH_BOX:策略中心小黑盒场景,AD_STRATEGY_LAXIN_XINXIANG:策略中心拉新新享一笔钱场景,AD_STRATEGY_RUHUI_NEW:策略中心入会快新会员场景,AD_STRATEGY_RUHUI_OLD:策略中心入会快老会员场景,AD_STRATEGY_LIUZI:策略中心留资快场景,AD_STRATEGY_YURE:策略中心活动加速场景,AD_STRATEGY_HUODONG_ESCORT:活动全周期护航计划,AD_STRATEGY_CROWD_INDUSTRY:策略中心行业人群通,AD_STRATEGY_WHOLE_SHOP_SMART:策略中心全店放心推-智能托管,AD_STRATEGY_WHOLE_SHOP_CUSTOM:策略中心全店放心推-自定义,AD_STRATEGY_EVEN_BUDGET:策略中心均匀曝光场景,AD_STRATEGY_SXK_GUARANTEE:万相台上新快确定性保障,AD_STRATEGY_HPJS_GUARANTEE:万相台货品加速确定性保障,AD_STRATEGY_HPJS_TIME_LIMIT:万相台货品加速限时加速,AD_STRATEGY_FANS_HEADLINES:万相台粉丝头条,WXT_AGENCY_XST_CPC:服务商版一键起量（CPC）,WXT_AGENCY_XST_CPA:服务商版客源多（CPA）,WXT_AGENCY_SHEQUN:服务商版社群（CPA+CPS）
	PromotionScene string `json:"promotion_scene,omitempty" xml:"promotion_scene,omitempty"`
	// 计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 类目查询
	Category *CategoryQueryVo `json:"category,omitempty" xml:"category,omitempty"`
}

var poolCrowdRecQueryVo = sync.Pool{
	New: func() any {
		return new(CrowdRecQueryVo)
	},
}

// GetCrowdRecQueryVo() 从对象池中获取CrowdRecQueryVo
func GetCrowdRecQueryVo() *CrowdRecQueryVo {
	return poolCrowdRecQueryVo.Get().(*CrowdRecQueryVo)
}

// ReleaseCrowdRecQueryVo 释放CrowdRecQueryVo
func ReleaseCrowdRecQueryVo(v *CrowdRecQueryVo) {
	v.RecSceneList = v.RecSceneList[:0]
	v.MaterialIdList = v.MaterialIdList[:0]
	v.ItemSelectedMode = ""
	v.ShopItemType = ""
	v.PromotionType = ""
	v.SubPromotionType = ""
	v.OptimizeTarget = ""
	v.PromotionScene = ""
	v.CampaignId = 0
	v.Category = nil
	poolCrowdRecQueryVo.Put(v)
}
