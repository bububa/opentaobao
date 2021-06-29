package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】度假线路商品发布接口 API请求
taobao.alitrip.travel.item.base.add

旅行度假新商品发布接口。目前支持的类目包括：境内跟团游、出境跟团游、境内自由行、出境自由行、境内当地玩乐、境外玩乐套餐、境内邮轮、国际邮轮
*/
type TaobaoAlitripTravelItemBaseAddRequest struct {
    model.Params
    // 必填，商品基本信息
    baseInfo   *PontusTravelItemBaseInfo
    // 必填，预定规则结构。示例： [{  "rule_type": "fee_excluded",  "rule_desc": "费用包含描述"},{  "rule_type": "fee_included",  "rule_desc": "费用不含描述"},{  "rule_type": "order_info",  "rule_desc": "预定须知描述"}]
    bookingRules   []PontusTravelBookingRuleInfo
    // 特殊选填（当发布自由行商品时必填，其他情况不填）。自由行商品相关信息，自由行商品必填。大交通、酒店住宿、景点门票、租车、保险、餐饮、其他费用，自由行商品至少需要填写其中两项
    freedomItemExt   *PontusTravelFreedomItemExt
    // 特殊选填（当发布跟团游商品时必填，其他情况不填）。跟团游商品相关信息，发布跟团游商品时必填
    groupItemExt   *PontusTravelGroupItemExt
    // 选填，详细行程描述结构
    itineraries   []PontusTravelItemItineraryInfo
    // 选填，退款规则结构
    refundInfo   *PontusTravelItemRefundInfo
    // 选填，商品的销售属性相关信息
    salesInfo   *PontusTravelItemSaleInfo
    // 特殊选填（当发布邮轮商品时必填，其他情况不填）邮轮商品相关信息，发布邮轮商品时必填
    cruiseItemExt   *CruiseItemExt
    // 特殊选填（当发布同城活动商品时必填，其他情况不填）同城活动商品相关信息，发布同城活动商品时必填
    tcwlItemExt   *TcwlItemExt
}

// 初始化TaobaoAlitripTravelItemBaseAddRequest对象
func NewTaobaoAlitripTravelItemBaseAddRequest() *TaobaoAlitripTravelItemBaseAddRequest{
    return &TaobaoAlitripTravelItemBaseAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelItemBaseAddRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.item.base.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelItemBaseAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BaseInfo Setter
// 必填，商品基本信息
func (r *TaobaoAlitripTravelItemBaseAddRequest) SetBaseInfo(baseInfo *PontusTravelItemBaseInfo) error {
    r.baseInfo = baseInfo
    r.Set("base_info", baseInfo)
    return nil
}

// BaseInfo Getter
func (r TaobaoAlitripTravelItemBaseAddRequest) GetBaseInfo() *PontusTravelItemBaseInfo {
    return r.baseInfo
}
// BookingRules Setter
// 必填，预定规则结构。示例： [{  "rule_type": "fee_excluded",  "rule_desc": "费用包含描述"},{  "rule_type": "fee_included",  "rule_desc": "费用不含描述"},{  "rule_type": "order_info",  "rule_desc": "预定须知描述"}]
func (r *TaobaoAlitripTravelItemBaseAddRequest) SetBookingRules(bookingRules []PontusTravelBookingRuleInfo) error {
    r.bookingRules = bookingRules
    r.Set("booking_rules", bookingRules)
    return nil
}

// BookingRules Getter
func (r TaobaoAlitripTravelItemBaseAddRequest) GetBookingRules() []PontusTravelBookingRuleInfo {
    return r.bookingRules
}
// FreedomItemExt Setter
// 特殊选填（当发布自由行商品时必填，其他情况不填）。自由行商品相关信息，自由行商品必填。大交通、酒店住宿、景点门票、租车、保险、餐饮、其他费用，自由行商品至少需要填写其中两项
func (r *TaobaoAlitripTravelItemBaseAddRequest) SetFreedomItemExt(freedomItemExt *PontusTravelFreedomItemExt) error {
    r.freedomItemExt = freedomItemExt
    r.Set("freedom_item_ext", freedomItemExt)
    return nil
}

// FreedomItemExt Getter
func (r TaobaoAlitripTravelItemBaseAddRequest) GetFreedomItemExt() *PontusTravelFreedomItemExt {
    return r.freedomItemExt
}
// GroupItemExt Setter
// 特殊选填（当发布跟团游商品时必填，其他情况不填）。跟团游商品相关信息，发布跟团游商品时必填
func (r *TaobaoAlitripTravelItemBaseAddRequest) SetGroupItemExt(groupItemExt *PontusTravelGroupItemExt) error {
    r.groupItemExt = groupItemExt
    r.Set("group_item_ext", groupItemExt)
    return nil
}

// GroupItemExt Getter
func (r TaobaoAlitripTravelItemBaseAddRequest) GetGroupItemExt() *PontusTravelGroupItemExt {
    return r.groupItemExt
}
// Itineraries Setter
// 选填，详细行程描述结构
func (r *TaobaoAlitripTravelItemBaseAddRequest) SetItineraries(itineraries []PontusTravelItemItineraryInfo) error {
    r.itineraries = itineraries
    r.Set("itineraries", itineraries)
    return nil
}

// Itineraries Getter
func (r TaobaoAlitripTravelItemBaseAddRequest) GetItineraries() []PontusTravelItemItineraryInfo {
    return r.itineraries
}
// RefundInfo Setter
// 选填，退款规则结构
func (r *TaobaoAlitripTravelItemBaseAddRequest) SetRefundInfo(refundInfo *PontusTravelItemRefundInfo) error {
    r.refundInfo = refundInfo
    r.Set("refund_info", refundInfo)
    return nil
}

// RefundInfo Getter
func (r TaobaoAlitripTravelItemBaseAddRequest) GetRefundInfo() *PontusTravelItemRefundInfo {
    return r.refundInfo
}
// SalesInfo Setter
// 选填，商品的销售属性相关信息
func (r *TaobaoAlitripTravelItemBaseAddRequest) SetSalesInfo(salesInfo *PontusTravelItemSaleInfo) error {
    r.salesInfo = salesInfo
    r.Set("sales_info", salesInfo)
    return nil
}

// SalesInfo Getter
func (r TaobaoAlitripTravelItemBaseAddRequest) GetSalesInfo() *PontusTravelItemSaleInfo {
    return r.salesInfo
}
// CruiseItemExt Setter
// 特殊选填（当发布邮轮商品时必填，其他情况不填）邮轮商品相关信息，发布邮轮商品时必填
func (r *TaobaoAlitripTravelItemBaseAddRequest) SetCruiseItemExt(cruiseItemExt *CruiseItemExt) error {
    r.cruiseItemExt = cruiseItemExt
    r.Set("cruise_item_ext", cruiseItemExt)
    return nil
}

// CruiseItemExt Getter
func (r TaobaoAlitripTravelItemBaseAddRequest) GetCruiseItemExt() *CruiseItemExt {
    return r.cruiseItemExt
}
// TcwlItemExt Setter
// 特殊选填（当发布同城活动商品时必填，其他情况不填）同城活动商品相关信息，发布同城活动商品时必填
func (r *TaobaoAlitripTravelItemBaseAddRequest) SetTcwlItemExt(tcwlItemExt *TcwlItemExt) error {
    r.tcwlItemExt = tcwlItemExt
    r.Set("tcwl_item_ext", tcwlItemExt)
    return nil
}

// TcwlItemExt Getter
func (r TaobaoAlitripTravelItemBaseAddRequest) GetTcwlItemExt() *TcwlItemExt {
    return r.tcwlItemExt
}
