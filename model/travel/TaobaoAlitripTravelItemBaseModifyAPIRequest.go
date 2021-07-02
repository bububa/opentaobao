package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelItemBaseModifyAPIRequest 【API3.0】度假线路商品编辑接口 API请求
// taobao.alitrip.travel.item.base.modify
//
// 旅行度假新商品基本信息修改接口 第三版。提供商家通过TOP API方式修改商品除sku外的基本信息。
type TaobaoAlitripTravelItemBaseModifyAPIRequest struct {
	model.Params
	// 跟团游商品相关信息
	_groupItemExt *PontusTravelGroupItemExt
	// 商品基本信息
	_baseInfo *PontusTravelItemBaseInfo
	// 详细行程描述结构
	_itineraries []PontusTravelItemItineraryInfo
	// 商品id
	_itemId int64
	// 退款规则结构
	_refundInfo *PontusTravelItemRefundInfo
	// 预定规则结构
	_bookingRules []PontusTravelBookingRuleInfo
	// 自由行商品相关信息
	_freedomItemExt *PontusTravelFreedomItemExt
	// 商品的销售属性相关信息
	_salesInfo *PontusTravelItemSaleInfo
	// 可选，支持清空商品某些字段内容。多个待清空字段名之间以英文逗号分隔。目前支持清除以下字段内容：seller_cids表示删除关联的店铺类目id
	_fieldsToClean string
	// 邮轮商品相关扩展信息
	_cruiseItemExt *CruiseItemExt
	// 同城活动商品相关扩展信息
	_tcwlItemExt *TcwlItemExt
}

// NewTaobaoAlitripTravelItemBaseModifyRequest 初始化TaobaoAlitripTravelItemBaseModifyAPIRequest对象
func NewTaobaoAlitripTravelItemBaseModifyRequest() *TaobaoAlitripTravelItemBaseModifyAPIRequest {
	return &TaobaoAlitripTravelItemBaseModifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelItemBaseModifyAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.item.base.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelItemBaseModifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is GroupItemExt Setter
// 跟团游商品相关信息
func (r *TaobaoAlitripTravelItemBaseModifyAPIRequest) SetGroupItemExt(_groupItemExt *PontusTravelGroupItemExt) error {
	r._groupItemExt = _groupItemExt
	r.Set("group_item_ext", _groupItemExt)
	return nil
}

// Get GroupItemExt Getter
func (r TaobaoAlitripTravelItemBaseModifyAPIRequest) GetGroupItemExt() *PontusTravelGroupItemExt {
	return r._groupItemExt
}

// Set is BaseInfo Setter
// 商品基本信息
func (r *TaobaoAlitripTravelItemBaseModifyAPIRequest) SetBaseInfo(_baseInfo *PontusTravelItemBaseInfo) error {
	r._baseInfo = _baseInfo
	r.Set("base_info", _baseInfo)
	return nil
}

// Get BaseInfo Getter
func (r TaobaoAlitripTravelItemBaseModifyAPIRequest) GetBaseInfo() *PontusTravelItemBaseInfo {
	return r._baseInfo
}

// Set is Itineraries Setter
// 详细行程描述结构
func (r *TaobaoAlitripTravelItemBaseModifyAPIRequest) SetItineraries(_itineraries []PontusTravelItemItineraryInfo) error {
	r._itineraries = _itineraries
	r.Set("itineraries", _itineraries)
	return nil
}

// Get Itineraries Getter
func (r TaobaoAlitripTravelItemBaseModifyAPIRequest) GetItineraries() []PontusTravelItemItineraryInfo {
	return r._itineraries
}

// Set is ItemId Setter
// 商品id
func (r *TaobaoAlitripTravelItemBaseModifyAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TaobaoAlitripTravelItemBaseModifyAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is RefundInfo Setter
// 退款规则结构
func (r *TaobaoAlitripTravelItemBaseModifyAPIRequest) SetRefundInfo(_refundInfo *PontusTravelItemRefundInfo) error {
	r._refundInfo = _refundInfo
	r.Set("refund_info", _refundInfo)
	return nil
}

// Get RefundInfo Getter
func (r TaobaoAlitripTravelItemBaseModifyAPIRequest) GetRefundInfo() *PontusTravelItemRefundInfo {
	return r._refundInfo
}

// Set is BookingRules Setter
// 预定规则结构
func (r *TaobaoAlitripTravelItemBaseModifyAPIRequest) SetBookingRules(_bookingRules []PontusTravelBookingRuleInfo) error {
	r._bookingRules = _bookingRules
	r.Set("booking_rules", _bookingRules)
	return nil
}

// Get BookingRules Getter
func (r TaobaoAlitripTravelItemBaseModifyAPIRequest) GetBookingRules() []PontusTravelBookingRuleInfo {
	return r._bookingRules
}

// Set is FreedomItemExt Setter
// 自由行商品相关信息
func (r *TaobaoAlitripTravelItemBaseModifyAPIRequest) SetFreedomItemExt(_freedomItemExt *PontusTravelFreedomItemExt) error {
	r._freedomItemExt = _freedomItemExt
	r.Set("freedom_item_ext", _freedomItemExt)
	return nil
}

// Get FreedomItemExt Getter
func (r TaobaoAlitripTravelItemBaseModifyAPIRequest) GetFreedomItemExt() *PontusTravelFreedomItemExt {
	return r._freedomItemExt
}

// Set is SalesInfo Setter
// 商品的销售属性相关信息
func (r *TaobaoAlitripTravelItemBaseModifyAPIRequest) SetSalesInfo(_salesInfo *PontusTravelItemSaleInfo) error {
	r._salesInfo = _salesInfo
	r.Set("sales_info", _salesInfo)
	return nil
}

// Get SalesInfo Getter
func (r TaobaoAlitripTravelItemBaseModifyAPIRequest) GetSalesInfo() *PontusTravelItemSaleInfo {
	return r._salesInfo
}

// Set is FieldsToClean Setter
// 可选，支持清空商品某些字段内容。多个待清空字段名之间以英文逗号分隔。目前支持清除以下字段内容：seller_cids表示删除关联的店铺类目id
func (r *TaobaoAlitripTravelItemBaseModifyAPIRequest) SetFieldsToClean(_fieldsToClean string) error {
	r._fieldsToClean = _fieldsToClean
	r.Set("fields_to_clean", _fieldsToClean)
	return nil
}

// Get FieldsToClean Getter
func (r TaobaoAlitripTravelItemBaseModifyAPIRequest) GetFieldsToClean() string {
	return r._fieldsToClean
}

// Set is CruiseItemExt Setter
// 邮轮商品相关扩展信息
func (r *TaobaoAlitripTravelItemBaseModifyAPIRequest) SetCruiseItemExt(_cruiseItemExt *CruiseItemExt) error {
	r._cruiseItemExt = _cruiseItemExt
	r.Set("cruise_item_ext", _cruiseItemExt)
	return nil
}

// Get CruiseItemExt Getter
func (r TaobaoAlitripTravelItemBaseModifyAPIRequest) GetCruiseItemExt() *CruiseItemExt {
	return r._cruiseItemExt
}

// Set is TcwlItemExt Setter
// 同城活动商品相关扩展信息
func (r *TaobaoAlitripTravelItemBaseModifyAPIRequest) SetTcwlItemExt(_tcwlItemExt *TcwlItemExt) error {
	r._tcwlItemExt = _tcwlItemExt
	r.Set("tcwl_item_ext", _tcwlItemExt)
	return nil
}

// Get TcwlItemExt Getter
func (r TaobaoAlitripTravelItemBaseModifyAPIRequest) GetTcwlItemExt() *TcwlItemExt {
	return r._tcwlItemExt
}
