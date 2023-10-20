package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelproductbaseaddAPIRequest 供应商新增产品API API请求
// taobao.alitrip.travel.product.base.add
//
// 飞猪供销平台供应商可通过该API发布新产品
type TaobaoalitriptravelproductbaseaddAPIRequest struct {
	model.Params
	// 选填，详细行程描述结构
	_itineraries []ItemItineraryInfo
	// 必填，预定规则结构。示例： [{  "rule_type": "fee_excluded",  "rule_desc": "费用包含描述"},{  "rule_type": "fee_included",  "rule_desc": "费用不含描述"},{  "rule_type": "order_info",  "rule_desc": "预定须知描述"}]
	_bookingRules []BookingRuleInfo
	// 产品基本信息
	_baseInfo *ProductBaseInfo
	// 选填，退款规则结构
	_refundInfo *ItemRefundInfo
	// 特殊选填（当发布邮轮商品时必填，其他情况不填）邮轮商品相关信息，发布邮轮商品时必填
	_cruiseProductExt *CruiseProductExt
	// 选填，商品的销售属性相关信息
	_productSaleInfo *ProductSaleInfo
}

// NewTaobaoalitriptravelproductbaseaddRequest 初始化TaobaoalitriptravelproductbaseaddAPIRequest对象
func NewTaobaoalitriptravelproductbaseaddRequest() *TaobaoalitriptravelproductbaseaddAPIRequest {
	return &TaobaoalitriptravelproductbaseaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelproductbaseaddAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.product.base.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelproductbaseaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelproductbaseaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItineraries is Itineraries Setter
// 选填，详细行程描述结构
func (r *TaobaoalitriptravelproductbaseaddAPIRequest) SetItineraries(_itineraries []ItemItineraryInfo) error {
	r._itineraries = _itineraries
	r.Set("itineraries", _itineraries)
	return nil
}

// GetItineraries Itineraries Getter
func (r TaobaoalitriptravelproductbaseaddAPIRequest) GetItineraries() []ItemItineraryInfo {
	return r._itineraries
}

// SetBookingRules is BookingRules Setter
// 必填，预定规则结构。示例： [{  &#34;rule_type&#34;: &#34;fee_excluded&#34;,  &#34;rule_desc&#34;: &#34;费用包含描述&#34;},{  &#34;rule_type&#34;: &#34;fee_included&#34;,  &#34;rule_desc&#34;: &#34;费用不含描述&#34;},{  &#34;rule_type&#34;: &#34;order_info&#34;,  &#34;rule_desc&#34;: &#34;预定须知描述&#34;}]
func (r *TaobaoalitriptravelproductbaseaddAPIRequest) SetBookingRules(_bookingRules []BookingRuleInfo) error {
	r._bookingRules = _bookingRules
	r.Set("booking_rules", _bookingRules)
	return nil
}

// GetBookingRules BookingRules Getter
func (r TaobaoalitriptravelproductbaseaddAPIRequest) GetBookingRules() []BookingRuleInfo {
	return r._bookingRules
}

// SetBaseInfo is BaseInfo Setter
// 产品基本信息
func (r *TaobaoalitriptravelproductbaseaddAPIRequest) SetBaseInfo(_baseInfo *ProductBaseInfo) error {
	r._baseInfo = _baseInfo
	r.Set("base_info", _baseInfo)
	return nil
}

// GetBaseInfo BaseInfo Getter
func (r TaobaoalitriptravelproductbaseaddAPIRequest) GetBaseInfo() *ProductBaseInfo {
	return r._baseInfo
}

// SetRefundInfo is RefundInfo Setter
// 选填，退款规则结构
func (r *TaobaoalitriptravelproductbaseaddAPIRequest) SetRefundInfo(_refundInfo *ItemRefundInfo) error {
	r._refundInfo = _refundInfo
	r.Set("refund_info", _refundInfo)
	return nil
}

// GetRefundInfo RefundInfo Getter
func (r TaobaoalitriptravelproductbaseaddAPIRequest) GetRefundInfo() *ItemRefundInfo {
	return r._refundInfo
}

// SetCruiseProductExt is CruiseProductExt Setter
// 特殊选填（当发布邮轮商品时必填，其他情况不填）邮轮商品相关信息，发布邮轮商品时必填
func (r *TaobaoalitriptravelproductbaseaddAPIRequest) SetCruiseProductExt(_cruiseProductExt *CruiseProductExt) error {
	r._cruiseProductExt = _cruiseProductExt
	r.Set("cruise_product_ext", _cruiseProductExt)
	return nil
}

// GetCruiseProductExt CruiseProductExt Getter
func (r TaobaoalitriptravelproductbaseaddAPIRequest) GetCruiseProductExt() *CruiseProductExt {
	return r._cruiseProductExt
}

// SetProductSaleInfo is ProductSaleInfo Setter
// 选填，商品的销售属性相关信息
func (r *TaobaoalitriptravelproductbaseaddAPIRequest) SetProductSaleInfo(_productSaleInfo *ProductSaleInfo) error {
	r._productSaleInfo = _productSaleInfo
	r.Set("product_sale_info", _productSaleInfo)
	return nil
}

// GetProductSaleInfo ProductSaleInfo Getter
func (r TaobaoalitriptravelproductbaseaddAPIRequest) GetProductSaleInfo() *ProductSaleInfo {
	return r._productSaleInfo
}
