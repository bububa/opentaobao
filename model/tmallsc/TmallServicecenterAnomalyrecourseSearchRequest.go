package tmallsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫服务平台服务商一键求助单查询 APIRequest
tmall.servicecenter.anomalyrecourse.search

天猫服务平台服务商一键求助单查询
*/
type TmallServicecenterAnomalyrecourseSearchRequest struct {
    model.Params

    // 开始时间
    start   int64 

    // 结束时间
    end   int64 

}

func NewTmallServicecenterAnomalyrecourseSearchRequest() *TmallServicecenterAnomalyrecourseSearchRequest{
    return &TmallServicecenterAnomalyrecourseSearchRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterAnomalyrecourseSearchRequest) GetApiMethodName() string {
    return "tmall.servicecenter.anomalyrecourse.search"
}

func (r TmallServicecenterAnomalyrecourseSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterAnomalyrecourseSearchRequest) SetStart(start int64) error {
    r.start = start
    r.Set("start", start)
    return nil
}

func (r TmallServicecenterAnomalyrecourseSearchRequest) GetStart() int64 {
    return r.start
}

func (r *TmallServicecenterAnomalyrecourseSearchRequest) SetEnd(end int64) error {
    r.end = end
    r.Set("end", end)
    return nil
}

func (r TmallServicecenterAnomalyrecourseSearchRequest) GetEnd() int64 {
    return r.end
}

