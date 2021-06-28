package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
根据时间段查询服务商的服务预警消息列表(15分钟内) 
tmall.servicecenter.servicemonitormessage.search

根据时间段查询服务商的服务预警消息列表(15分钟内)
*/
func TmallServicecenterServicemonitormessageSearch(clt *core.SDKClient, req *tmallservice.TmallServicecenterServicemonitormessageSearchRequest, session string) (*tmallservice.TmallServicecenterServicemonitormessageSearchAPIResponse, error) {
    var resp tmallservice.TmallServicecenterServicemonitormessageSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
