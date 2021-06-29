package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询需求容量 API请求
tmall.servicecenter.worker.querycapacitytask

查询需求容量
*/
type TmallServicecenterWorkerQuerycapacitytaskRequest struct {
    model.Params
    // 查询对象
    _query   *CapacityTaskQueryDTO
}

// 初始化TmallServicecenterWorkerQuerycapacitytaskRequest对象
func NewTmallServicecenterWorkerQuerycapacitytaskRequest() *TmallServicecenterWorkerQuerycapacitytaskRequest{
    return &TmallServicecenterWorkerQuerycapacitytaskRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkerQuerycapacitytaskRequest) GetApiMethodName() string {
    return "tmall.servicecenter.worker.querycapacitytask"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkerQuerycapacitytaskRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 查询对象
func (r *TmallServicecenterWorkerQuerycapacitytaskRequest) SetQuery(_query *CapacityTaskQueryDTO) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r TmallServicecenterWorkerQuerycapacitytaskRequest) GetQuery() *CapacityTaskQueryDTO {
    return r._query
}
