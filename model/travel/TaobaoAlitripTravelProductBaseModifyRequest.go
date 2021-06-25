package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商编辑产品API APIRequest
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

func NewTaobaoAlitripTravelProductBaseModifyRequest() *TaobaoAlitripTravelProductBaseModifyRequest{
    return &TaobaoAlitripTravelProductBaseModifyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripTravelProductBaseModifyRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.product.base.modify"
}

func (r TaobaoAlitripTravelProductBaseModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripTravelProductBaseModifyRequest) SetItineraries(itineraries []ItemItineraryInfo) error {
    r.itineraries = itineraries
    r.Set("itineraries", itineraries)
    return nil
}

func (r TaobaoAlitripTravelProductBaseModifyRequest) GetItineraries() []ItemItineraryInfo {
    return r.itineraries
}

func (r *TaobaoAlitripTravelProductBaseModifyRequest) SetBaseInfo(baseInfo *ProductBaseInfo) error {
    r.baseInfo = baseInfo
    r.Set("base_info", baseInfo)
    return nil
}

func (r TaobaoAlitripTravelProductBaseModifyRequest) GetBaseInfo() *ProductBaseInfo {
    return r.baseInfo
}

func (r *TaobaoAlitripTravelProductBaseModifyRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoAlitripTravelProductBaseModifyRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoAlitripTravelProductBaseModifyRequest) SetRefundInfo(refundInfo *ItemRefundInfo) error {
    r.refundInfo = refundInfo
    r.Set("refund_info", refundInfo)
    return nil
}

func (r TaobaoAlitripTravelProductBaseModifyRequest) GetRefundInfo() *ItemRefundInfo {
    return r.refundInfo
}

func (r *TaobaoAlitripTravelProductBaseModifyRequest) SetBookingRules(bookingRules []BookingRuleInfo) error {
    r.bookingRules = bookingRules
    r.Set("booking_rules", bookingRules)
    return nil
}

func (r TaobaoAlitripTravelProductBaseModifyRequest) GetBookingRules() []BookingRuleInfo {
    return r.bookingRules
}

func (r *TaobaoAlitripTravelProductBaseModifyRequest) SetCruiseProductExt(cruiseProductExt *CruiseProductExt) error {
    r.cruiseProductExt = cruiseProductExt
    r.Set("cruise_product_ext", cruiseProductExt)
    return nil
}

func (r TaobaoAlitripTravelProductBaseModifyRequest) GetCruiseProductExt() *CruiseProductExt {
    return r.cruiseProductExt
}

func (r *TaobaoAlitripTravelProductBaseModifyRequest) SetProductSaleInfo(productSaleInfo *ProductSaleInfo) error {
    r.productSaleInfo = productSaleInfo
    r.Set("product_sale_info", productSaleInfo)
    return nil
}

func (r TaobaoAlitripTravelProductBaseModifyRequest) GetProductSaleInfo() *ProductSaleInfo {
    return r.productSaleInfo
}

