package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelordersearchAPIRequest 酒店产品库订单查询 API请求
// taobao.xhotel.order.search
//
// 酒店产品库订单查询功能，查询90天内的订单
type TaobaoxhotelordersearchAPIRequest struct {
	model.Params
	// 酒店订单oids列表，多个oid用英文逗号隔开，一次不超过20个。
	_orderIds string
	// 酒店订单tids列表，多个tid用英文逗号隔开，一次不超过20个。oids和tids都传的情况下默认使用tids
	_orderTids string
	// 订单创建时间查询起始时间，格式为：yyyy-MM-dd HH:mm:ss
	_createdStart string
	// 订单创建时间查询结束时间，格式为：yyyy-MM-dd HH:mm:ss。不能早于created_start或者间隔不能大于30
	_createdEnd string
	// 系统商标识
	_vendor string
	// 外部订单号out_oids列表，多个oid用英文逗号隔开，一次不超过20个。
	_outOids string
	// 入住时间查询起始时间，格式为：yyyy-MM-dd HH:mm:ss
	_checkinDateStart string
	// 入住时间查询结束时间，格式为：yyyy-MM-dd HH:mm:ss。不能早于checkin_date_start或者间隔不能大于30
	_checkinDateEnd string
	// 离店时间查询起始时间，格式为：yyyy-MM-dd HH:mm:ss
	_checkoutDateStart string
	// 离店时间查询结束时间，格式为：yyyy-MM-dd HH:mm:ss。不能早于checkout_date_start或者间隔不能大于30
	_checkoutDateEnd string
	// 订单状态（可发多个，逗号隔开）
	_tradeStatus string
	// 分页页码。取值范围，大于零的整数，默认值1，即返回第一页的数据。页面大小为20
	_pageNo int64
	// 订单类型（true为直连，false为非直连）
	_direct bool
}

// NewTaobaoxhotelordersearchRequest 初始化TaobaoxhotelordersearchAPIRequest对象
func NewTaobaoxhotelordersearchRequest() *TaobaoxhotelordersearchAPIRequest {
	return &TaobaoxhotelordersearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelordersearchAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.order.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelordersearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelordersearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderIds is OrderIds Setter
// 酒店订单oids列表，多个oid用英文逗号隔开，一次不超过20个。
func (r *TaobaoxhotelordersearchAPIRequest) SetOrderIds(_orderIds string) error {
	r._orderIds = _orderIds
	r.Set("order_ids", _orderIds)
	return nil
}

// GetOrderIds OrderIds Getter
func (r TaobaoxhotelordersearchAPIRequest) GetOrderIds() string {
	return r._orderIds
}

// SetOrderTids is OrderTids Setter
// 酒店订单tids列表，多个tid用英文逗号隔开，一次不超过20个。oids和tids都传的情况下默认使用tids
func (r *TaobaoxhotelordersearchAPIRequest) SetOrderTids(_orderTids string) error {
	r._orderTids = _orderTids
	r.Set("order_tids", _orderTids)
	return nil
}

// GetOrderTids OrderTids Getter
func (r TaobaoxhotelordersearchAPIRequest) GetOrderTids() string {
	return r._orderTids
}

// SetCreatedStart is CreatedStart Setter
// 订单创建时间查询起始时间，格式为：yyyy-MM-dd HH:mm:ss
func (r *TaobaoxhotelordersearchAPIRequest) SetCreatedStart(_createdStart string) error {
	r._createdStart = _createdStart
	r.Set("created_start", _createdStart)
	return nil
}

// GetCreatedStart CreatedStart Getter
func (r TaobaoxhotelordersearchAPIRequest) GetCreatedStart() string {
	return r._createdStart
}

// SetCreatedEnd is CreatedEnd Setter
// 订单创建时间查询结束时间，格式为：yyyy-MM-dd HH:mm:ss。不能早于created_start或者间隔不能大于30
func (r *TaobaoxhotelordersearchAPIRequest) SetCreatedEnd(_createdEnd string) error {
	r._createdEnd = _createdEnd
	r.Set("created_end", _createdEnd)
	return nil
}

// GetCreatedEnd CreatedEnd Getter
func (r TaobaoxhotelordersearchAPIRequest) GetCreatedEnd() string {
	return r._createdEnd
}

// SetVendor is Vendor Setter
// 系统商标识
func (r *TaobaoxhotelordersearchAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoxhotelordersearchAPIRequest) GetVendor() string {
	return r._vendor
}

// SetOutOids is OutOids Setter
// 外部订单号out_oids列表，多个oid用英文逗号隔开，一次不超过20个。
func (r *TaobaoxhotelordersearchAPIRequest) SetOutOids(_outOids string) error {
	r._outOids = _outOids
	r.Set("out_oids", _outOids)
	return nil
}

// GetOutOids OutOids Getter
func (r TaobaoxhotelordersearchAPIRequest) GetOutOids() string {
	return r._outOids
}

// SetCheckinDateStart is CheckinDateStart Setter
// 入住时间查询起始时间，格式为：yyyy-MM-dd HH:mm:ss
func (r *TaobaoxhotelordersearchAPIRequest) SetCheckinDateStart(_checkinDateStart string) error {
	r._checkinDateStart = _checkinDateStart
	r.Set("checkin_date_start", _checkinDateStart)
	return nil
}

// GetCheckinDateStart CheckinDateStart Getter
func (r TaobaoxhotelordersearchAPIRequest) GetCheckinDateStart() string {
	return r._checkinDateStart
}

// SetCheckinDateEnd is CheckinDateEnd Setter
// 入住时间查询结束时间，格式为：yyyy-MM-dd HH:mm:ss。不能早于checkin_date_start或者间隔不能大于30
func (r *TaobaoxhotelordersearchAPIRequest) SetCheckinDateEnd(_checkinDateEnd string) error {
	r._checkinDateEnd = _checkinDateEnd
	r.Set("checkin_date_end", _checkinDateEnd)
	return nil
}

// GetCheckinDateEnd CheckinDateEnd Getter
func (r TaobaoxhotelordersearchAPIRequest) GetCheckinDateEnd() string {
	return r._checkinDateEnd
}

// SetCheckoutDateStart is CheckoutDateStart Setter
// 离店时间查询起始时间，格式为：yyyy-MM-dd HH:mm:ss
func (r *TaobaoxhotelordersearchAPIRequest) SetCheckoutDateStart(_checkoutDateStart string) error {
	r._checkoutDateStart = _checkoutDateStart
	r.Set("checkout_date_start", _checkoutDateStart)
	return nil
}

// GetCheckoutDateStart CheckoutDateStart Getter
func (r TaobaoxhotelordersearchAPIRequest) GetCheckoutDateStart() string {
	return r._checkoutDateStart
}

// SetCheckoutDateEnd is CheckoutDateEnd Setter
// 离店时间查询结束时间，格式为：yyyy-MM-dd HH:mm:ss。不能早于checkout_date_start或者间隔不能大于30
func (r *TaobaoxhotelordersearchAPIRequest) SetCheckoutDateEnd(_checkoutDateEnd string) error {
	r._checkoutDateEnd = _checkoutDateEnd
	r.Set("checkout_date_end", _checkoutDateEnd)
	return nil
}

// GetCheckoutDateEnd CheckoutDateEnd Getter
func (r TaobaoxhotelordersearchAPIRequest) GetCheckoutDateEnd() string {
	return r._checkoutDateEnd
}

// SetTradeStatus is TradeStatus Setter
// 订单状态（可发多个，逗号隔开）
func (r *TaobaoxhotelordersearchAPIRequest) SetTradeStatus(_tradeStatus string) error {
	r._tradeStatus = _tradeStatus
	r.Set("trade_status", _tradeStatus)
	return nil
}

// GetTradeStatus TradeStatus Getter
func (r TaobaoxhotelordersearchAPIRequest) GetTradeStatus() string {
	return r._tradeStatus
}

// SetPageNo is PageNo Setter
// 分页页码。取值范围，大于零的整数，默认值1，即返回第一页的数据。页面大小为20
func (r *TaobaoxhotelordersearchAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoxhotelordersearchAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetDirect is Direct Setter
// 订单类型（true为直连，false为非直连）
func (r *TaobaoxhotelordersearchAPIRequest) SetDirect(_direct bool) error {
	r._direct = _direct
	r.Set("direct", _direct)
	return nil
}

// GetDirect Direct Getter
func (r TaobaoxhotelordersearchAPIRequest) GetDirect() bool {
	return r._direct
}
