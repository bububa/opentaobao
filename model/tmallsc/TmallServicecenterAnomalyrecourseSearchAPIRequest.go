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
type TmallServicecenterAnomalyrecourseSearchAPIRequest struct {
    model.Params
    // 开始时间
    _start   int64
    // 结束时间
    _end   int64
}

// 初始化TmallServicecenterAnomalyrecourseSearchAPIRequest对象
func NewTmallServicecenterAnomalyrecourseSearchRequest() *TmallServicecenterAnomalyrecourseSearchAPIRequest{
    return &TmallServicecenterAnomalyrecourseSearchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterAnomalyrecourseSearchAPIRequest) GetApiMethodName() string {
    return "tmall.servicecenter.anomalyrecourse.search"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterAnomalyrecourseSearchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Start Setter
// 开始时间
func (r *TmallServicecenterAnomalyrecourseSearchAPIRequest) SetStart(_start int64) error {
    r._start = _start
    r.Set("start", _start)
    return nil
}

// Start Getter
func (r TmallServicecenterAnomalyrecourseSearchAPIRequest) GetStart() int64 {
    return r._start
}
// End Setter
// 结束时间
func (r *TmallServicecenterAnomalyrecourseSearchAPIRequest) SetEnd(_end int64) error {
    r._end = _end
    r.Set("end", _end)
    return nil
}

// End Getter
func (r TmallServicecenterAnomalyrecourseSearchAPIRequest) GetEnd() int64 {
    return r._end
}
