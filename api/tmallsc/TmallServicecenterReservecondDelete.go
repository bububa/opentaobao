package tmallsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallsc"
)

/* 
删除主动预约开通条件 
tmall.servicecenter.reservecond.delete

删除主动预约开通条件
*/
func TmallServicecenterReservecondDelete(clt *core.SDKClient, req *tmallsc.TmallServicecenterReservecondDeleteRequest, session string) (*tmallsc.TmallServicecenterReservecondDeleteResponse, error) {
    var resp tmallsc.TmallServicecenterReservecondDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
