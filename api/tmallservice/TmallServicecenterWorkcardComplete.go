package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
工单完结 
tmall.servicecenter.workcard.complete

工单完结
*/
func TmallServicecenterWorkcardComplete(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardCompleteRequest, session string) (*tmallservice.TmallServicecenterWorkcardCompleteResponse, error) {
    var resp tmallservice.TmallServicecenterWorkcardCompleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
