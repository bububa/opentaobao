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
    _start   int64
    // 结束时间:  开始时间和结束时间不能超过15分钟
    _end   int64
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
func (r *TmallServicecenterContractsSearchRequest) SetStart(_start int64) error {
    r._start = _start
    r.Set("start", _start)
    return nil
}

// Start Getter
func (r TmallServicecenterContractsSearchRequest) GetStart() int64 {
    return r._start
}
// End Setter
// 结束时间:  开始时间和结束时间不能超过15分钟
func (r *TmallServicecenterContractsSearchRequest) SetEnd(_end int64) error {
    r._end = _end
    r.Set("end", _end)
    return nil
}

// End Getter
func (r TmallServicecenterContractsSearchRequest) GetEnd() int64 {
    return r._end
}
