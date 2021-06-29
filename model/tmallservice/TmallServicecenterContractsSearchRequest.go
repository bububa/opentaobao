package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取合同类的服务工单信息 API请求
tmall.servicecenter.contracts.search

获取合同类的服务工单信息
*/
type TmallServicecenterContractsSearchRequest struct {
    model.Params
    // 开始时间:  开始时间和结束时间不能超过15分钟
    start   int64
    // 结束时间:  开始时间和结束时间不能超过15分钟
    end   int64
}

// 初始化TmallServicecenterContractsSearchRequest对象
func NewTmallServicecenterContractsSearchRequest() *TmallServicecenterContractsSearchRequest{
    return &TmallServicecenterContractsSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterContractsSearchRequest) GetApiMethodName() string {
    return "tmall.servicecenter.contracts.search"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterContractsSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Start Setter
// 开始时间:  开始时间和结束时间不能超过15分钟
func (r *TmallServicecenterContractsSearchRequest) SetStart(start int64) error {
    r.start = start
    r.Set("start", start)
    return nil
}

// Start Getter
func (r TmallServicecenterContractsSearchRequest) GetStart() int64 {
    return r.start
}
// End Setter
// 结束时间:  开始时间和结束时间不能超过15分钟
func (r *TmallServicecenterContractsSearchRequest) SetEnd(end int64) error {
    r.end = end
    r.Set("end", end)
    return nil
}

// End Getter
func (r TmallServicecenterContractsSearchRequest) GetEnd() int64 {
    return r.end
}
