package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
核销单查询 API请求
alibaba.servicecenter.fulfiltask.query

当系统生成核销单之后，需要派单到服务商，服务商根据核销里的服务信息和用户信息，给消费者提供服务
*/
type AlibabaServicecenterFulfiltaskQueryRequest struct {
    model.Params
    // 时间段查询，核销单创建时间，时间段跨度不超过15分钟
    _gmtCreateStart   string
    // 每页条数，默认50，最大50
    _pageSize   int64
    // 核销单外部单号
    _outerId   string
    // 时间段查询，核销单创建时间，时间段跨度不超过15分钟
    _gmtCreateEnd   string
    // 查询第几页
    _currentPage   int64
    // 核销单id列表
    _fulfilTaskIdList   []int64
}

// 初始化AlibabaServicecenterFulfiltaskQueryRequest对象
func NewAlibabaServicecenterFulfiltaskQueryRequest() *AlibabaServicecenterFulfiltaskQueryRequest{
    return &AlibabaServicecenterFulfiltaskQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaServicecenterFulfiltaskQueryRequest) GetApiMethodName() string {
    return "alibaba.servicecenter.fulfiltask.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaServicecenterFulfiltaskQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GmtCreateStart Setter
// 时间段查询，核销单创建时间，时间段跨度不超过15分钟
func (r *AlibabaServicecenterFulfiltaskQueryRequest) SetGmtCreateStart(_gmtCreateStart string) error {
    r._gmtCreateStart = _gmtCreateStart
    r.Set("gmt_create_start", _gmtCreateStart)
    return nil
}

// GmtCreateStart Getter
func (r AlibabaServicecenterFulfiltaskQueryRequest) GetGmtCreateStart() string {
    return r._gmtCreateStart
}
// PageSize Setter
// 每页条数，默认50，最大50
func (r *AlibabaServicecenterFulfiltaskQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaServicecenterFulfiltaskQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
// OuterId Setter
// 核销单外部单号
func (r *AlibabaServicecenterFulfiltaskQueryRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r AlibabaServicecenterFulfiltaskQueryRequest) GetOuterId() string {
    return r._outerId
}
// GmtCreateEnd Setter
// 时间段查询，核销单创建时间，时间段跨度不超过15分钟
func (r *AlibabaServicecenterFulfiltaskQueryRequest) SetGmtCreateEnd(_gmtCreateEnd string) error {
    r._gmtCreateEnd = _gmtCreateEnd
    r.Set("gmt_create_end", _gmtCreateEnd)
    return nil
}

// GmtCreateEnd Getter
func (r AlibabaServicecenterFulfiltaskQueryRequest) GetGmtCreateEnd() string {
    return r._gmtCreateEnd
}
// CurrentPage Setter
// 查询第几页
func (r *AlibabaServicecenterFulfiltaskQueryRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r AlibabaServicecenterFulfiltaskQueryRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// FulfilTaskIdList Setter
// 核销单id列表
func (r *AlibabaServicecenterFulfiltaskQueryRequest) SetFulfilTaskIdList(_fulfilTaskIdList []int64) error {
    r._fulfilTaskIdList = _fulfilTaskIdList
    r.Set("fulfil_task_id_list", _fulfilTaskIdList)
    return nil
}

// FulfilTaskIdList Getter
func (r AlibabaServicecenterFulfiltaskQueryRequest) GetFulfilTaskIdList() []int64 {
    return r._fulfilTaskIdList
}
