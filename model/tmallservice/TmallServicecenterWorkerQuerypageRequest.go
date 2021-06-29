package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询工人列表 API请求
tmall.servicecenter.worker.querypage

服务商查询工人列表
*/
type TmallServicecenterWorkerQuerypageRequest struct {
    model.Params
    // 页码
    pageIndex   int64
}

// 初始化TmallServicecenterWorkerQuerypageRequest对象
func NewTmallServicecenterWorkerQuerypageRequest() *TmallServicecenterWorkerQuerypageRequest{
    return &TmallServicecenterWorkerQuerypageRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkerQuerypageRequest) GetApiMethodName() string {
    return "tmall.servicecenter.worker.querypage"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkerQuerypageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageIndex Setter
// 页码
func (r *TmallServicecenterWorkerQuerypageRequest) SetPageIndex(pageIndex int64) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

// PageIndex Getter
func (r TmallServicecenterWorkerQuerypageRequest) GetPageIndex() int64 {
    return r.pageIndex
}
