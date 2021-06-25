package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询任务类工单信息 APIRequest
tmall.servicecenter.tasks.search

查询任务类工单信息
*/
type TmallServicecenterTasksSearchRequest struct {
    model.Params

    // 开始时间:  开始时间和结束时间不能超过15分钟
    start   int64 

    // 结束时间:  开始时间和结束时间不能超过15分钟
    end   int64 

}

func NewTmallServicecenterTasksSearchRequest() *TmallServicecenterTasksSearchRequest{
    return &TmallServicecenterTasksSearchRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterTasksSearchRequest) GetApiMethodName() string {
    return "tmall.servicecenter.tasks.search"
}

func (r TmallServicecenterTasksSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterTasksSearchRequest) SetStart(start int64) error {
    r.start = start
    r.Set("start", start)
    return nil
}

func (r TmallServicecenterTasksSearchRequest) GetStart() int64 {
    return r.start
}

func (r *TmallServicecenterTasksSearchRequest) SetEnd(end int64) error {
    r.end = end
    r.Set("end", end)
    return nil
}

func (r TmallServicecenterTasksSearchRequest) GetEnd() int64 {
    return r.end
}

