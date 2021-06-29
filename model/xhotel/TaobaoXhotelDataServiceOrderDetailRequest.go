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
    startDate   string
    // 是否查无订单 1:是 0:否
    isCallNoOrder   int64
    // 酒店id
    hid   int64
    // 是否特殊时段订单 1:是 0:否
    isSpecTimeOrder   int64
    // 渠道商名称
    vendor   string
    // 分页大小
    pageSize   int64
    // 查询结束时间
    endDate   string
    // 是否到店无房 1:是
    isNoRoomCompen   int64
    // 分页参数
    startRow   int64
    // 是否拒单 1:是 0:否
    isSellerDeny   int64
    // 是否卖家原因退款 1:是 0:否
    isSellerRefund   int64
    // 供应商名称
    supplier   string
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
func (r *TaobaoXhotelDataServiceOrderDetailRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

// StartDate Getter
func (r TaobaoXhotelDataServiceOrderDetailRequest) GetStartDate() string {
    return r.startDate
}
// IsCallNoOrder Setter
// 是否查无订单 1:是 0:否
func (r *TaobaoXhotelDataServiceOrderDetailRequest) SetIsCallNoOrder(isCallNoOrder int64) error {
    r.isCallNoOrder = isCallNoOrder
    r.Set("is_call_no_order", isCallNoOrder)
    return nil
}

// IsCallNoOrder Getter
func (r TaobaoXhotelDataServiceOrderDetailRequest) GetIsCallNoOrder() int64 {
    return r.isCallNoOrder
}
// Hid Setter
// 酒店id
func (r *TaobaoXhotelDataServiceOrderDetailRequest) SetHid(hid int64) error {
    r.hid = hid
    r.Set("hid", hid)
    return nil
}

// Hid Getter
func (r TaobaoXhotelDataServiceOrderDetailRequest) GetHid() int64 {
    return r.hid
}
// IsSpecTimeOrder Setter
// 是否特殊时段订单 1:是 0:否
func (r *TaobaoXhotelDataServiceOrderDetailRequest) SetIsSpecTimeOrder(isSpecTimeOrder int64) error {
    r.isSpecTimeOrder = isSpecTimeOrder
    r.Set("is_spec_time_order", isSpecTimeOrder)
    return nil
}

// IsSpecTimeOrder Getter
func (r TaobaoXhotelDataServiceOrderDetailRequest) GetIsSpecTimeOrder() int64 {
    return r.isSpecTimeOrder
}
// Vendor Setter
// 渠道商名称
func (r *TaobaoXhotelDataServiceOrderDetailRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelDataServiceOrderDetailRequest) GetVendor() string {
    return r.vendor
}
// PageSize Setter
// 分页大小
func (r *TaobaoXhotelDataServiceOrderDetailRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoXhotelDataServiceOrderDetailRequest) GetPageSize() int64 {
    return r.pageSize
}
// EndDate Setter
// 查询结束时间
func (r *TaobaoXhotelDataServiceOrderDetailRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r TaobaoXhotelDataServiceOrderDetailRequest) GetEndDate() string {
    return r.endDate
}
// IsNoRoomCompen Setter
// 是否到店无房 1:是
func (r *TaobaoXhotelDataServiceOrderDetailRequest) SetIsNoRoomCompen(isNoRoomCompen int64) error {
    r.isNoRoomCompen = isNoRoomCompen
    r.Set("is_no_room_compen", isNoRoomCompen)
    return nil
}

// IsNoRoomCompen Getter
func (r TaobaoXhotelDataServiceOrderDetailRequest) GetIsNoRoomCompen() int64 {
    return r.isNoRoomCompen
}
// StartRow Setter
// 分页参数
func (r *TaobaoXhotelDataServiceOrderDetailRequest) SetStartRow(startRow int64) error {
    r.startRow = startRow
    r.Set("start_row", startRow)
    return nil
}

// StartRow Getter
func (r TaobaoXhotelDataServiceOrderDetailRequest) GetStartRow() int64 {
    return r.startRow
}
// IsSellerDeny Setter
// 是否拒单 1:是 0:否
func (r *TaobaoXhotelDataServiceOrderDetailRequest) SetIsSellerDeny(isSellerDeny int64) error {
    r.isSellerDeny = isSellerDeny
    r.Set("is_seller_deny", isSellerDeny)
    return nil
}

// IsSellerDeny Getter
func (r TaobaoXhotelDataServiceOrderDetailRequest) GetIsSellerDeny() int64 {
    return r.isSellerDeny
}
// IsSellerRefund Setter
// 是否卖家原因退款 1:是 0:否
func (r *TaobaoXhotelDataServiceOrderDetailRequest) SetIsSellerRefund(isSellerRefund int64) error {
    r.isSellerRefund = isSellerRefund
    r.Set("is_seller_refund", isSellerRefund)
    return nil
}

// IsSellerRefund Getter
func (r TaobaoXhotelDataServiceOrderDetailRequest) GetIsSellerRefund() int64 {
    return r.isSellerRefund
}
// Supplier Setter
// 供应商名称
func (r *TaobaoXhotelDataServiceOrderDetailRequest) SetSupplier(supplier string) error {
    r.supplier = supplier
    r.Set("supplier", supplier)
    return nil
}

// Supplier Getter
func (r TaobaoXhotelDataServiceOrderDetailRequest) GetSupplier() string {
    return r.supplier
}
