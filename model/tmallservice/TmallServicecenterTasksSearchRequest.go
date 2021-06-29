package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询任务类工单信息 API请求
tmall.servicecenter.tasks.search

查询任务类工单信息
*/
type TmallServicecenterTasksSearchRequest struct {
    model.Params
    // 开始时间:  开始时间和结束时间不能超过15分钟
    _start   int64
    // 结束时间:  开始时间和结束时间不能超过15分钟
    _end   int64
}

// 初始化TmallServicecenterTasksSearchRequest对象
func NewTmallServicecenterTasksSearchRequest() *TmallServicecenterTasksSearchRequest{
    return &TmallServicecenterTasksSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterTasksSearchRequest) GetApiMethodName() string {
    return "tmall.servicecenter.tasks.search"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterTasksSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Start Setter
// 开始时间:  开始时间和结束时间不能超过15分钟
func (r *TmallServicecenterTasksSearchRequest) SetStart(_start int64) error {
    r._start = _start
    r.Set("start", _start)
    return nil
}

// Start Getter
func (r TmallServicecenterTasksSearchRequest) GetStart() int64 {
    return r._start
}
// End Setter
// 结束时间:  开始时间和结束时间不能超过15分钟
func (r *TmallServicecenterTasksSearchRequest) SetEnd(_end int64) error {
    r._end = _end
    r.Set("end", _end)
    return nil
}

// End Getter
func (r TmallServicecenterTasksSearchRequest) GetEnd() int64 {
    return r._end
}
