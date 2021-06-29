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
    gmtCreateStart   string
    // 每页条数，默认50，最大50
    pageSize   int64
    // 核销单外部单号
    outerId   string
    // 时间段查询，核销单创建时间，时间段跨度不超过15分钟
    gmtCreateEnd   string
    // 查询第几页
    currentPage   int64
    // 核销单id列表
    fulfilTaskIdList   []int64
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
func (r *AlibabaServicecenterFulfiltaskQueryRequest) SetGmtCreateStart(gmtCreateStart string) error {
    r.gmtCreateStart = gmtCreateStart
    r.Set("gmt_create_start", gmtCreateStart)
    return nil
}

// GmtCreateStart Getter
func (r AlibabaServicecenterFulfiltaskQueryRequest) GetGmtCreateStart() string {
    return r.gmtCreateStart
}
// PageSize Setter
// 每页条数，默认50，最大50
func (r *AlibabaServicecenterFulfiltaskQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaServicecenterFulfiltaskQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
// OuterId Setter
// 核销单外部单号
func (r *AlibabaServicecenterFulfiltaskQueryRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r AlibabaServicecenterFulfiltaskQueryRequest) GetOuterId() string {
    return r.outerId
}
// GmtCreateEnd Setter
// 时间段查询，核销单创建时间，时间段跨度不超过15分钟
func (r *AlibabaServicecenterFulfiltaskQueryRequest) SetGmtCreateEnd(gmtCreateEnd string) error {
    r.gmtCreateEnd = gmtCreateEnd
    r.Set("gmt_create_end", gmtCreateEnd)
    return nil
}

// GmtCreateEnd Getter
func (r AlibabaServicecenterFulfiltaskQueryRequest) GetGmtCreateEnd() string {
    return r.gmtCreateEnd
}
// CurrentPage Setter
// 查询第几页
func (r *AlibabaServicecenterFulfiltaskQueryRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

// CurrentPage Getter
func (r AlibabaServicecenterFulfiltaskQueryRequest) GetCurrentPage() int64 {
    return r.currentPage
}
// FulfilTaskIdList Setter
// 核销单id列表
func (r *AlibabaServicecenterFulfiltaskQueryRequest) SetFulfilTaskIdList(fulfilTaskIdList []int64) error {
    r.fulfilTaskIdList = fulfilTaskIdList
    r.Set("fulfil_task_id_list", fulfilTaskIdList)
    return nil
}

// FulfilTaskIdList Getter
func (r AlibabaServicecenterFulfiltaskQueryRequest) GetFulfilTaskIdList() []int64 {
    return r.fulfilTaskIdList
}
