package travel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelProductBaseModifyAPIRequest 供应商编辑产品API API请求
// taobao.alitrip.travel.product.base.modify
//
// 飞猪供销平台供应商可通过该API编辑产品
type TaobaoAlitripTravelProductBaseModifyAPIRequest struct {
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

// NewTaobaoAlitripTravelProductBaseModifyRequest 初始化TaobaoAlitripTravelProductBaseModifyAPIRequest对象
func NewTaobaoAlitripTravelProductBaseModifyRequest() *TaobaoAlitripTravelProductBaseModifyAPIRequest {
	return &TaobaoAlitripTravelProductBaseModifyAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelProductBaseModifyAPIRequest) Reset() {
	r._itineraries = r._itineraries[:0]
	r._bookingRules = r._bookingRules[:0]
	r._baseInfo = nil
	r._itemId = 0
	r._refundInfo = nil
	r._cruiseProductExt = nil
	r._productSaleInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelProductBaseModifyAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.product.base.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelProductBaseModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelProductBaseModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItineraries is Itineraries Setter
// 详细行程描述结构
func (r *TaobaoAlitripTravelProductBaseModifyAPIRequest) SetItineraries(_itineraries []ItemItineraryInfo) error {
	r._itineraries = _itineraries
	r.Set("itineraries", _itineraries)
	return nil
}

// GetItineraries Itineraries Getter
func (r TaobaoAlitripTravelProductBaseModifyAPIRequest) GetItineraries() []ItemItineraryInfo {
	return r._itineraries
}

// SetBookingRules is BookingRules Setter
// 预定规则结构。示例： [{ &#34;rule_type&#34;: &#34;fee_excluded&#34;, &#34;rule_desc&#34;: &#34;费用包含描述&#34;},{ &#34;rule_type&#34;: &#34;fee_included&#34;, &#34;rule_desc&#34;: &#34;费用不含描述&#34;},{ &#34;rule_type&#34;: &#34;order_info&#34;, &#34;rule_desc&#34;: &#34;预定须知描述&#34;}]
func (r *TaobaoAlitripTravelProductBaseModifyAPIRequest) SetBookingRules(_bookingRules []BookingRuleInfo) error {
	r._bookingRules = _bookingRules
	r.Set("booking_rules", _bookingRules)
	return nil
}

// GetBookingRules BookingRules Getter
func (r TaobaoAlitripTravelProductBaseModifyAPIRequest) GetBookingRules() []BookingRuleInfo {
	return r._bookingRules
}

// SetBaseInfo is BaseInfo Setter
// 产品基本信息
func (r *TaobaoAlitripTravelProductBaseModifyAPIRequest) SetBaseInfo(_baseInfo *ProductBaseInfo) error {
	r._baseInfo = _baseInfo
	r.Set("base_info", _baseInfo)
	return nil
}

// GetBaseInfo BaseInfo Getter
func (r TaobaoAlitripTravelProductBaseModifyAPIRequest) GetBaseInfo() *ProductBaseInfo {
	return r._baseInfo
}

// SetItemId is ItemId Setter
// 产品id
func (r *TaobaoAlitripTravelProductBaseModifyAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoAlitripTravelProductBaseModifyAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetRefundInfo is RefundInfo Setter
// 退款规则结构
func (r *TaobaoAlitripTravelProductBaseModifyAPIRequest) SetRefundInfo(_refundInfo *ItemRefundInfo) error {
	r._refundInfo = _refundInfo
	r.Set("refund_info", _refundInfo)
	return nil
}

// GetRefundInfo RefundInfo Getter
func (r TaobaoAlitripTravelProductBaseModifyAPIRequest) GetRefundInfo() *ItemRefundInfo {
	return r._refundInfo
}

// SetCruiseProductExt is CruiseProductExt Setter
// 邮轮商品相关信息
func (r *TaobaoAlitripTravelProductBaseModifyAPIRequest) SetCruiseProductExt(_cruiseProductExt *CruiseProductExt) error {
	r._cruiseProductExt = _cruiseProductExt
	r.Set("cruise_product_ext", _cruiseProductExt)
	return nil
}

// GetCruiseProductExt CruiseProductExt Getter
func (r TaobaoAlitripTravelProductBaseModifyAPIRequest) GetCruiseProductExt() *CruiseProductExt {
	return r._cruiseProductExt
}

// SetProductSaleInfo is ProductSaleInfo Setter
// 商品的销售属性相关信息
func (r *TaobaoAlitripTravelProductBaseModifyAPIRequest) SetProductSaleInfo(_productSaleInfo *ProductSaleInfo) error {
	r._productSaleInfo = _productSaleInfo
	r.Set("product_sale_info", _productSaleInfo)
	return nil
}

// GetProductSaleInfo ProductSaleInfo Getter
func (r TaobaoAlitripTravelProductBaseModifyAPIRequest) GetProductSaleInfo() *ProductSaleInfo {
	return r._productSaleInfo
}

var poolTaobaoAlitripTravelProductBaseModifyAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelProductBaseModifyRequest()
	},
}

// GetTaobaoAlitripTravelProductBaseModifyRequest 从 sync.Pool 获取 TaobaoAlitripTravelProductBaseModifyAPIRequest
func GetTaobaoAlitripTravelProductBaseModifyAPIRequest() *TaobaoAlitripTravelProductBaseModifyAPIRequest {
	return poolTaobaoAlitripTravelProductBaseModifyAPIRequest.Get().(*TaobaoAlitripTravelProductBaseModifyAPIRequest)
}

// ReleaseTaobaoAlitripTravelProductBaseModifyAPIRequest 将 TaobaoAlitripTravelProductBaseModifyAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelProductBaseModifyAPIRequest(v *TaobaoAlitripTravelProductBaseModifyAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelProductBaseModifyAPIRequest.Put(v)
}
