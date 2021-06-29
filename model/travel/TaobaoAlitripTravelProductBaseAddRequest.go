package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商新增产品API API请求
taobao.alitrip.travel.product.base.add

飞猪供销平台供应商可通过该API发布新产品
*/
type TaobaoAlitripTravelProductBaseAddRequest struct {
    model.Params
    // 产品基本信息
    _baseInfo   *ProductBaseInfo
    // 选填，详细行程描述结构
    _itineraries   []ItemItineraryInfo
    // 选填，退款规则结构
    _refundInfo   *ItemRefundInfo
    // 必填，预定规则结构。示例： [{  "rule_type": "fee_excluded",  "rule_desc": "费用包含描述"},{  "rule_type": "fee_included",  "rule_desc": "费用不含描述"},{  "rule_type": "order_info",  "rule_desc": "预定须知描述"}]
    _bookingRules   []BookingRuleInfo
    // 特殊选填（当发布邮轮商品时必填，其他情况不填）邮轮商品相关信息，发布邮轮商品时必填
    _cruiseProductExt   *CruiseProductExt
    // 选填，商品的销售属性相关信息
    _productSaleInfo   *ProductSaleInfo
}

// 初始化TaobaoAlitripTravelProductBaseAddRequest对象
func NewTaobaoAlitripTravelProductBaseAddRequest() *TaobaoAlitripTravelProductBaseAddRequest{
    return &TaobaoAlitripTravelProductBaseAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelProductBaseAddRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.product.base.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelProductBaseAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BaseInfo Setter
// 产品基本信息
func (r *TaobaoAlitripTravelProductBaseAddRequest) SetBaseInfo(_baseInfo *ProductBaseInfo) error {
    r._baseInfo = _baseInfo
    r.Set("base_info", _baseInfo)
    return nil
}

// BaseInfo Getter
func (r TaobaoAlitripTravelProductBaseAddRequest) GetBaseInfo() *ProductBaseInfo {
    return r._baseInfo
}
// Itineraries Setter
// 选填，详细行程描述结构
func (r *TaobaoAlitripTravelProductBaseAddRequest) SetItineraries(_itineraries []ItemItineraryInfo) error {
    r._itineraries = _itineraries
    r.Set("itineraries", _itineraries)
    return nil
}

// Itineraries Getter
func (r TaobaoAlitripTravelProductBaseAddRequest) GetItineraries() []ItemItineraryInfo {
    return r._itineraries
}
// RefundInfo Setter
// 选填，退款规则结构
func (r *TaobaoAlitripTravelProductBaseAddRequest) SetRefundInfo(_refundInfo *ItemRefundInfo) error {
    r._refundInfo = _refundInfo
    r.Set("refund_info", _refundInfo)
    return nil
}

// RefundInfo Getter
func (r TaobaoAlitripTravelProductBaseAddRequest) GetRefundInfo() *ItemRefundInfo {
    return r._refundInfo
}
// BookingRules Setter
// 必填，预定规则结构。示例： [{  "rule_type": "fee_excluded",  "rule_desc": "费用包含描述"},{  "rule_type": "fee_included",  "rule_desc": "费用不含描述"},{  "rule_type": "order_info",  "rule_desc": "预定须知描述"}]
func (r *TaobaoAlitripTravelProductBaseAddRequest) SetBookingRules(_bookingRules []BookingRuleInfo) error {
    r._bookingRules = _bookingRules
    r.Set("booking_rules", _bookingRules)
    return nil
}

// BookingRules Getter
func (r TaobaoAlitripTravelProductBaseAddRequest) GetBookingRules() []BookingRuleInfo {
    return r._bookingRules
}
// CruiseProductExt Setter
// 特殊选填（当发布邮轮商品时必填，其他情况不填）邮轮商品相关信息，发布邮轮商品时必填
func (r *TaobaoAlitripTravelProductBaseAddRequest) SetCruiseProductExt(_cruiseProductExt *CruiseProductExt) error {
    r._cruiseProductExt = _cruiseProductExt
    r.Set("cruise_product_ext", _cruiseProductExt)
    return nil
}

// CruiseProductExt Getter
func (r TaobaoAlitripTravelProductBaseAddRequest) GetCruiseProductExt() *CruiseProductExt {
    return r._cruiseProductExt
}
// ProductSaleInfo Setter
// 选填，商品的销售属性相关信息
func (r *TaobaoAlitripTravelProductBaseAddRequest) SetProductSaleInfo(_productSaleInfo *ProductSaleInfo) error {
    r._productSaleInfo = _productSaleInfo
    r.Set("product_sale_info", _productSaleInfo)
    return nil
}

// ProductSaleInfo Getter
func (r TaobaoAlitripTravelProductBaseAddRequest) GetProductSaleInfo() *ProductSaleInfo {
    return r._productSaleInfo
}
