package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
历史订单查询（对账） API请求
taobao.bus.historyorder.get

历史订单查询，对账接口
*/
type TaobaoBusHistoryorderGetRequest struct {
    model.Params
    // 开始时间 2017-04-23 13:33:43
    _fromDate   string
    // 分页大小 不超过1w
    _pageSize   int64
    // 结束时间 2017-04-23 13:33:43
    _toDate   string
    // offline_ticket 线下自助机； online_ticket：线上售票； 空 代表查全部
    _type   string
    // 第几页 从1开始
    _pageIndex   int64
}

// 初始化TaobaoBusHistoryorderGetRequest对象
func NewTaobaoBusHistoryorderGetRequest() *TaobaoBusHistoryorderGetRequest{
    return &TaobaoBusHistoryorderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusHistoryorderGetRequest) GetApiMethodName() string {
    return "taobao.bus.historyorder.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusHistoryorderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FromDate Setter
// 开始时间 2017-04-23 13:33:43
func (r *TaobaoBusHistoryorderGetRequest) SetFromDate(_fromDate string) error {
    r._fromDate = _fromDate
    r.Set("from_date", _fromDate)
    return nil
}

// FromDate Getter
func (r TaobaoBusHistoryorderGetRequest) GetFromDate() string {
    return r._fromDate
}
// PageSize Setter
// 分页大小 不超过1w
func (r *TaobaoBusHistoryorderGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoBusHistoryorderGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// ToDate Setter
// 结束时间 2017-04-23 13:33:43
func (r *TaobaoBusHistoryorderGetRequest) SetToDate(_toDate string) error {
    r._toDate = _toDate
    r.Set("to_date", _toDate)
    return nil
}

// ToDate Getter
func (r TaobaoBusHistoryorderGetRequest) GetToDate() string {
    return r._toDate
}
// Type Setter
// offline_ticket 线下自助机； online_ticket：线上售票； 空 代表查全部
func (r *TaobaoBusHistoryorderGetRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoBusHistoryorderGetRequest) GetType() string {
    return r._type
}
// PageIndex Setter
// 第几页 从1开始
func (r *TaobaoBusHistoryorderGetRequest) SetPageIndex(_pageIndex int64) error {
    r._pageIndex = _pageIndex
    r.Set("page_index", _pageIndex)
    return nil
}

// PageIndex Getter
func (r TaobaoBusHistoryorderGetRequest) GetPageIndex() int64 {
    return r._pageIndex
}
