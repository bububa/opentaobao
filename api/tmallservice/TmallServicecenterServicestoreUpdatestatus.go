package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
网点/门店状态修改 
tmall.servicecenter.servicestore.updatestatus

修改网点/门店状态
*/
func TmallServicecenterServicestoreUpdatestatus(clt *core.SDKClient, req *tmallservice.TmallServicecenterServicestoreUpdatestatusRequest, session string) (*tmallservice.TmallServicecenterServicestoreUpdatestatusAPIResponse, error) {
    var resp tmallservice.TmallServicecenterServicestoreUpdatestatusAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
