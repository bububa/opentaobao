package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelItemBaseAddAPIRequest
【API3.0】度假线路商品发布接口 API请求
taobao.alitrip.travel.item.base.add

旅行度假新商品发布接口。目前支持的类目包括：境内跟团游、出境跟团游、境内自由行、出境自由行、境内当地玩乐、境外玩乐套餐、境内邮轮、国际邮轮 */
type TaobaoAlitripTravelItemBaseAddAPIRequest struct {
	model.Params
	// 必填，商品基本信息
	_baseInfo *PontusTravelItemBaseInfo
	// 必填，预定规则结构。示例： [{  "rule_type": "fee_excluded",  "rule_desc": "费用包含描述"},{  "rule_type": "fee_included",  "rule_desc": "费用不含描述"},{  "rule_type": "order_info",  "rule_desc": "预定须知描述"}]
	_bookingRules []PontusTravelBookingRuleInfo
	// 特殊选填（当发布自由行商品时必填，其他情况不填）。自由行商品相关信息，自由行商品必填。大交通、酒店住宿、景点门票、租车、保险、餐饮、其他费用，自由行商品至少需要填写其中两项
	_freedomItemExt *PontusTravelFreedomItemExt
	// 特殊选填（当发布跟团游商品时必填，其他情况不填）。跟团游商品相关信息，发布跟团游商品时必填
	_groupItemExt *PontusTravelGroupItemExt
	// 选填，详细行程描述结构
	_itineraries []PontusTravelItemItineraryInfo
	// 选填，退款规则结构
	_refundInfo *PontusTravelItemRefundInfo
	// 选填，商品的销售属性相关信息
	_salesInfo *PontusTravelItemSaleInfo
	// 特殊选填（当发布邮轮商品时必填，其他情况不填）邮轮商品相关信息，发布邮轮商品时必填
	_cruiseItemExt *CruiseItemExt
	// 特殊选填（当发布同城活动商品时必填，其他情况不填）同城活动商品相关信息，发布同城活动商品时必填
	_tcwlItemExt *TcwlItemExt
}

// New
