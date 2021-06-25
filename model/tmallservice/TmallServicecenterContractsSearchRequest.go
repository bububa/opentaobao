package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取合同类的服务工单信息 APIRequest
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

func NewTmallServicecenterContractsSearchRequest() *TmallServicecenterContractsSearchRequest{
    return &TmallServicecenterContractsSearchRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterContractsSearchRequest) GetApiMethodName() string {
    return "tmall.servicecenter.contracts.search"
}

func (r TmallServicecenterContractsSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterContractsSearchRequest) SetStart(start int64) error {
    r.start = start
    r.Set("start", start)
    return nil
}

func (r TmallServicecenterContractsSearchRequest) GetStart() int64 {
    return r.start
}

func (r *TmallServicecenterContractsSearchRequest) SetEnd(end int64) error {
    r.end = end
    r.Set("end", end)
    return nil
}

func (r TmallServicecenterContractsSearchRequest) GetEnd() int64 {
    return r.end
}

