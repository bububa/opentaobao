package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询账单信息 API请求
taobao.xhotel.order.statement.get

阿里根据此接口定义输出订单账务明细，结账状态发生变化时阿里需推送账单信息。系统商可实时调用该接口来查询订单的详情
*/
type TaobaoXhotelOrderStatementGetRequest struct {
    model.Params
    // 要查询的tid列表，逗号分隔,列表查询;当此值不为空时候，其余参数忽略。最多单次20条。
    _orderTids   string
    // 查询条数，最大支持500条
    _pageSize   int64
    // 数据查询开始下标
    _start   int64
    // 0：check_in, 1：check_out,2：分账时间
    _dateType   int64
    // 查询结束时间
    _to   string
    // 查询开始时间
    _from   string
    // 淘宝订单号
    _tid   int64
    // 外部酒店编码
    _hotelCode   string
    // 系统商vendor
    _vendor   string
}

// 初始化TaobaoXhotelOrderStatementGetRequest对象
func NewTaobaoXhotelOrderStatementGetRequest() *TaobaoXhotelOrderStatementGetRequest{
    return &TaobaoXhotelOrderStatementGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderStatementGetRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.statement.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderStatementGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderTids Setter
// 要查询的tid列表，逗号分隔,列表查询;当此值不为空时候，其余参数忽略。最多单次20条。
func (r *TaobaoXhotelOrderStatementGetRequest) SetOrderTids(_orderTids string) error {
    r._orderTids = _orderTids
    r.Set("order_tids", _orderTids)
    return nil
}

// OrderTids Getter
func (r TaobaoXhotelOrderStatementGetRequest) GetOrderTids() string {
    return r._orderTids
}
// PageSize Setter
// 查询条数，最大支持500条
func (r *TaobaoXhotelOrderStatementGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoXhotelOrderStatementGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// Start Setter
// 数据查询开始下标
func (r *TaobaoXhotelOrderStatementGetRequest) SetStart(_start int64) error {
    r._start = _start
    r.Set("start", _start)
    return nil
}

// Start Getter
func (r TaobaoXhotelOrderStatementGetRequest) GetStart() int64 {
    return r._start
}
// DateType Setter
// 0：check_in, 1：check_out,2：分账时间
func (r *TaobaoXhotelOrderStatementGetRequest) SetDateType(_dateType int64) error {
    r._dateType = _dateType
    r.Set("date_type", _dateType)
    return nil
}

// DateType Getter
func (r TaobaoXhotelOrderStatementGetRequest) GetDateType() int64 {
    return r._dateType
}
// To Setter
// 查询结束时间
func (r *TaobaoXhotelOrderStatementGetRequest) SetTo(_to string) error {
    r._to = _to
    r.Set("to", _to)
    return nil
}

// To Getter
func (r TaobaoXhotelOrderStatementGetRequest) GetTo() string {
    return r._to
}
// From Setter
// 查询开始时间
func (r *TaobaoXhotelOrderStatementGetRequest) SetFrom(_from string) error {
    r._from = _from
    r.Set("from", _from)
    return nil
}

// From Getter
func (r TaobaoXhotelOrderStatementGetRequest) GetFrom() string {
    return r._from
}
// Tid Setter
// 淘宝订单号
func (r *TaobaoXhotelOrderStatementGetRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoXhotelOrderStatementGetRequest) GetTid() int64 {
    return r._tid
}
// HotelCode Setter
// 外部酒店编码
func (r *TaobaoXhotelOrderStatementGetRequest) SetHotelCode(_hotelCode string) error {
    r._hotelCode = _hotelCode
    r.Set("hotel_code", _hotelCode)
    return nil
}

// HotelCode Getter
func (r TaobaoXhotelOrderStatementGetRequest) GetHotelCode() string {
    return r._hotelCode
}
// Vendor Setter
// 系统商vendor
func (r *TaobaoXhotelOrderStatementGetRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelOrderStatementGetRequest) GetVendor() string {
    return r._vendor
}
