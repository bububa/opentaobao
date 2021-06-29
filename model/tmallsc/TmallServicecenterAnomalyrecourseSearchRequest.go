package tmallsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫服务平台服务商一键求助单查询 API请求
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

// 初始化TmallServicecenterAnomalyrecourseSearchRequest对象
func NewTmallServicecenterAnomalyrecourseSearchRequest() *TmallServicecenterAnomalyrecourseSearchRequest{
    return &TmallServicecenterAnomalyrecourseSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterAnomalyrecourseSearchRequest) GetApiMethodName() string {
    return "tmall.servicecenter.anomalyrecourse.search"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterAnomalyrecourseSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Start Setter
// 开始时间
func (r *TmallServicecenterAnomalyrecourseSearchRequest) SetStart(start int64) error {
    r.start = start
    r.Set("start", start)
    return nil
}

// Start Getter
func (r TmallServicecenterAnomalyrecourseSearchRequest) GetStart() int64 {
    return r.start
}
// End Setter
// 结束时间
func (r *TmallServicecenterAnomalyrecourseSearchRequest) SetEnd(end int64) error {
    r.end = end
    r.Set("end", end)
    return nil
}

// End Getter
func (r TmallServicecenterAnomalyrecourseSearchRequest) GetEnd() int64 {
    return r.end
}
