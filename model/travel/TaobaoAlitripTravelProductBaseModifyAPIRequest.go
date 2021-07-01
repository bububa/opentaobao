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

// NewTaobaoAlitripTravelProductBaseModifyRequest 初始化TaobaoAlitripTravelProductBaseModifyAPIRequest对象
func NewTaobaoAlitripTravelProductBaseModifyRequest() *TaobaoAlitripTravelProductBaseModifyAPIRequest {
	return &TaobaoAlitripTravelProductBaseModifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelProductBaseModifyAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.product.base.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelProductBaseModifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Itineraries Setter
// 详细行程描述结构
func (r *TaobaoAlitripTravelProductBaseModifyAPIRequest) SetItineraries(_itineraries []ItemItineraryInfo) error {
	r._itineraries = _itineraries
	r.Set("itineraries", _itineraries)
	return nil
}

// Get Itineraries Getter
func (r TaobaoAlitripTravelProductBaseModifyAPIRequest) GetItineraries() []ItemItineraryInfo {
	return r._itineraries
}

// Set is BaseInfo Setter
// 产品基本信息
func (r *TaobaoAlitripTravelProductBaseModifyAPIRequest) SetBaseInfo(_baseInfo *ProductBaseInfo) error {
	r._baseInfo = _baseInfo
	r.Set("base_info", _baseInfo)
	return nil
}

// Get BaseInfo Getter
func (r TaobaoAlitripTravelProductBaseModifyAPIRequest) GetBaseInfo() *ProductBaseInfo {
	return r._baseInfo
}

// Set is ItemId Setter
// 产品id
func (r *TaobaoAlitripTravelProductBaseModifyAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TaobaoAlitripTravelProductBaseModifyAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is RefundInfo Setter
// 退款规则结构
func (r *TaobaoAlitripTravelProductBaseModifyAPIRequest) SetRefundInfo(_refundInfo *ItemRefundInfo) error {
	r._refundInfo = _refundInfo
	r.Set("refund_info", _refundInfo)
	return nil
}

// Get RefundInfo Getter
func (r TaobaoAlitripTravelProductBaseModifyAPIRequest) GetRefundInfo() *ItemRefundInfo {
	return r._refundInfo
}

// Set is BookingRules Setter
// 预定规则结构。示例： [{ "rule_type": "fee_excluded", "rule_desc": "费用包含描述"},{ "rule_type": "fee_included", "rule_desc": "费用不含描述"},{ "rule_type": "order_info", "rule_desc": "预定须知描述"}]
func (r *TaobaoAlitripTravelProductBaseModifyAPIRequest) SetBookingRules(_bookingRules []BookingRuleInfo) error {
	r._bookingRules = _bookingRules
	r.Set("booking_rules", _bookingRules)
	return nil
}

// Get BookingRules Getter
func (r TaobaoAlitripTravelProductBaseModifyAPIRequest) GetBookingRules() []BookingRuleInfo {
	return r._bookingRules
}

// Set is CruiseProductExt Setter
// 邮轮商品相关信息
func (r *TaobaoAlitripTravelProductBaseModifyAPIRequest) SetCruiseProductExt(_cruiseProductExt *CruiseProductExt) error {
	r._cruiseProductExt = _cruiseProductExt
	r.Set("cruise_product_ext", _cruiseProductExt)
	return nil
}

// Get CruiseProductExt Getter
func (r TaobaoAlitripTravelProductBaseModifyAPIRequest) GetCruiseProductExt() *CruiseProductExt {
	return r._cruiseProductExt
}

// Set is ProductSaleInfo Setter
// 商品的销售属性相关信息
func (r *TaobaoAlitripTravelProductBaseModifyAPIRequest) SetProductSaleInfo(_productSaleInfo *ProductSaleInfo) error {
	r._productSaleInfo = _productSaleInfo
	r.Set("product_sale_info", _productSaleInfo)
	return nil
}

// Get ProductSaleInfo Getter
func (r TaobaoAlitripTravelProductBaseModifyAPIRequest) GetProductSaleInfo() *ProductSaleInfo {
	return r._productSaleInfo
}
