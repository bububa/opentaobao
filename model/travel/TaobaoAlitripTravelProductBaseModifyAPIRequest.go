package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelproductbasemodifyAPIRequest 供应商编辑产品API API请求
// taobao.alitrip.travel.product.base.modify
//
// 飞猪供销平台供应商可通过该API编辑产品
type TaobaoalitriptravelproductbasemodifyAPIRequest struct {
	model.Params
	// 详细行程描述结构
	_itineraries []ItemItineraryInfo
	// 预定规则结构。示例： [{ "rule_type": "fee_excluded", "rule_desc": "费用包含描述"},{ "rule_type": "fee_included", "rule_desc": "费用不含描述"},{ "rule_type": "order_info", "rule_desc": "预定须知描述"}]
	_bookingRules []BookingRuleInfo
	// 产品基本信息
	_baseInfo *ProductBaseInfo
	// 产品id
	_itemId int64
	// 退款规则结构
	_refundInfo *ItemRefundInfo
	// 邮轮商品相关信息
	_cruiseProductExt *CruiseProductExt
	// 商品的销售属性相关信息
	_productSaleInfo *ProductSaleInfo
}

// NewTaobaoalitriptravelproductbasemodifyRequest 初始化TaobaoalitriptravelproductbasemodifyAPIRequest对象
func NewTaobaoalitriptravelproductbasemodifyRequest() *TaobaoalitriptravelproductbasemodifyAPIRequest {
	return &TaobaoalitriptravelproductbasemodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelproductbasemodifyAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.product.base.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelproductbasemodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelproductbasemodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItineraries is Itineraries Setter
// 详细行程描述结构
func (r *TaobaoalitriptravelproductbasemodifyAPIRequest) SetItineraries(_itineraries []ItemItineraryInfo) error {
	r._itineraries = _itineraries
	r.Set("itineraries", _itineraries)
	return nil
}

// GetItineraries Itineraries Getter
func (r TaobaoalitriptravelproductbasemodifyAPIRequest) GetItineraries() []ItemItineraryInfo {
	return r._itineraries
}

// SetBookingRules is BookingRules Setter
// 预定规则结构。示例： [{ &#34;rule_type&#34;: &#34;fee_excluded&#34;, &#34;rule_desc&#34;: &#34;费用包含描述&#34;},{ &#34;rule_type&#34;: &#34;fee_included&#34;, &#34;rule_desc&#34;: &#34;费用不含描述&#34;},{ &#34;rule_type&#34;: &#34;order_info&#34;, &#34;rule_desc&#34;: &#34;预定须知描述&#34;}]
func (r *TaobaoalitriptravelproductbasemodifyAPIRequest) SetBookingRules(_bookingRules []BookingRuleInfo) error {
	r._bookingRules = _bookingRules
	r.Set("booking_rules", _bookingRules)
	return nil
}

// GetBookingRules BookingRules Getter
func (r TaobaoalitriptravelproductbasemodifyAPIRequest) GetBookingRules() []BookingRuleInfo {
	return r._bookingRules
}

// SetBaseInfo is BaseInfo Setter
// 产品基本信息
func (r *TaobaoalitriptravelproductbasemodifyAPIRequest) SetBaseInfo(_baseInfo *ProductBaseInfo) error {
	r._baseInfo = _baseInfo
	r.Set("base_info", _baseInfo)
	return nil
}

// GetBaseInfo BaseInfo Getter
func (r TaobaoalitriptravelproductbasemodifyAPIRequest) GetBaseInfo() *ProductBaseInfo {
	return r._baseInfo
}

// SetItemId is ItemId Setter
// 产品id
func (r *TaobaoalitriptravelproductbasemodifyAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoalitriptravelproductbasemodifyAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetRefundInfo is RefundInfo Setter
// 退款规则结构
func (r *TaobaoalitriptravelproductbasemodifyAPIRequest) SetRefundInfo(_refundInfo *ItemRefundInfo) error {
	r._refundInfo = _refundInfo
	r.Set("refund_info", _refundInfo)
	return nil
}

// GetRefundInfo RefundInfo Getter
func (r TaobaoalitriptravelproductbasemodifyAPIRequest) GetRefundInfo() *ItemRefundInfo {
	return r._refundInfo
}

// SetCruiseProductExt is CruiseProductExt Setter
// 邮轮商品相关信息
func (r *TaobaoalitriptravelproductbasemodifyAPIRequest) SetCruiseProductExt(_cruiseProductExt *CruiseProductExt) error {
	r._cruiseProductExt = _cruiseProductExt
	r.Set("cruise_product_ext", _cruiseProductExt)
	return nil
}

// GetCruiseProductExt CruiseProductExt Getter
func (r TaobaoalitriptravelproductbasemodifyAPIRequest) GetCruiseProductExt() *CruiseProductExt {
	return r._cruiseProductExt
}

// SetProductSaleInfo is ProductSaleInfo Setter
// 商品的销售属性相关信息
func (r *TaobaoalitriptravelproductbasemodifyAPIRequest) SetProductSaleInfo(_productSaleInfo *ProductSaleInfo) error {
	r._productSaleInfo = _productSaleInfo
	r.Set("product_sale_info", _productSaleInfo)
	return nil
}

// GetProductSaleInfo ProductSaleInfo Getter
func (r TaobaoalitriptravelproductbasemodifyAPIRequest) GetProductSaleInfo() *ProductSaleInfo {
	return r._productSaleInfo
}
