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
    itineraries   []ItemItineraryInfo
    // 产品基本信息
    baseInfo   *ProductBaseInfo
    // 产品id
    itemId   int64
    // 退款规则结构
    refundInfo   *ItemRefundInfo
    // 预定规则结构。示例： [{ "rule_type": "fee_excluded", "rule_desc": "费用包含描述"},{ "rule_type": "fee_included", "rule_desc": "费用不含描述"},{ "rule_type": "order_info", "rule_desc": "预定须知描述"}]
    bookingRules   []BookingRuleInfo
    // 邮轮商品相关信息
    cruiseProductExt   *CruiseProductExt
    // 商品的销售属性相关信息
    productSaleInfo   *ProductSaleInfo
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
func (r *TaobaoAlitripTravelProductBaseModifyRequest) SetItineraries(itineraries []ItemItineraryInfo) error {
    r.itineraries = itineraries
    r.Set("itineraries", itineraries)
    return nil
}

// Itineraries Getter
func (r TaobaoAlitripTravelProductBaseModifyRequest) GetItineraries() []ItemItineraryInfo {
    return r.itineraries
}
// BaseInfo Setter
// 产品基本信息
func (r *TaobaoAlitripTravelProductBaseModifyRequest) SetBaseInfo(baseInfo *ProductBaseInfo) error {
    r.baseInfo = baseInfo
    r.Set("base_info", baseInfo)
    return nil
}

// BaseInfo Getter
func (r TaobaoAlitripTravelProductBaseModifyRequest) GetBaseInfo() *ProductBaseInfo {
    return r.baseInfo
}
// ItemId Setter
// 产品id
func (r *TaobaoAlitripTravelProductBaseModifyRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoAlitripTravelProductBaseModifyRequest) GetItemId() int64 {
    return r.itemId
}
// RefundInfo Setter
// 退款规则结构
func (r *TaobaoAlitripTravelProductBaseModifyRequest) SetRefundInfo(refundInfo *ItemRefundInfo) error {
    r.refundInfo = refundInfo
    r.Set("refund_info", refundInfo)
    return nil
}

// RefundInfo Getter
func (r TaobaoAlitripTravelProductBaseModifyRequest) GetRefundInfo() *ItemRefundInfo {
    return r.refundInfo
}
// BookingRules Setter
// 预定规则结构。示例： [{ "rule_type": "fee_excluded", "rule_desc": "费用包含描述"},{ "rule_type": "fee_included", "rule_desc": "费用不含描述"},{ "rule_type": "order_info", "rule_desc": "预定须知描述"}]
func (r *TaobaoAlitripTravelProductBaseModifyRequest) SetBookingRules(bookingRules []BookingRuleInfo) error {
    r.bookingRules = bookingRules
    r.Set("booking_rules", bookingRules)
    return nil
}

// BookingRules Getter
func (r TaobaoAlitripTravelProductBaseModifyRequest) GetBookingRules() []BookingRuleInfo {
    return r.bookingRules
}
// CruiseProductExt Setter
// 邮轮商品相关信息
func (r *TaobaoAlitripTravelProductBaseModifyRequest) SetCruiseProductExt(cruiseProductExt *CruiseProductExt) error {
    r.cruiseProductExt = cruiseProductExt
    r.Set("cruise_product_ext", cruiseProductExt)
    return nil
}

// CruiseProductExt Getter
func (r TaobaoAlitripTravelProductBaseModifyRequest) GetCruiseProductExt() *CruiseProductExt {
    return r.cruiseProductExt
}
// ProductSaleInfo Setter
// 商品的销售属性相关信息
func (r *TaobaoAlitripTravelProductBaseModifyRequest) SetProductSaleInfo(productSaleInfo *ProductSaleInfo) error {
    r.productSaleInfo = productSaleInfo
    r.Set("product_sale_info", productSaleInfo)
    return nil
}

// ProductSaleInfo Getter
func (r TaobaoAlitripTravelProductBaseModifyRequest) GetProductSaleInfo() *ProductSaleInfo {
    return r.productSaleInfo
}
