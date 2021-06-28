package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
根据一键求助id查询指定服务商的一键求助单 
tmall.servicecenter.anomalyrecourse.querybyid

根据一键求助id查询指定服务商的一键求助单
*/
func TmallServicecenterAnomalyrecourseQuerybyid(clt *core.SDKClient, req *tmallservice.TmallServicecenterAnomalyrecourseQuerybyidRequest, session string) (*tmallservice.TmallServicecenterAnomalyrecourseQuerybyidAPIResponse, error) {
    var resp tmallservice.TmallServicecenterAnomalyrecourseQuerybyidAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
