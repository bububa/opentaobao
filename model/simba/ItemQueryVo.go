package simba

import (
	"sync"
)

// ItemQueryVo 结构体
type ItemQueryVo struct {
	// 宝贝id集合
	ItemIdList []int64 `json:"item_id_list,omitempty" xml:"item_id_list>int64,omitempty"`
	// 宝贝名称，支持模糊查询
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 推广场景,promotion_scene_crowd:人群方舟,promotion_scene_item:店铺宝贝运营,promotion_scene_traffic:宝藏场景,AD_STRATEGY_LAXIN:策略中心拉新场景,AD_STRATEGY_SHANGXINLI:策略中心上新礼场景,AD_STRATEGY_FRESH_BOX:策略中心小黑盒场景,AD_STRATEGY_LAXIN_XINXIANG:策略中心拉新新享一笔钱场景,AD_STRATEGY_RUHUI_NEW:策略中心入会快新会员场景,AD_STRATEGY_RUHUI_OLD:策略中心入会快老会员场景,AD_STRATEGY_LIUZI:策略中心留资快场景,AD_STRATEGY_YURE:策略中心活动加速场景,AD_STRATEGY_HUODONG_ESCORT:活动全周期护航计划,AD_STRATEGY_CROWD_INDUSTRY:策略中心行业人群通,AD_STRATEGY_WHOLE_SHOP_SMART:策略中心全店放心推-智能托管,AD_STRATEGY_WHOLE_SHOP_CUSTOM:策略中心全店放心推-自定义,AD_STRATEGY_EVEN_BUDGET:策略中心均匀曝光场景,AD_STRATEGY_SXK_GUARANTEE:万相台上新快确定性保障,AD_STRATEGY_HPJS_GUARANTEE:万相台货品加速确定性保障,AD_STRATEGY_HPJS_TIME_LIMIT:万相台货品加速限时加速,AD_STRATEGY_FANS_HEADLINES:万相台粉丝头条,WXT_AGENCY_XST_CPC:服务商版一键起量（CPC）,WXT_AGENCY_XST_CPA:服务商版客源多（CPA）,WXT_AGENCY_SHEQUN:服务商版社群（CPA+CPS）
	PromotionScene string `json:"promotion_scene,omitempty" xml:"promotion_scene,omitempty"`
	// 优化目标，conv:促进成交,coll_cart:促进加购,click:促进点击,deal_count:促进成交笔数,exposure_pv:促进曝光,mini_view:促进橱窗宝贝浏览个数,mini_interactive:促进橱窗宝贝互动,ad_strategy_wangwang:策略中心旺旺咨询,shop_potential:提升潜客人数,shop_interest_new:提升兴趣新客人数,shop_purchase_new:提升首购新客人数,shop_visit_new:提升访问新客人数,shop_repurchase:提升复购人数,high_cvr:提升高转换人群成交人数,deeplink_d:提升品牌发现人数-D,deeplink_e1:提升品牌种草人数-E,deeplink_e2:提升品牌互动人数-E,deeplink_p:提升品牌行动人数-P,deeplink_i:提升品牌首购人数-I,deeplink_n:提升品牌复购人数-N,deeplink_k:提升品牌挚爱人数-K,hf_grass_plant:预热种草,hf_impoundment:预售蓄水,hf_harvest:爆发收割,nd_click:提高互动量,nd_cart:提高加购量,nd_deal:提高成交量,ad_strategy_ruhui_count:策略中心入会快新会员场景入会量目标,ad_strategy_lzl:策略中心留资快留资量目标,ad_strategy_try:策略中心派样量,ad_strategy_view:策略中心优化展现,ad_strategy_auto_trans:策略中心自动流转,ad_strategy_trial_order:策略中心表单获客成本,ad_strategy_liuzi_cost:策略中心广义留资目标,ad_strategy_cool_start:策略中心自动冷启动,ad_strategy_cool_start_mini_aim:策略中心冷启动分阶段目标流转,wxt_agency_ai:A转I人群流转,wxt_agency_smart:自定义场景,form_submit:表单提交,trial_order:试用下单,wangwang_liuzi:旺旺留资
	OptimizeTarget string `json:"optimize_target,omitempty" xml:"optimize_target,omitempty"`
	// 选品方式,shop:全店优选,user_define:用户自定义选品,scope:范围圈选,trend:趋势选品
	ItemSelectedMode string `json:"item_selected_mode,omitempty" xml:"item_selected_mode,omitempty"`
	// 优选集合,effect:效果优选,growing:全店成长,activity:活动选品,industry:行业尖货,scope:定义选品
	ShopItemType string `json:"shop_item_type,omitempty" xml:"shop_item_type,omitempty"`
	// 物料来源标识,1:全部宝贝,2:推荐宝贝,3:活动宝贝,4:有好货宝贝,5:图文,6:有好货帖子,7:直播,8:淘积木,9:短视频,10:图文手猜,11:可用津贴券商品,12:全部新品,13:618活动商品,14:可用店铺券宝贝,15:推荐新品,16:首页新品推荐,20:首页推荐,21:图文V任务,22:短视频V任务,23:直播V任务,25:店铺微淘,26:直播合作,24:落地页,27:活动商品,28:首页大促商品,29:投中探测爆款商品,30:已报名的新享宝贝,31:店铺视频,32:包含单元宝贝的淘积木列表,34:达人视频,33:不计入最低价新品,35:首猜精品库,36:首猜精品中的新品,37:行业优选,38:大促推荐新品,39:推广有礼&#34;:关于推广有礼,40:效果优选,41:潜力爆品,42:全店成长,43:双11,44:推广有礼,49:新品飞车数据,50:效果优选,53:全店优选二期,55:首单礼金,56:新品直降,57:限时优惠,58:618大促,59:扩容后新品飞车商品池,101:全部,102:优选宝贝&#34;:根据该宝贝历史数据预测为适合推广的宝贝。,103:优选流量&#34;:根据该宝贝历史数据预测为在引流方面有潜力宝贝。,104:优选转化&#34;:根据该宝贝历史数据预测为在转化方面有潜力宝贝。,105:生意参谋&#34;:根据店铺宝贝历史数据优选推荐的适合推广的宝贝。,106:生意参谋-物流红包,1000:组件tag开始index,1999:组件tag结束,209:策略中心货品加速单品,201:策略中心推荐宝贝,202:策略中心推荐相似宝贝,203:策略中心智能营销推荐宝贝,204:策略中心大快消TOP单品,205:策略中心新享报名商品,206:策略中心未报名商品,207:策略中心上新单品,214:策略中心上新礼,224:策略中心小黑盒上新,212:策略中心产业带单品,210:策略中心货品加速新的tag,211:策略中心货品加速版本2,213:策略中心大快消入会快,215:策略中心大快消留资快,216:策略中心大快消派样快,219:策略中心大快消预热蓄水,217:策略中心大快消爆发现货,218:策略中心大快消爆发预售,220:策略中心大快消测款,221:策略中心销量明星商品,222:策略中心测款均匀曝光,223:策略中心大测款均匀预算,225:策略中心上新数智经营争霸赛,226:策略中心货品加速数智经营争霸赛,227:策略中心数预热蓄水智经营争霸赛,228:策略中心爆发现货数智经营争霸赛,229:策略中心爆发预售数智经营争霸赛,230:策略中心大快消数智经营争霸赛,1016001:系统优选,1016002:相似宝贝,1100001:极速推宝贝推荐,101111301:冷启动新品,101111302:好货推荐,101111303:趋势主题选品
	TagId int64 `json:"tag_id,omitempty" xml:"tag_id,omitempty"`
	// 偏移量
	Offset int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// 页面大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolItemQueryVo = sync.Pool{
	New: func() any {
		return new(ItemQueryVo)
	},
}

// GetItemQueryVo() 从对象池中获取ItemQueryVo
func GetItemQueryVo() *ItemQueryVo {
	return poolItemQueryVo.Get().(*ItemQueryVo)
}

// ReleaseItemQueryVo 释放ItemQueryVo
func ReleaseItemQueryVo(v *ItemQueryVo) {
	v.ItemIdList = v.ItemIdList[:0]
	v.Title = ""
	v.PromotionScene = ""
	v.OptimizeTarget = ""
	v.ItemSelectedMode = ""
	v.ShopItemType = ""
	v.TagId = 0
	v.Offset = 0
	v.PageSize = 0
	poolItemQueryVo.Put(v)
}
