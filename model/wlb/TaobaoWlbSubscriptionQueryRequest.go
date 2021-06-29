package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商家定购的所有服务 API请求
taobao.wlb.subscription.query

查询商家定购的所有服务,可通过入参状态来筛选
*/
type TaobaoWlbSubscriptionQueryRequest struct {
    model.Params
    // 状态 <br/>AUDITING 1-待审核; <br/>CANCEL 2-撤销 ;<br/>CHECKED 3-审核通过 ;<br/>FAILED 4-审核未通过 ;<br/>SYNCHRONIZING 5-同步中;<br/>只允许输入上面指定的值，且可以为空，为空时查询所有状态。若输错了，则按AUDITING处理。
    status   string
    // 当前页
    pageNo   int64
    // 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
    pageSize   int64
}

// 初始化TaobaoWlbSubscriptionQueryRequest对象
func NewTaobaoWlbSubscriptionQueryRequest() *TaobaoWlbSubscriptionQueryRequest{
    return &TaobaoWlbSubscriptionQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbSubscriptionQueryRequest) GetApiMethodName() string {
    return "taobao.wlb.subscription.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbSubscriptionQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Status Setter
// 状态 <br/>AUDITING 1-待审核; <br/>CANCEL 2-撤销 ;<br/>CHECKED 3-审核通过 ;<br/>FAILED 4-审核未通过 ;<br/>SYNCHRONIZING 5-同步中;<br/>只允许输入上面指定的值，且可以为空，为空时查询所有状态。若输错了，则按AUDITING处理。
func (r *TaobaoWlbSubscriptionQueryRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoWlbSubscriptionQueryRequest) GetStatus() string {
    return r.status
}
// PageNo Setter
// 当前页
func (r *TaobaoWlbSubscriptionQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoWlbSubscriptionQueryRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
func (r *TaobaoWlbSubscriptionQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoWlbSubscriptionQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
