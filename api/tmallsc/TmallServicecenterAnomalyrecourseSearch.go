package tmallsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallsc"
)

/* 
天猫服务平台服务商一键求助单查询 
tmall.servicecenter.anomalyrecourse.search

天猫服务平台服务商一键求助单查询
*/
func TmallServicecenterAnomalyrecourseSearch(clt *core.SDKClient, req *tmallsc.TmallServicecenterAnomalyrecourseSearchRequest, session string) (*tmallsc.TmallServicecenterAnomalyrecourseSearchResponse, error) {
    var resp tmallsc.TmallServicecenterAnomalyrecourseSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
