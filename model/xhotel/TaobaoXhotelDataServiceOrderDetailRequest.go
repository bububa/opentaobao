package xhotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务订单详情 API请求
taobao.xhotel.data.service.order.detail

服务订单详情top接口构建
*/
type TaobaoXhotelDataServiceOrderDetailRequest struct {
    model.Params
    // 查询开始日期
    _startDate   string
    // 是否查无订单 1:是 0:否
    _isCallNoOrder   int64
    // 酒店id
    _hid   int64
    // 是否特殊时段订单 1:是 0:否
    _isSpecTimeOrder   int64
    // 渠道商名称
    _vendor   string
    // 分页大小
    _pageSize   int64
    // 查询结束时间
    _endDate   string
    // 是否到店无房 1:是
    _isNoRoomCompen   int64
    // 分页参数
    _startRow   int64
    // 是否拒单 1:是 0:否
    _isSellerDeny   int64
    // 是否卖家原因退款 1:是 0:否
    _isSellerRefund   int64
    // 供应商名称
    _supplier   string
}

// 初始化TaobaoXhotelDataServiceOrderDetailRequest对象
func NewTaobaoXhotelDataServiceOrderDetailRequest() *TaobaoXhotelDataServiceOrderDetailRequest{
    return &TaobaoXhotelDataServiceOrderDetailRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelDataServiceOrderDetailRequest) GetApiMethodName() string {
    return "taobao.xhotel.data.service.order.detail"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelDataServiceOrderDetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartDate Setter
// 查询开始日期
func (r *TaobaoXhotelDataServiceOrderDetailRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r TaobaoXhotelDataServiceOrderDetailRequest) GetStartDate() string {
    return r._startDate
}
// IsCallNoOrder Setter
// 是否查无订单 1:是 0:否
func (r *TaobaoXhotelDataServiceOrderDetailRequest) SetIsCallNoOrder(_isCallNoOrder int64) error {
    r._isCallNoOrder = _isCallNoOrder
    r.Set("is_call_no_order", _isCallNoOrder)
    return nil
}

// IsCallNoOrder Getter
func (r TaobaoXhotelDataServiceOrderDetailRequest) GetIsCallNoOrder() int64 {
    return r._isCallNoOrder
}
// Hid Setter
// 酒店id
func (r *TaobaoXhotelDataServiceOrderDetailRequest) SetHid(_hid int64) error {
    r._hid = _hid
    r.Set("hid", _hid)
    return nil
}

// Hid Getter
func (r TaobaoXhotelDataServiceOrderDetailRequest) GetHid() int64 {
    return r._hid
}
// IsSpecTimeOrder Setter
// 是否特殊时段订单 1:是 0:否
func (r *TaobaoXhotelDataServiceOrderDetailRequest) SetIsSpecTimeOrder(_isSpecTimeOrder int64) error {
    r._isSpecTimeOrder = _isSpecTimeOrder
    r.Set("is_spec_time_order", _isSpecTimeOrder)
    return nil
}

// IsSpecTimeOrder Getter
func (r TaobaoXhotelDataServiceOrderDetailRequest) GetIsSpecTimeOrder() int64 {
    return r._isSpecTimeOrder
}
// Vendor Setter
// 渠道商名称
func (r *TaobaoXhotelDataServiceOrderDetailRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelDataServiceOrderDetailRequest) GetVendor() string {
    return r._vendor
}
// PageSize Setter
// 分页大小
func (r *TaobaoXhotelDataServiceOrderDetailRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoXhotelDataServiceOrderDetailRequest) GetPageSize() int64 {
    return r._pageSize
}
// EndDate Setter
// 查询结束时间
func (r *TaobaoXhotelDataServiceOrderDetailRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r TaobaoXhotelDataServiceOrderDetailRequest) GetEndDate() string {
    return r._endDate
}
// IsNoRoomCompen Setter
// 是否到店无房 1:是
func (r *TaobaoXhotelDataServiceOrderDetailRequest) SetIsNoRoomCompen(_isNoRoomCompen int64) error {
    r._isNoRoomCompen = _isNoRoomCompen
    r.Set("is_no_room_compen", _isNoRoomCompen)
    return nil
}

// IsNoRoomCompen Getter
func (r TaobaoXhotelDataServiceOrderDetailRequest) GetIsNoRoomCompen() int64 {
    return r._isNoRoomCompen
}
// StartRow Setter
// 分页参数
func (r *TaobaoXhotelDataServiceOrderDetailRequest) SetStartRow(_startRow int64) error {
    r._startRow = _startRow
    r.Set("start_row", _startRow)
    return nil
}

// StartRow Getter
func (r TaobaoXhotelDataServiceOrderDetailRequest) GetStartRow() int64 {
    return r._startRow
}
// IsSellerDeny Setter
// 是否拒单 1:是 0:否
func (r *TaobaoXhotelDataServiceOrderDetailRequest) SetIsSellerDeny(_isSellerDeny int64) error {
    r._isSellerDeny = _isSellerDeny
    r.Set("is_seller_deny", _isSellerDeny)
    return nil
}

// IsSellerDeny Getter
func (r TaobaoXhotelDataServiceOrderDetailRequest) GetIsSellerDeny() int64 {
    return r._isSellerDeny
}
// IsSellerRefund Setter
// 是否卖家原因退款 1:是 0:否
func (r *TaobaoXhotelDataServiceOrderDetailRequest) SetIsSellerRefund(_isSellerRefund int64) error {
    r._isSellerRefund = _isSellerRefund
    r.Set("is_seller_refund", _isSellerRefund)
    return nil
}

// IsSellerRefund Getter
func (r TaobaoXhotelDataServiceOrderDetailRequest) GetIsSellerRefund() int64 {
    return r._isSellerRefund
}
// Supplier Setter
// 供应商名称
func (r *TaobaoXhotelDataServiceOrderDetailRequest) SetSupplier(_supplier string) error {
    r._supplier = _supplier
    r.Set("supplier", _supplier)
    return nil
}

// Supplier Getter
func (r TaobaoXhotelDataServiceOrderDetailRequest) GetSupplier() string {
    return r._supplier
}
