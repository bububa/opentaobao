package simba

import (
	"sync"
)

// CampaignVo 结构体
type CampaignVo struct {
	// 主体类型，支持一个计划下存在多种类型
	PromotionTypeList []string `json:"promotion_type_list,omitempty" xml:"promotion_type_list>string,omitempty"`
	// 子主体类型，支持一个计划下存在多种类型
	SubPromotionTypeList []string `json:"sub_promotion_type_list,omitempty" xml:"sub_promotion_type_list>string,omitempty"`
	// 地域码
	LaunchAreaStrList []string `json:"launch_area_str_list,omitempty" xml:"launch_area_str_list>string,omitempty"`
	// 投放折扣时段设置
	LaunchPeriodList []LaunchPeriodVo `json:"launch_period_list,omitempty" xml:"launch_period_list>launch_period_vo,omitempty"`
	// 人群过滤
	CrowdFilterList []CampaignCrowdFilterVo `json:"crowd_filter_list,omitempty" xml:"crowd_filter_list>campaign_crowd_filter_vo,omitempty"`
	// 品牌过滤
	BrandFilterList []CampaignBrandFilterVo `json:"brand_filter_list,omitempty" xml:"brand_filter_list>campaign_brand_filter_vo,omitempty"`
	// 品牌正向过滤
	DeeplinkBrandList []DeeplinkBrandVo `json:"deeplink_brand_list,omitempty" xml:"deeplink_brand_list>deeplink_brand_vo,omitempty"`
	// 性别年龄过滤
	GenderAgeFilterList []CampaignGenderAgeFilterVo `json:"gender_age_filter_list,omitempty" xml:"gender_age_filter_list>campaign_gender_age_filter_vo,omitempty"`
	// 自动化选品时屏蔽宝贝
	ShieldItems []int64 `json:"shield_items,omitempty" xml:"shield_items>int64,omitempty"`
	// 宝贝优选里面的指定宝贝
	ScopeItems []int64 `json:"scope_items,omitempty" xml:"scope_items>int64,omitempty"`
	// 单元信息
	AdgroupList []AdgroupVo `json:"adgroup_list,omitempty" xml:"adgroup_list>adgroup_vo,omitempty"`
	// 业务身份,枚举信息同&#39;account.get.can.use.bizcode&#39;服务返回结果
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 计划组名称
	CampaignGroupName string `json:"campaign_group_name,omitempty" xml:"campaign_group_name,omitempty"`
	// 计划名称
	CampaignName string `json:"campaign_name,omitempty" xml:"campaign_name,omitempty"`
	// 推广场景,promotion_scene_crowd:人群方舟,promotion_scene_item:店铺宝贝运营,promotion_scene_traffic:宝藏场景,AD_STRATEGY_LAXIN:策略中心拉新场景,AD_STRATEGY_SHANGXINLI:策略中心上新礼场景,AD_STRATEGY_FRESH_BOX:策略中心小黑盒场景,AD_STRATEGY_LAXIN_XINXIANG:策略中心拉新新享一笔钱场景,AD_STRATEGY_RUHUI_NEW:策略中心入会快新会员场景,AD_STRATEGY_RUHUI_OLD:策略中心入会快老会员场景,AD_STRATEGY_LIUZI:策略中心留资快场景,AD_STRATEGY_YURE:策略中心活动加速场景,AD_STRATEGY_HUODONG_ESCORT:活动全周期护航计划,AD_STRATEGY_CROWD_INDUSTRY:策略中心行业人群通,AD_STRATEGY_WHOLE_SHOP_SMART:策略中心全店放心推-智能托管,AD_STRATEGY_WHOLE_SHOP_CUSTOM:策略中心全店放心推-自定义,AD_STRATEGY_EVEN_BUDGET:策略中心均匀曝光场景,AD_STRATEGY_SXK_GUARANTEE:万相台上新快确定性保障,AD_STRATEGY_HPJS_GUARANTEE:万相台货品加速确定性保障,AD_STRATEGY_HPJS_TIME_LIMIT:万相台货品加速限时加速,AD_STRATEGY_FANS_HEADLINES:万相台粉丝头条,WXT_AGENCY_XST_CPC:服务商版一键起量（CPC）,WXT_AGENCY_XST_CPA:服务商版客源多（CPA）,WXT_AGENCY_SHEQUN:服务商版社群（CPA+CPS）
	PromotionScene string `json:"promotion_scene,omitempty" xml:"promotion_scene,omitempty"`
	// 推广目标,shop_item:全店宝贝推广,new_item_speed_car:新品飞车,shop_crowd:店铺人群运营,brand_crowd:品牌人群运营,traffic_shoujiao:淘系焦点图,traffic_nd:全屏微详情
	PromotionGoals string `json:"promotion_goals,omitempty" xml:"promotion_goals,omitempty"`
	// 推广类型
	PromotionModel string `json:"promotion_model,omitempty" xml:"promotion_model,omitempty"`
	// 推广主体类型,item:商品,item_private_mini:独享橱窗,shop:店铺,content:内容,short_video:短视频,user_define:自定义;
	PromotionType string `json:"promotion_type,omitempty" xml:"promotion_type,omitempty"`
	// 推广子主体类型,item:商品，item_private_mini:独享橱窗，shop:店铺，you_hao_huo:有好货，picture:图文，short_video:短视频，live_room:直播间，live_spot:看点，tao_blocks:淘积木，user_define_url:自定义url
	SubPromotionType string `json:"sub_promotion_type,omitempty" xml:"sub_promotion_type,omitempty"`
	// 优化目标，conv:促进成交,coll_cart:促进加购,click:促进点击,deal_count:促进成交笔数,exposure_pv:促进曝光,mini_view:促进橱窗宝贝浏览个数,mini_interactive:促进橱窗宝贝互动,ad_strategy_wangwang:策略中心旺旺咨询,shop_potential:提升潜客人数,shop_interest_new:提升兴趣新客人数,shop_purchase_new:提升首购新客人数,shop_visit_new:提升访问新客人数,shop_repurchase:提升复购人数,high_cvr:提升高转换人群成交人数,deeplink_d:提升品牌发现人数-D,deeplink_e1:提升品牌种草人数-E,deeplink_e2:提升品牌互动人数-E,deeplink_p:提升品牌行动人数-P,deeplink_i:提升品牌首购人数-I,deeplink_n:提升品牌复购人数-N,deeplink_k:提升品牌挚爱人数-K,hf_grass_plant:预热种草,hf_impoundment:预售蓄水,hf_harvest:爆发收割,nd_click:提高互动量,nd_cart:提高加购量,nd_deal:提高成交量,ad_strategy_ruhui_count:策略中心入会快新会员场景入会量目标,ad_strategy_lzl:策略中心留资快留资量目标,ad_strategy_try:策略中心派样量,ad_strategy_view:策略中心优化展现,ad_strategy_auto_trans:策略中心自动流转,ad_strategy_trial_order:策略中心表单获客成本,ad_strategy_liuzi_cost:策略中心广义留资目标,ad_strategy_cool_start:策略中心自动冷启动,ad_strategy_cool_start_mini_aim:策略中心冷启动分阶段目标流转,wxt_agency_ai:A转I人群流转,wxt_agency_smart:自定义场景,form_submit:表单提交,trial_order:试用下单,wangwang_liuzi:旺旺留资
	OptimizeTarget string `json:"optimize_target,omitempty" xml:"optimize_target,omitempty"`
	// 前端展示状态,pause:暂停推广,start:正在推广,terminate:结束推广,wait:等待推广,pending:准备推广,wait_pay:计划未付款
	DisplayStatus string `json:"display_status,omitempty" xml:"display_status,omitempty"`
	// 预算类型,normal:日预算,total:总预算,cycle:活动周期预算,day_freeze:日预算冻结,unlimit:不限预算
	DmcType string `json:"dmc_type,omitempty" xml:"dmc_type,omitempty"`
	// 日预算金额
	DayBudget string `json:"day_budget,omitempty" xml:"day_budget,omitempty"`
	// 日预算投放方式, quick:天内尽快,smooth:天内平滑
	SmoothOption string `json:"smooth_option,omitempty" xml:"smooth_option,omitempty"`
	// 周期预算金额
	TotalBudget string `json:"total_budget,omitempty" xml:"total_budget,omitempty"`
	// 周期预算投放方式, quick:天内尽快,smooth:天内平滑
	PeriodSmooth string `json:"period_smooth,omitempty" xml:"period_smooth,omitempty"`
	// 出价方式,custom_bid:手动出价,max_amount:最大化拿量,cost_control:控成本,roi_control:控投产比
	BidType string `json:"bid_type,omitempty" xml:"bid_type,omitempty"`
	// 手动出价时，最大金额
	MaxPrice string `json:"max_price,omitempty" xml:"max_price,omitempty"`
	// 手动出价时，最小金额
	MinPrice string `json:"min_price,omitempty" xml:"min_price,omitempty"`
	// 出价单位
	BidUnit string `json:"bid_unit,omitempty" xml:"bid_unit,omitempty"`
	// 出价约束类型,non:无约束,click:点击成本,coll_cart:收藏加购成本(非场景推广),cart:收藏加购成本(场景推广),conv:成交成本(非场景推广),deal:成交成本(场景推广),dir_conv:直接成交成本,roi:投产比约束
	ConstraintType string `json:"constraint_type,omitempty" xml:"constraint_type,omitempty"`
	// 出价约束值
	ConstraintValue string `json:"constraint_value,omitempty" xml:"constraint_value,omitempty"`
	// 冷启动时间
	ColdBootTime string `json:"cold_boot_time,omitempty" xml:"cold_boot_time,omitempty"`
	// 当前计划冷启动阶段，cold:冷启动期,mature:成熟期,acc_fail:冷启动失败,acc_success:冷启动成功,accelerate:加速中
	ColdBootStage string `json:"cold_boot_stage,omitempty" xml:"cold_boot_stage,omitempty"`
	// 分时折扣比例（取当前时刻对应的实际折扣，格式：80%）
	LaunchPeriodDiscount string `json:"launch_period_discount,omitempty" xml:"launch_period_discount,omitempty"`
	// 分时折扣展示时段（取当前时刻对应的半小时区间，格式：09:00-09:30）
	LaunchPeriodDisplayTime string `json:"launch_period_display_time,omitempty" xml:"launch_period_display_time,omitempty"`
	// 选品方式,shop:全店优选,user_define:用户自定义选品,scope:范围圈选,trend:趋势选品
	ItemSelectedMode string `json:"item_selected_mode,omitempty" xml:"item_selected_mode,omitempty"`
	// 优选集合,effect:效果优选,growing:全店成长,activity:活动选品,industry:行业尖货,scope:定义选品
	ShopItemType string `json:"shop_item_type,omitempty" xml:"shop_item_type,omitempty"`
	// 置顶时间
	TopTime string `json:"top_time,omitempty" xml:"top_time,omitempty"`
	// 计划组id
	CampaignGroupId int64 `json:"campaign_group_id,omitempty" xml:"campaign_group_id,omitempty"`
	// 计划id,计划已经存在场景必填,eg:添加主体、编辑计划状态等场景
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 投放状态,-1:未进入投放状态,0:用户设置状态-下线,1:用户设置状态-上线,2:合同未生效,9:投放结束,105:待支付
	OnlineStatus int64 `json:"online_status,omitempty" xml:"online_status,omitempty"`
	// 周期预算周期天数
	BudgetPeriod int64 `json:"budget_period,omitempty" xml:"budget_period,omitempty"`
	// ocpc出价
	CampaignOcpc *CampaignOcpcVo `json:"campaign_ocpc,omitempty" xml:"campaign_ocpc,omitempty"`
	// 冷启动标识，0:否，1:是
	ColdBoot int64 `json:"cold_boot,omitempty" xml:"cold_boot,omitempty"`
	// 计划投放时间
	LaunchTime *CampaignLaunchTimeVo `json:"launch_time,omitempty" xml:"launch_time,omitempty"`
	// 橱窗相关
	MiniDetail *CampaignMiniDetailVo `json:"mini_detail,omitempty" xml:"mini_detail,omitempty"`
	// 置顶状态,true:置顶，false:不置顶
	TopStatus bool `json:"top_status,omitempty" xml:"top_status,omitempty"`
}

var poolCampaignVo = sync.Pool{
	New: func() any {
		return new(CampaignVo)
	},
}

// GetCampaignVo() 从对象池中获取CampaignVo
func GetCampaignVo() *CampaignVo {
	return poolCampaignVo.Get().(*CampaignVo)
}

// ReleaseCampaignVo 释放CampaignVo
func ReleaseCampaignVo(v *CampaignVo) {
	v.PromotionTypeList = v.PromotionTypeList[:0]
	v.SubPromotionTypeList = v.SubPromotionTypeList[:0]
	v.LaunchAreaStrList = v.LaunchAreaStrList[:0]
	v.LaunchPeriodList = v.LaunchPeriodList[:0]
	v.CrowdFilterList = v.CrowdFilterList[:0]
	v.BrandFilterList = v.BrandFilterList[:0]
	v.DeeplinkBrandList = v.DeeplinkBrandList[:0]
	v.GenderAgeFilterList = v.GenderAgeFilterList[:0]
	v.ShieldItems = v.ShieldItems[:0]
	v.ScopeItems = v.ScopeItems[:0]
	v.AdgroupList = v.AdgroupList[:0]
	v.BizCode = ""
	v.CampaignGroupName = ""
	v.CampaignName = ""
	v.PromotionScene = ""
	v.PromotionGoals = ""
	v.PromotionModel = ""
	v.PromotionType = ""
	v.SubPromotionType = ""
	v.OptimizeTarget = ""
	v.DisplayStatus = ""
	v.DmcType = ""
	v.DayBudget = ""
	v.SmoothOption = ""
	v.TotalBudget = ""
	v.PeriodSmooth = ""
	v.BidType = ""
	v.MaxPrice = ""
	v.MinPrice = ""
	v.BidUnit = ""
	v.ConstraintType = ""
	v.ConstraintValue = ""
	v.ColdBootTime = ""
	v.ColdBootStage = ""
	v.LaunchPeriodDiscount = ""
	v.LaunchPeriodDisplayTime = ""
	v.ItemSelectedMode = ""
	v.ShopItemType = ""
	v.TopTime = ""
	v.CampaignGroupId = 0
	v.CampaignId = 0
	v.OnlineStatus = 0
	v.BudgetPeriod = 0
	v.CampaignOcpc = nil
	v.ColdBoot = 0
	v.LaunchTime = nil
	v.MiniDetail = nil
	v.TopStatus = false
	poolCampaignVo.Put(v)
}
