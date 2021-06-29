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
    groupItemExt   *PontusTravelGroupItemExt
    // 商品基本信息
    baseInfo   *PontusTravelItemBaseInfo
    // 详细行程描述结构
    itineraries   []PontusTravelItemItineraryInfo
    // 商品id
    itemId   int64
    // 退款规则结构
    refundInfo   *PontusTravelItemRefundInfo
    // 预定规则结构
    bookingRules   []PontusTravelBookingRuleInfo
    // 自由行商品相关信息
    freedomItemExt   *PontusTravelFreedomItemExt
    // 商品的销售属性相关信息
    salesInfo   *PontusTravelItemSaleInfo
    // 可选，支持清空商品某些字段内容。多个待清空字段名之间以英文逗号分隔。目前支持清除以下字段内容：seller_cids表示删除关联的店铺类目id
    fieldsToClean   string
    // 邮轮商品相关扩展信息
    cruiseItemExt   *CruiseItemExt
    // 同城活动商品相关扩展信息
    tcwlItemExt   *TcwlItemExt
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
func (r *TaobaoAlitripTravelItemBaseModifyRequest) SetGroupItemExt(groupItemExt *PontusTravelGroupItemExt) error {
    r.groupItemExt = groupItemExt
    r.Set("group_item_ext", groupItemExt)
    return nil
}

// GroupItemExt Getter
func (r TaobaoAlitripTravelItemBaseModifyRequest) GetGroupItemExt() *PontusTravelGroupItemExt {
    return r.groupItemExt
}
// BaseInfo Setter
// 商品基本信息
func (r *TaobaoAlitripTravelItemBaseModifyRequest) SetBaseInfo(baseInfo *PontusTravelItemBaseInfo) error {
    r.baseInfo = baseInfo
    r.Set("base_info", baseInfo)
    return nil
}

// BaseInfo Getter
func (r TaobaoAlitripTravelItemBaseModifyRequest) GetBaseInfo() *PontusTravelItemBaseInfo {
    return r.baseInfo
}
// Itineraries Setter
// 详细行程描述结构
func (r *TaobaoAlitripTravelItemBaseModifyRequest) SetItineraries(itineraries []PontusTravelItemItineraryInfo) error {
    r.itineraries = itineraries
    r.Set("itineraries", itineraries)
    return nil
}

// Itineraries Getter
func (r TaobaoAlitripTravelItemBaseModifyRequest) GetItineraries() []PontusTravelItemItineraryInfo {
    return r.itineraries
}
// ItemId Setter
// 商品id
func (r *TaobaoAlitripTravelItemBaseModifyRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoAlitripTravelItemBaseModifyRequest) GetItemId() int64 {
    return r.itemId
}
// RefundInfo Setter
// 退款规则结构
func (r *TaobaoAlitripTravelItemBaseModifyRequest) SetRefundInfo(refundInfo *PontusTravelItemRefundInfo) error {
    r.refundInfo = refundInfo
    r.Set("refund_info", refundInfo)
    return nil
}

// RefundInfo Getter
func (r TaobaoAlitripTravelItemBaseModifyRequest) GetRefundInfo() *PontusTravelItemRefundInfo {
    return r.refundInfo
}
// BookingRules Setter
// 预定规则结构
func (r *TaobaoAlitripTravelItemBaseModifyRequest) SetBookingRules(bookingRules []PontusTravelBookingRuleInfo) error {
    r.bookingRules = bookingRules
    r.Set("booking_rules", bookingRules)
    return nil
}

// BookingRules Getter
func (r TaobaoAlitripTravelItemBaseModifyRequest) GetBookingRules() []PontusTravelBookingRuleInfo {
    return r.bookingRules
}
// FreedomItemExt Setter
// 自由行商品相关信息
func (r *TaobaoAlitripTravelItemBaseModifyRequest) SetFreedomItemExt(freedomItemExt *PontusTravelFreedomItemExt) error {
    r.freedomItemExt = freedomItemExt
    r.Set("freedom_item_ext", freedomItemExt)
    return nil
}

// FreedomItemExt Getter
func (r TaobaoAlitripTravelItemBaseModifyRequest) GetFreedomItemExt() *PontusTravelFreedomItemExt {
    return r.freedomItemExt
}
// SalesInfo Setter
// 商品的销售属性相关信息
func (r *TaobaoAlitripTravelItemBaseModifyRequest) SetSalesInfo(salesInfo *PontusTravelItemSaleInfo) error {
    r.salesInfo = salesInfo
    r.Set("sales_info", salesInfo)
    return nil
}

// SalesInfo Getter
func (r TaobaoAlitripTravelItemBaseModifyRequest) GetSalesInfo() *PontusTravelItemSaleInfo {
    return r.salesInfo
}
// FieldsToClean Setter
// 可选，支持清空商品某些字段内容。多个待清空字段名之间以英文逗号分隔。目前支持清除以下字段内容：seller_cids表示删除关联的店铺类目id
func (r *TaobaoAlitripTravelItemBaseModifyRequest) SetFieldsToClean(fieldsToClean string) error {
    r.fieldsToClean = fieldsToClean
    r.Set("fields_to_clean", fieldsToClean)
    return nil
}

// FieldsToClean Getter
func (r TaobaoAlitripTravelItemBaseModifyRequest) GetFieldsToClean() string {
    return r.fieldsToClean
}
// CruiseItemExt Setter
// 邮轮商品相关扩展信息
func (r *TaobaoAlitripTravelItemBaseModifyRequest) SetCruiseItemExt(cruiseItemExt *CruiseItemExt) error {
    r.cruiseItemExt = cruiseItemExt
    r.Set("cruise_item_ext", cruiseItemExt)
    return nil
}

// CruiseItemExt Getter
func (r TaobaoAlitripTravelItemBaseModifyRequest) GetCruiseItemExt() *CruiseItemExt {
    return r.cruiseItemExt
}
// TcwlItemExt Setter
// 同城活动商品相关扩展信息
func (r *TaobaoAlitripTravelItemBaseModifyRequest) SetTcwlItemExt(tcwlItemExt *TcwlItemExt) error {
    r.tcwlItemExt = tcwlItemExt
    r.Set("tcwl_item_ext", tcwlItemExt)
    return nil
}

// TcwlItemExt Getter
func (r TaobaoAlitripTravelItemBaseModifyRequest) GetTcwlItemExt() *TcwlItemExt {
    return r.tcwlItemExt
}
