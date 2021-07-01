package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelProductBaseModifyAPIRequest
供应商编辑产品API API请求
taobao.alitrip.travel.product.base.modify

飞猪供销平台供应商可通过该API编辑产品 */
type TaobaoAlitripTravelProductBaseModifyAPIRequest struct {
	model.Params
	// 详细行程描述结构
	_itineraries []ItemItineraryInfo
	// 产品基本信息
	_baseInfo *ProductBaseInfo
	// 产品id
	_itemId int64
	// 退款规则结构
	_refundInfo *ItemRefundInfo
	// 预定规则结构。示例： [{ "rule_type": "fee_excluded", "rule_desc": "费用包含描述"},{ "rule_type": "fee_included", "rule_desc": "费用不含描述"},{ "rule_type": "order_info", "rule_desc": "预定须知描述"}]
	_bookingRules []BookingRuleInfo
	// 邮轮商品相关信息
	_cruiseProductExt *CruiseProductExt
	// 商品的销售属性相关信息
	_productSaleInfo *ProductSaleInfo
}

// New
