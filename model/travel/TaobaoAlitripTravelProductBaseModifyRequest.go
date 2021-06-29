package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商编辑产品API API请求
taobao.alitrip.travel.product.base.modify

飞猪供销平台供应商可通过该API编辑产品
*/
type TaobaoAlitripTravelProductBaseModifyRequest struct {
    model.Params
    // 详细行程描述结构
    _itineraries   []ItemItineraryInfo
    // 产品基本信息
    _baseInfo   *ProductBaseInfo
    // 产品id
    _itemId   int64
    // 退款规则结构
    _refundInfo   *ItemRefundInfo
    // 预定规则结构。示例： [{ "rule_type": "fee_excluded", "rule_desc": "费用包含描述"},{ "rule_type": "fee_included", "rule_desc": "费用不含描述"},{ "rule_type": "order_info", "rule_desc": "预定须知描述"}]
    _bookingRules   []BookingRuleInfo
    // 邮轮商品相关信息
    _cruiseProductExt   *CruiseProductExt
    // 商品的销售属性相关信息
    _productSaleInfo   *ProductSaleInfo
}

// 初始化TaobaoAlitripTravelProductBaseModifyRequest对象
func NewTaobaoAlitripTravelProductBaseModifyRequest() *TaobaoAlitripTravelProductBaseModifyRequest{
    return &TaobaoAlitripTravelProductBaseModifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelProductBaseModifyRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.product.base.modify"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelProductBaseModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Itineraries Setter
// 详细行程描述结构
func (r *TaobaoAlitripTravelProductBaseModifyRequest) SetItineraries(_itineraries []ItemItineraryInfo) error {
    r._itineraries = _itineraries
    r.Set("itineraries", _itineraries)
    return nil
}

// Itineraries Getter
func (r TaobaoAlitripTravelProductBaseModifyRequest) GetItineraries() []ItemItineraryInfo {
    return r._itineraries
}
// BaseInfo Setter
// 产品基本信息
func (r *TaobaoAlitripTravelProductBaseModifyRequest) SetBaseInfo(_baseInfo *ProductBaseInfo) error {
    r._baseInfo = _baseInfo
    r.Set("base_info", _baseInfo)
    return nil
}

// BaseInfo Getter
func (r TaobaoAlitripTravelProductBaseModifyRequest) GetBaseInfo() *ProductBaseInfo {
    return r._baseInfo
}
// ItemId Setter
// 产品id
func (r *TaobaoAlitripTravelProductBaseModifyRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoAlitripTravelProductBaseModifyRequest) GetItemId() int64 {
    return r._itemId
}
// RefundInfo Setter
// 退款规则结构
func (r *TaobaoAlitripTravelProductBaseModifyRequest) SetRefundInfo(_refundInfo *ItemRefundInfo) error {
    r._refundInfo = _refundInfo
    r.Set("refund_info", _refundInfo)
    return nil
}

// RefundInfo Getter
func (r TaobaoAlitripTravelProductBaseModifyRequest) GetRefundInfo() *ItemRefundInfo {
    return r._refundInfo
}
// BookingRules Setter
// 预定规则结构。示例： [{ "rule_type": "fee_excluded", "rule_desc": "费用包含描述"},{ "rule_type": "fee_included", "rule_desc": "费用不含描述"},{ "rule_type": "order_info", "rule_desc": "预定须知描述"}]
func (r *TaobaoAlitripTravelProductBaseModifyRequest) SetBookingRules(_bookingRules []BookingRuleInfo) error {
    r._bookingRules = _bookingRules
    r.Set("booking_rules", _bookingRules)
    return nil
}

// BookingRules Getter
func (r TaobaoAlitripTravelProductBaseModifyRequest) GetBookingRules() []BookingRuleInfo {
    return r._bookingRules
}
// CruiseProductExt Setter
// 邮轮商品相关信息
func (r *TaobaoAlitripTravelProductBaseModifyRequest) SetCruiseProductExt(_cruiseProductExt *CruiseProductExt) error {
    r._cruiseProductExt = _cruiseProductExt
    r.Set("cruise_product_ext", _cruiseProductExt)
    return nil
}

// CruiseProductExt Getter
func (r TaobaoAlitripTravelProductBaseModifyRequest) GetCruiseProductExt() *CruiseProductExt {
    return r._cruiseProductExt
}
// ProductSaleInfo Setter
// 商品的销售属性相关信息
func (r *TaobaoAlitripTravelProductBaseModifyRequest) SetProductSaleInfo(_productSaleInfo *ProductSaleInfo) error {
    r._productSaleInfo = _productSaleInfo
    r.Set("product_sale_info", _productSaleInfo)
    return nil
}

// ProductSaleInfo Getter
func (r TaobaoAlitripTravelProductBaseModifyRequest) GetProductSaleInfo() *ProductSaleInfo {
    return r._productSaleInfo
}
