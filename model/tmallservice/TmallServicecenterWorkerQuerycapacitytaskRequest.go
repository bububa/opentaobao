package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询需求容量 APIRequest
tmall.servicecenter.worker.querycapacitytask

查询需求容量
*/
type TmallServicecenterWorkerQuerycapacitytaskRequest struct {
    model.Params

    // 查询对象
    query   *CapacityTaskQueryDto 

}

func NewTmallServicecenterWorkerQuerycapacitytaskRequest() *TmallServicecenterWorkerQuerycapacitytaskRequest{
    return &TmallServicecenterWorkerQuerycapacitytaskRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterWorkerQuerycapacitytaskRequest) GetApiMethodName() string {
    return "tmall.servicecenter.worker.querycapacitytask"
}

func (r TmallServicecenterWorkerQuerycapacitytaskRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterWorkerQuerycapacitytaskRequest) SetQuery(query *CapacityTaskQueryDto) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r TmallServicecenterWorkerQuerycapacitytaskRequest) GetQuery() *CapacityTaskQueryDto {
    return r.query
}

