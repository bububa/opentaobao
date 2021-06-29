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
    fromDate   string
    // 分页大小 不超过1w
    pageSize   int64
    // 结束时间 2017-04-23 13:33:43
    toDate   string
    // offline_ticket 线下自助机； online_ticket：线上售票； 空 代表查全部
    type   string
    // 第几页 从1开始
    pageIndex   int64
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
func (r *TaobaoBusHistoryorderGetRequest) SetFromDate(fromDate string) error {
    r.fromDate = fromDate
    r.Set("from_date", fromDate)
    return nil
}

// FromDate Getter
func (r TaobaoBusHistoryorderGetRequest) GetFromDate() string {
    return r.fromDate
}
// PageSize Setter
// 分页大小 不超过1w
func (r *TaobaoBusHistoryorderGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoBusHistoryorderGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// ToDate Setter
// 结束时间 2017-04-23 13:33:43
func (r *TaobaoBusHistoryorderGetRequest) SetToDate(toDate string) error {
    r.toDate = toDate
    r.Set("to_date", toDate)
    return nil
}

// ToDate Getter
func (r TaobaoBusHistoryorderGetRequest) GetToDate() string {
    return r.toDate
}
// Type Setter
// offline_ticket 线下自助机； online_ticket：线上售票； 空 代表查全部
func (r *TaobaoBusHistoryorderGetRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r TaobaoBusHistoryorderGetRequest) GetType() string {
    return r.type
}
// PageIndex Setter
// 第几页 从1开始
func (r *TaobaoBusHistoryorderGetRequest) SetPageIndex(pageIndex int64) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

// PageIndex Getter
func (r TaobaoBusHistoryorderGetRequest) GetPageIndex() int64 {
    return r.pageIndex
}
