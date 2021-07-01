package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据一键求助id查询指定服务商的一键求助单 API请求
tmall.servicecenter.anomalyrecourse.querybyid

根据一键求助id查询指定服务商的一键求助单
*/
type TmallServicecenterAnomalyrecourseQuerybyidAPIRequest struct {
    model.Params
    // 一键求助的id
    _anomalyRecourseId   int64
}

// 初始化TmallServicecenterAnomalyrecourseQuerybyidAPIRequest对象
func NewTmallServicecenterAnomalyrecourseQuerybyidRequest() *TmallServicecenterAnomalyrecourseQuerybyidAPIRequest{
    return &TmallServicecenterAnomalyrecourseQuerybyidAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterAnomalyrecourseQuerybyidAPIRequest) GetApiMethodName() string {
    return "tmall.servicecenter.anomalyrecourse.querybyid"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterAnomalyrecourseQuerybyidAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AnomalyRecourseId Setter
// 一键求助的id
func (r *TmallServicecenterAnomalyrecourseQuerybyidAPIRequest) SetAnomalyRecourseId(_anomalyRecourseId int64) error {
    r._anomalyRecourseId = _anomalyRecourseId
    r.Set("anomaly_recourse_id", _anomalyRecourseId)
    return nil
}

// AnomalyRecourseId Getter
func (r TmallServicecenterAnomalyrecourseQuerybyidAPIRequest) GetAnomalyRecourseId() int64 {
    return r._anomalyRecourseId
}
