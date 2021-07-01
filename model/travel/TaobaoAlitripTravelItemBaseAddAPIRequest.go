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
type TaobaoAlitripTravelItemBaseAddAPIRequest struct {
    model.Params
    // 必填，商品基本信息
    _baseInfo   *PontusTravelItemBaseInfo
    // 必填，预定规则结构。示例： [{  "rule_type": "fee_excluded",  "rule_desc": "费用包含描述"},{  "rule_type": "fee_included",  "rule_desc": "费用不含描述"},{  "rule_type": "order_info",  "rule_desc": "预定须知描述"}]
    _bookingRules   []PontusTravelBookingRuleInfo
    // 特殊选填（当发布自由行商品时必填，其他情况不填）。自由行商品相关信息，自由行商品必填。大交通、酒店住宿、景点门票、租车、保险、餐饮、其他费用，自由行商品至少需要填写其中两项
    _freedomItemExt   *PontusTravelFreedomItemExt
    // 特殊选填（当发布跟团游商品时必填，其他情况不填）。跟团游商品相关信息，发布跟团游商品时必填
    _groupItemExt   *PontusTravelGroupItemExt
    // 选填，详细行程描述结构
    _itineraries   []PontusTravelItemItineraryInfo
    // 选填，退款规则结构
    _refundInfo   *PontusTravelItemRefundInfo
    // 选填，商品的销售属性相关信息
    _salesInfo   *PontusTravelItemSaleInfo
    // 特殊选填（当发布邮轮商品时必填，其他情况不填）邮轮商品相关信息，发布邮轮商品时必填
    _cruiseItemExt   *CruiseItemExt
    // 特殊选填（当发布同城活动商品时必填，其他情况不填）同城活动商品相关信息，发布同城活动商品时必填
    _tcwlItemExt   *TcwlItemExt
}

// 初始化TaobaoAlitripTravelItemBaseAddAPIRequest对象
func NewTaobaoAlitripTravelItemBaseAddRequest() *TaobaoAlitripTravelItemBaseAddAPIRequest{
    return &TaobaoAlitripTravelItemBaseAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelItemBaseAddAPIRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.item.base.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelItemBaseAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BaseInfo Setter
// 必填，商品基本信息
func (r *TaobaoAlitripTravelItemBaseAddAPIRequest) SetBaseInfo(_baseInfo *PontusTravelItemBaseInfo) error {
    r._baseInfo = _baseInfo
    r.Set("base_info", _baseInfo)
    return nil
}

// BaseInfo Getter
func (r TaobaoAlitripTravelItemBaseAddAPIRequest) GetBaseInfo() *PontusTravelItemBaseInfo {
    return r._baseInfo
}
// BookingRules Setter
// 必填，预定规则结构。示例： [{  "rule_type": "fee_excluded",  "rule_desc": "费用包含描述"},{  "rule_type": "fee_included",  "rule_desc": "费用不含描述"},{  "rule_type": "order_info",  "rule_desc": "预定须知描述"}]
func (r *TaobaoAlitripTravelItemBaseAddAPIRequest) SetBookingRules(_bookingRules []PontusTravelBookingRuleInfo) error {
    r._bookingRules = _bookingRules
    r.Set("booking_rules", _bookingRules)
    return nil
}

// BookingRules Getter
func (r TaobaoAlitripTravelItemBaseAddAPIRequest) GetBookingRules() []PontusTravelBookingRuleInfo {
    return r._bookingRules
}
// FreedomItemExt Setter
// 特殊选填（当发布自由行商品时必填，其他情况不填）。自由行商品相关信息，自由行商品必填。大交通、酒店住宿、景点门票、租车、保险、餐饮、其他费用，自由行商品至少需要填写其中两项
func (r *TaobaoAlitripTravelItemBaseAddAPIRequest) SetFreedomItemExt(_freedomItemExt *PontusTravelFreedomItemExt) error {
    r._freedomItemExt = _freedomItemExt
    r.Set("freedom_item_ext", _freedomItemExt)
    return nil
}

// FreedomItemExt Getter
func (r TaobaoAlitripTravelItemBaseAddAPIRequest) GetFreedomItemExt() *PontusTravelFreedomItemExt {
    return r._freedomItemExt
}
// GroupItemExt Setter
// 特殊选填（当发布跟团游商品时必填，其他情况不填）。跟团游商品相关信息，发布跟团游商品时必填
func (r *TaobaoAlitripTravelItemBaseAddAPIRequest) SetGroupItemExt(_groupItemExt *PontusTravelGroupItemExt) error {
    r._groupItemExt = _groupItemExt
    r.Set("group_item_ext", _groupItemExt)
    return nil
}

// GroupItemExt Getter
func (r TaobaoAlitripTravelItemBaseAddAPIRequest) GetGroupItemExt() *PontusTravelGroupItemExt {
    return r._groupItemExt
}
// Itineraries Setter
// 选填，详细行程描述结构
func (r *TaobaoAlitripTravelItemBaseAddAPIRequest) SetItineraries(_itineraries []PontusTravelItemItineraryInfo) error {
    r._itineraries = _itineraries
    r.Set("itineraries", _itineraries)
    return nil
}

// Itineraries Getter
func (r TaobaoAlitripTravelItemBaseAddAPIRequest) GetItineraries() []PontusTravelItemItineraryInfo {
    return r._itineraries
}
// RefundInfo Setter
// 选填，退款规则结构
func (r *TaobaoAlitripTravelItemBaseAddAPIRequest) SetRefundInfo(_refundInfo *PontusTravelItemRefundInfo) error {
    r._refundInfo = _refundInfo
    r.Set("refund_info", _refundInfo)
    return nil
}

// RefundInfo Getter
func (r TaobaoAlitripTravelItemBaseAddAPIRequest) GetRefundInfo() *PontusTravelItemRefundInfo {
    return r._refundInfo
}
// SalesInfo Setter
// 选填，商品的销售属性相关信息
func (r *TaobaoAlitripTravelItemBaseAddAPIRequest) SetSalesInfo(_salesInfo *PontusTravelItemSaleInfo) error {
    r._salesInfo = _salesInfo
    r.Set("sales_info", _salesInfo)
    return nil
}

// SalesInfo Getter
func (r TaobaoAlitripTravelItemBaseAddAPIRequest) GetSalesInfo() *PontusTravelItemSaleInfo {
    return r._salesInfo
}
// CruiseItemExt Setter
// 特殊选填（当发布邮轮商品时必填，其他情况不填）邮轮商品相关信息，发布邮轮商品时必填
func (r *TaobaoAlitripTravelItemBaseAddAPIRequest) SetCruiseItemExt(_cruiseItemExt *CruiseItemExt) error {
    r._cruiseItemExt = _cruiseItemExt
    r.Set("cruise_item_ext", _cruiseItemExt)
    return nil
}

// CruiseItemExt Getter
func (r TaobaoAlitripTravelItemBaseAddAPIRequest) GetCruiseItemExt() *CruiseItemExt {
    return r._cruiseItemExt
}
// TcwlItemExt Setter
// 特殊选填（当发布同城活动商品时必填，其他情况不填）同城活动商品相关信息，发布同城活动商品时必填
func (r *TaobaoAlitripTravelItemBaseAddAPIRequest) SetTcwlItemExt(_tcwlItemExt *TcwlItemExt) error {
    r._tcwlItemExt = _tcwlItemExt
    r.Set("tcwl_item_ext", _tcwlItemExt)
    return nil
}

// TcwlItemExt Getter
func (r TaobaoAlitripTravelItemBaseAddAPIRequest) GetTcwlItemExt() *TcwlItemExt {
    return r._tcwlItemExt
}
