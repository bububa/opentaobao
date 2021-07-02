package xhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelDataServiceOrderDetailAPIRequest 服务订单详情 API请求
// taobao.xhotel.data.service.order.detail
//
// 服务订单详情top接口构建
type TaobaoXhotelDataServiceOrderDetailAPIRequest struct {
	model.Params
	// 查询开始日期
	_startDate string
	// 是否查无订单 1:是 0:否
	_isCallNoOrder int64
	// 酒店id
	_hid int64
	// 是否特殊时段订单 1:是 0:否
	_isSpecTimeOrder int64
	// 渠道商名称
	_vendor string
	// 分页大小
	_pageSize int64
	// 查询结束时间
	_endDate string
	// 是否到店无房 1:是
	_isNoRoomCompen int64
	// 分页参数
	_startRow int64
	// 是否拒单 1:是 0:否
	_isSellerDeny int64
	// 是否卖家原因退款 1:是 0:否
	_isSellerRefund int64
	// 供应商名称
	_supplier string
}

// NewTaobaoXhotelDataServiceOrderDetailRequest 初始化TaobaoXhotelDataServiceOrderDetailAPIRequest对象
func NewTaobaoXhotelDataServiceOrderDetailRequest() *TaobaoXhotelDataServiceOrderDetailAPIRequest {
	return &TaobaoXhotelDataServiceOrderDetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelDataServiceOrderDetailAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.data.service.order.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelDataServiceOrderDetailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetStartDate is StartDate Setter
// 查询开始日期
func (r *TaobaoXhotelDataServiceOrderDetailAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaoXhotelDataServiceOrderDetailAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetIsCallNoOrder is IsCallNoOrder Setter
// 是否查无订单 1:是 0:否
func (r *TaobaoXhotelDataServiceOrderDetailAPIRequest) SetIsCallNoOrder(_isCallNoOrder int64) error {
	r._isCallNoOrder = _isCallNoOrder
	r.Set("is_call_no_order", _isCallNoOrder)
	return nil
}

// GetIsCallNoOrder IsCallNoOrder Getter
func (r TaobaoXhotelDataServiceOrderDetailAPIRequest) GetIsCallNoOrder() int64 {
	return r._isCallNoOrder
}

// SetHid is Hid Setter
// 酒店id
func (r *TaobaoXhotelDataServiceOrderDetailAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r TaobaoXhotelDataServiceOrderDetailAPIRequest) GetHid() int64 {
	return r._hid
}

// SetIsSpecTimeOrder is IsSpecTimeOrder Setter
// 是否特殊时段订单 1:是 0:否
func (r *TaobaoXhotelDataServiceOrderDetailAPIRequest) SetIsSpecTimeOrder(_isSpecTimeOrder int64) error {
	r._isSpecTimeOrder = _isSpecTimeOrder
	r.Set("is_spec_time_order", _isSpecTimeOrder)
	return nil
}

// GetIsSpecTimeOrder IsSpecTimeOrder Getter
func (r TaobaoXhotelDataServiceOrderDetailAPIRequest) GetIsSpecTimeOrder() int64 {
	return r._isSpecTimeOrder
}

// SetVendor is Vendor Setter
// 渠道商名称
func (r *TaobaoXhotelDataServiceOrderDetailAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelDataServiceOrderDetailAPIRequest) GetVendor() string {
	return r._vendor
}

// SetPageSize is PageSize Setter
// 分页大小
func (r *TaobaoXhotelDataServiceOrderDetailAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoXhotelDataServiceOrderDetailAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetEndDate is EndDate Setter
// 查询结束时间
func (r *TaobaoXhotelDataServiceOrderDetailAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaoXhotelDataServiceOrderDetailAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetIsNoRoomCompen is IsNoRoomCompen Setter
// 是否到店无房 1:是
func (r *TaobaoXhotelDataServiceOrderDetailAPIRequest) SetIsNoRoomCompen(_isNoRoomCompen int64) error {
	r._isNoRoomCompen = _isNoRoomCompen
	r.Set("is_no_room_compen", _isNoRoomCompen)
	return nil
}

// GetIsNoRoomCompen IsNoRoomCompen Getter
func (r TaobaoXhotelDataServiceOrderDetailAPIRequest) GetIsNoRoomCompen() int64 {
	return r._isNoRoomCompen
}

// SetStartRow is StartRow Setter
// 分页参数
func (r *TaobaoXhotelDataServiceOrderDetailAPIRequest) SetStartRow(_startRow int64) error {
	r._startRow = _startRow
	r.Set("start_row", _startRow)
	return nil
}

// GetStartRow StartRow Getter
func (r TaobaoXhotelDataServiceOrderDetailAPIRequest) GetStartRow() int64 {
	return r._startRow
}

// SetIsSellerDeny is IsSellerDeny Setter
// 是否拒单 1:是 0:否
func (r *TaobaoXhotelDataServiceOrderDetailAPIRequest) SetIsSellerDeny(_isSellerDeny int64) error {
	r._isSellerDeny = _isSellerDeny
	r.Set("is_seller_deny", _isSellerDeny)
	return nil
}

// GetIsSellerDeny IsSellerDeny Getter
func (r TaobaoXhotelDataServiceOrderDetailAPIRequest) GetIsSellerDeny() int64 {
	return r._isSellerDeny
}

// SetIsSellerRefund is IsSellerRefund Setter
// 是否卖家原因退款 1:是 0:否
func (r *TaobaoXhotelDataServiceOrderDetailAPIRequest) SetIsSellerRefund(_isSellerRefund int64) error {
	r._isSellerRefund = _isSellerRefund
	r.Set("is_seller_refund", _isSellerRefund)
	return nil
}

// GetIsSellerRefund IsSellerRefund Getter
func (r TaobaoXhotelDataServiceOrderDetailAPIRequest) GetIsSellerRefund() int64 {
	return r._isSellerRefund
}

// SetSupplier is Supplier Setter
// 供应商名称
func (r *TaobaoXhotelDataServiceOrderDetailAPIRequest) SetSupplier(_supplier string) error {
	r._supplier = _supplier
	r.Set("supplier", _supplier)
	return nil
}

// GetSupplier Supplier Getter
func (r TaobaoXhotelDataServiceOrderDetailAPIRequest) GetSupplier() string {
	return r._supplier
}
