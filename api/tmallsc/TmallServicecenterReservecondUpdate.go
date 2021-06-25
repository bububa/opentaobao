package tmallsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallsc"
)

/* 
主动预约条件更新 
tmall.servicecenter.reservecond.update

1、设置主动预约开通条件
*/
func TmallServicecenterReservecondUpdate(clt *core.SDKClient, req *tmallsc.TmallServicecenterReservecondUpdateRequest, session string) (*tmallsc.TmallServicecenterReservecondUpdateResponse, error) {
    var resp tmallsc.TmallServicecenterReservecondUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
