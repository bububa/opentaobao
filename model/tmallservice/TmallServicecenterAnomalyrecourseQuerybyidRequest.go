package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据一键求助id查询指定服务商的一键求助单 APIRequest
tmall.servicecenter.anomalyrecourse.querybyid

根据一键求助id查询指定服务商的一键求助单
*/
type TmallServicecenterAnomalyrecourseQuerybyidRequest struct {
    model.Params

    // 一键求助的id
    anomalyRecourseId   int64 

}

func NewTmallServicecenterAnomalyrecourseQuerybyidRequest() *TmallServicecenterAnomalyrecourseQuerybyidRequest{
    return &TmallServicecenterAnomalyrecourseQuerybyidRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterAnomalyrecourseQuerybyidRequest) GetApiMethodName() string {
    return "tmall.servicecenter.anomalyrecourse.querybyid"
}

func (r TmallServicecenterAnomalyrecourseQuerybyidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterAnomalyrecourseQuerybyidRequest) SetAnomalyRecourseId(anomalyRecourseId int64) error {
    r.anomalyRecourseId = anomalyRecourseId
    r.Set("anomaly_recourse_id", anomalyRecourseId)
    return nil
}

func (r TmallServicecenterAnomalyrecourseQuerybyidRequest) GetAnomalyRecourseId() int64 {
    return r.anomalyRecourseId
}

