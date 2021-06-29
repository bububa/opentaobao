package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】度假线路商品编辑接口 API请求
taobao.alitrip.travel.item.base.modify

旅行度假新商品基本信息修改接口 第三版。提供商家通过TOP API方式修改商品除sku外的基本信息。
*/
type TaobaoAlitripTravelItemBaseModifyRequest struct {
    model.Params
    // 跟团游商品相关信息
    _groupItemExt   *PontusTravelGroupItemExt
    // 商品基本信息
    _baseInfo   *PontusTravelItemBaseInfo
    // 详细行程描述结构
    _itineraries   []PontusTravelItemItineraryInfo
    // 商品id
    _itemId   int64
    // 退款规则结构
    _refundInfo   *PontusTravelItemRefundInfo
    // 预定规则结构
    _bookingRules   []PontusTravelBookingRuleInfo
    // 自由行商品相关信息
    _freedomItemExt   *PontusTravelFreedomItemExt
    // 商品的销售属性相关信息
    _salesInfo   *PontusTravelItemSaleInfo
    // 可选，支持清空商品某些字段内容。多个待清空字段名之间以英文逗号分隔。目前支持清除以下字段内容：seller_cids表示删除关联的店铺类目id
    _fieldsToClean   string
    // 邮轮商品相关扩展信息
    _cruiseItemExt   *CruiseItemExt
    // 同城活动商品相关扩展信息
    _tcwlItemExt   *TcwlItemExt
}

// 初始化TaobaoAlitripTravelItemBaseModifyRequest对象
func NewTaobaoAlitripTravelItemBaseModifyRequest() *TaobaoAlitripTravelItemBaseModifyRequest{
    return &TaobaoAlitripTravelItemBaseModifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelItemBaseModifyRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.item.base.modify"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelItemBaseModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GroupItemExt Setter
// 跟团游商品相关信息
func (r *TaobaoAlitripTravelItemBaseModifyRequest) SetGroupItemExt(_groupItemExt *PontusTravelGroupItemExt) error {
    r._groupItemExt = _groupItemExt
    r.Set("group_item_ext", _groupItemExt)
    return nil
}

// GroupItemExt Getter
func (r TaobaoAlitripTravelItemBaseModifyRequest) GetGroupItemExt() *PontusTravelGroupItemExt {
    return r._groupItemExt
}
// BaseInfo Setter
// 商品基本信息
func (r *TaobaoAlitripTravelItemBaseModifyRequest) SetBaseInfo(_baseInfo *PontusTravelItemBaseInfo) error {
    r._baseInfo = _baseInfo
    r.Set("base_info", _baseInfo)
    return nil
}

// BaseInfo Getter
func (r TaobaoAlitripTravelItemBaseModifyRequest) GetBaseInfo() *PontusTravelItemBaseInfo {
    return r._baseInfo
}
// Itineraries Setter
// 详细行程描述结构
func (r *TaobaoAlitripTravelItemBaseModifyRequest) SetItineraries(_itineraries []PontusTravelItemItineraryInfo) error {
    r._itineraries = _itineraries
    r.Set("itineraries", _itineraries)
    return nil
}

// Itineraries Getter
func (r TaobaoAlitripTravelItemBaseModifyRequest) GetItineraries() []PontusTravelItemItineraryInfo {
    return r._itineraries
}
// ItemId Setter
// 商品id
func (r *TaobaoAlitripTravelItemBaseModifyRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoAlitripTravelItemBaseModifyRequest) GetItemId() int64 {
    return r._itemId
}
// RefundInfo Setter
// 退款规则结构
func (r *TaobaoAlitripTravelItemBaseModifyRequest) SetRefundInfo(_refundInfo *PontusTravelItemRefundInfo) error {
    r._refundInfo = _refundInfo
    r.Set("refund_info", _refundInfo)
    return nil
}

// RefundInfo Getter
func (r TaobaoAlitripTravelItemBaseModifyRequest) GetRefundInfo() *PontusTravelItemRefundInfo {
    return r._refundInfo
}
// BookingRules Setter
// 预定规则结构
func (r *TaobaoAlitripTravelItemBaseModifyRequest) SetBookingRules(_bookingRules []PontusTravelBookingRuleInfo) error {
    r._bookingRules = _bookingRules
    r.Set("booking_rules", _bookingRules)
    return nil
}

// BookingRules Getter
func (r TaobaoAlitripTravelItemBaseModifyRequest) GetBookingRules() []PontusTravelBookingRuleInfo {
    return r._bookingRules
}
// FreedomItemExt Setter
// 自由行商品相关信息
func (r *TaobaoAlitripTravelItemBaseModifyRequest) SetFreedomItemExt(_freedomItemExt *PontusTravelFreedomItemExt) error {
    r._freedomItemExt = _freedomItemExt
    r.Set("freedom_item_ext", _freedomItemExt)
    return nil
}

// FreedomItemExt Getter
func (r TaobaoAlitripTravelItemBaseModifyRequest) GetFreedomItemExt() *PontusTravelFreedomItemExt {
    return r._freedomItemExt
}
// SalesInfo Setter
// 商品的销售属性相关信息
func (r *TaobaoAlitripTravelItemBaseModifyRequest) SetSalesInfo(_salesInfo *PontusTravelItemSaleInfo) error {
    r._salesInfo = _salesInfo
    r.Set("sales_info", _salesInfo)
    return nil
}

// SalesInfo Getter
func (r TaobaoAlitripTravelItemBaseModifyRequest) GetSalesInfo() *PontusTravelItemSaleInfo {
    return r._salesInfo
}
// FieldsToClean Setter
// 可选，支持清空商品某些字段内容。多个待清空字段名之间以英文逗号分隔。目前支持清除以下字段内容：seller_cids表示删除关联的店铺类目id
func (r *TaobaoAlitripTravelItemBaseModifyRequest) SetFieldsToClean(_fieldsToClean string) error {
    r._fieldsToClean = _fieldsToClean
    r.Set("fields_to_clean", _fieldsToClean)
    return nil
}

// FieldsToClean Getter
func (r TaobaoAlitripTravelItemBaseModifyRequest) GetFieldsToClean() string {
    return r._fieldsToClean
}
// CruiseItemExt Setter
// 邮轮商品相关扩展信息
func (r *TaobaoAlitripTravelItemBaseModifyRequest) SetCruiseItemExt(_cruiseItemExt *CruiseItemExt) error {
    r._cruiseItemExt = _cruiseItemExt
    r.Set("cruise_item_ext", _cruiseItemExt)
    return nil
}

// CruiseItemExt Getter
func (r TaobaoAlitripTravelItemBaseModifyRequest) GetCruiseItemExt() *CruiseItemExt {
    return r._cruiseItemExt
}
// TcwlItemExt Setter
// 同城活动商品相关扩展信息
func (r *TaobaoAlitripTravelItemBaseModifyRequest) SetTcwlItemExt(_tcwlItemExt *TcwlItemExt) error {
    r._tcwlItemExt = _tcwlItemExt
    r.Set("tcwl_item_ext", _tcwlItemExt)
    return nil
}

// TcwlItemExt Getter
func (r TaobaoAlitripTravelItemBaseModifyRequest) GetTcwlItemExt() *TcwlItemExt {
    return r._tcwlItemExt
}
