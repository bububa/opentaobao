package tmallsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallsc"
)

/* 
买家拒收 
tmall.servicecenter.workcard.refuse

买家拒收通知接口
*/
func TmallServicecenterWorkcardRefuse(clt *core.SDKClient, req *tmallsc.TmallServicecenterWorkcardRefuseRequest, session string) (*tmallsc.TmallServicecenterWorkcardRefuseResponse, error) {
    var resp tmallsc.TmallServicecenterWorkcardRefuseAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
