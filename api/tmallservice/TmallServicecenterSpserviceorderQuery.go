package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
服务单列表查询 
tmall.servicecenter.spserviceorder.query

查询服务单列表
*/
func TmallServicecenterSpserviceorderQuery(clt *core.SDKClient, req *tmallservice.TmallServicecenterSpserviceorderQueryAPIRequest, session string) (*tmallservice.TmallServicecenterSpserviceorderQueryAPIResponse, error) {
    var resp tmallservice.TmallServicecenterSpserviceorderQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
