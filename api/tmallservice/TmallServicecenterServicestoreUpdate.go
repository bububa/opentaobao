package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
修改门店信息 
tmall.servicecenter.servicestore.update

用于修改门店/网点信息。多个业务共用
*/
func TmallServicecenterServicestoreUpdate(clt *core.SDKClient, req *tmallservice.TmallServicecenterServicestoreUpdateRequest, session string) (*tmallservice.TmallServicecenterServicestoreUpdateResponse, error) {
    var resp tmallservice.TmallServicecenterServicestoreUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
