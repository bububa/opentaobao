package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelGereralitemUpdateAPIRequest 除度假线路、门票以外的其他类目商品维护接口（商品ID重复将自动更新） API请求
// alitrip.travel.gereralitem.update
//
// 除度假线路、门票以外的商品维护接口；目前该接口支持以下类目；
// （123740001：客栈周边交通服务、125038002：旅行设备/GPS/相机租赁、50018298：船票、124084006：酒店商品升级差价、125228016：预约卡券、50011954：旅游服务、50012913：酒店优惠券、50214003：旅游会员卡/酒店会员卡、50012917：巴士/地铁/交通卡/一卡通、50134002：代客烧香/还愿、50026091：境外火车票、123742001：客栈周边餐饮服务、50019817：海外套餐（该类目已废弃）、125210016：团建/outing、124212017：其他预定、50025888：机场行李托运取送寄存、50025831：旅游年票/年卡、124142009：旅游商品升级差价/押金、123744001：客栈周边其他服务、50012762：广深口岸港澳送关服务、50025880：旅行拍照/婚纱摄影、123166001：酒店餐饮美食（该类目已废弃）、50668002：手绘地图/明信片、50024210：旅游购物/纪念品、50024208：酒店用品、50024215：购物折扣卡券、50025878：酒店SPA/足浴/温泉、50024212：旅游必备品、123738001：客栈周边票务服务、123164002：游泳健身（该类目已废弃）、50686003：机票增值产品、123164001：酒店SPA（该类目已废弃）、124832008：美食卡券/酒店餐饮卡券、125408001：旅游定制服务、50018112：旅行社/网站优惠券、124258004：酒店客房优惠券（该类目已废弃）、50104001：机场周边停车位、124730002：内机机场/火车站送关服务、124090010：其他）
type AlitripTravelGereralitemUpdateAPIRequest struct {
	model.Params
	// 必填，预定规则结构。示例： [{ "rule_type": "fee_excluded", "rule_desc": "费用包含描述"},{ "rule_type": "fee_included", "rule_desc": "费用不含描述"},{ "rule_type": "order_info", "rule_desc": "预定须知描述"}]
	_bookingRules []BookingRuleInfo
	// 更新sku信息，仅限日历商品使用
	_dateSkuInfoList []DateSkuInfo
	// 更新sku信息，仅限非日历（普通）商品使用
	_commonSkuList []NoDateSkuInfo
	// 必填，商品基本信息
	_baseInfo *BaseInfo
	// 新版电子凭证信息。如果传递了此参数，则sales_info中旧版电子凭证信息将被忽略，否则电子凭证信息将以旧版电子凭证参数为准。如果新、旧版参数都没传，则该商品不支持电子凭证
	_itemEleCertInfo *ItemEleCertInfo
	// 选填，退款规则结构
	_itemRefundInfo *ItemRefundInfo
	// poi信息，个别类目必填，如演艺类目下场馆信息
	_poi *Poi
}

// NewAlitripTravelGereralitemUpdateRequest 初始化AlitripTravelGereralitemUpdateAPIRequest对象
func NewAlitripTravelGereralitemUpdateRequest() *AlitripTravelGereralitemUpdateAPIRequest {
	return &AlitripTravelGereralitemUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTravelGereralitemUpdateAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.gereralitem.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTravelGereralitemUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBookingRules is BookingRules Setter
// 必填，预定规则结构。示例： [{ "rule_type": "fee_excluded", "rule_desc": "费用包含描述"},{ "rule_type": "fee_included", "rule_desc": "费用不含描述"},{ "rule_type": "order_info", "rule_desc": "预定须知描述"}]
func (r *AlitripTravelGereralitemUpdateAPIRequest) SetBookingRules(_bookingRules []BookingRuleInfo) error {
	r._bookingRules = _bookingRules
	r.Set("booking_rules", _bookingRules)
	return nil
}

// GetBookingRules BookingRules Getter
func (r AlitripTravelGereralitemUpdateAPIRequest) GetBookingRules() []BookingRuleInfo {
	return r._bookingRules
}

// SetDateSkuInfoList is DateSkuInfoList Setter
// 更新sku信息，仅限日历商品使用
func (r *AlitripTravelGereralitemUpdateAPIRequest) SetDateSkuInfoList(_dateSkuInfoList []DateSkuInfo) error {
	r._dateSkuInfoList = _dateSkuInfoList
	r.Set("date_sku_info_list", _dateSkuInfoList)
	return nil
}

// GetDateSkuInfoList DateSkuInfoList Getter
func (r AlitripTravelGereralitemUpdateAPIRequest) GetDateSkuInfoList() []DateSkuInfo {
	return r._dateSkuInfoList
}

// SetCommonSkuList is CommonSkuList Setter
// 更新sku信息，仅限非日历（普通）商品使用
func (r *AlitripTravelGereralitemUpdateAPIRequest) SetCommonSkuList(_commonSkuList []NoDateSkuInfo) error {
	r._commonSkuList = _commonSkuList
	r.Set("common_sku_list", _commonSkuList)
	return nil
}

// GetCommonSkuList CommonSkuList Getter
func (r AlitripTravelGereralitemUpdateAPIRequest) GetCommonSkuList() []NoDateSkuInfo {
	return r._commonSkuList
}

// SetBaseInfo is BaseInfo Setter
// 必填，商品基本信息
func (r *AlitripTravelGereralitemUpdateAPIRequest) SetBaseInfo(_baseInfo *BaseInfo) error {
	r._baseInfo = _baseInfo
	r.Set("base_info", _baseInfo)
	return nil
}

// GetBaseInfo BaseInfo Getter
func (r AlitripTravelGereralitemUpdateAPIRequest) GetBaseInfo() *BaseInfo {
	return r._baseInfo
}

// SetItemEleCertInfo is ItemEleCertInfo Setter
// 新版电子凭证信息。如果传递了此参数，则sales_info中旧版电子凭证信息将被忽略，否则电子凭证信息将以旧版电子凭证参数为准。如果新、旧版参数都没传，则该商品不支持电子凭证
func (r *AlitripTravelGereralitemUpdateAPIRequest) SetItemEleCertInfo(_itemEleCertInfo *ItemEleCertInfo) error {
	r._itemEleCertInfo = _itemEleCertInfo
	r.Set("item_ele_cert_info", _itemEleCertInfo)
	return nil
}

// GetItemEleCertInfo ItemEleCertInfo Getter
func (r AlitripTravelGereralitemUpdateAPIRequest) GetItemEleCertInfo() *ItemEleCertInfo {
	return r._itemEleCertInfo
}

// SetItemRefundInfo is ItemRefundInfo Setter
// 选填，退款规则结构
func (r *AlitripTravelGereralitemUpdateAPIRequest) SetItemRefundInfo(_itemRefundInfo *ItemRefundInfo) error {
	r._itemRefundInfo = _itemRefundInfo
	r.Set("item_refund_info", _itemRefundInfo)
	return nil
}

// GetItemRefundInfo ItemRefundInfo Getter
func (r AlitripTravelGereralitemUpdateAPIRequest) GetItemRefundInfo() *ItemRefundInfo {
	return r._itemRefundInfo
}

// SetPoi is Poi Setter
// poi信息，个别类目必填，如演艺类目下场馆信息
func (r *AlitripTravelGereralitemUpdateAPIRequest) SetPoi(_poi *Poi) error {
	r._poi = _poi
	r.Set("poi", _poi)
	return nil
}

// GetPoi Poi Getter
func (r AlitripTravelGereralitemUpdateAPIRequest) GetPoi() *Poi {
	return r._poi
}
