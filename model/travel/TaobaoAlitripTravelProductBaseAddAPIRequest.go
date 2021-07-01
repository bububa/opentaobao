package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelProductBaseAddAPIRequest
供应商新增产品API API请求
taobao.alitrip.travel.product.base.add

飞猪供销平台供应商可通过该API发布新产品 */
type TaobaoAlitripTravelProductBaseAddAPIRequest struct {
	model.Params
	// 产品基本信息
	_baseInfo *ProductBaseInfo
	// 选填，详细行程描述结构
	_itineraries []ItemItineraryInfo
	// 选填，退款规则结构
	_refundInfo *ItemRefundInfo
	// 必填，预定规则结构。示例： [{  "rule_type": "fee_excluded",  "rule_desc": "费用包含描述"},{  "rule_type": "fee_included",  "rule_desc": "费用不含描述"},{  "rule_type": "order_info",  "rule_desc": "预定须知描述"}]
	_bookingRules []BookingRuleInfo
	// 特殊选填（当发布邮轮商品时必填，其他情况不填）邮轮商品相关信息，发布邮轮商品时必填
	_cruiseProductExt *CruiseProductExt
	// 选填，商品的销售属性相关信息
	_productSaleInfo *ProductSaleInfo
}

// New
