package cainiaohandover

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/cainiaohandover"
)

/* 
取消交接单 
cainiao.global.handover.cancel

提供给ISV通过该接口取消交接单
*/
func CainiaoGlobalHandoverCancel(clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalHandoverCancelAPIRequest, session string) (*cainiaohandover.CainiaoGlobalHandoverCancelAPIResponse, error) {
    var resp cainiaohandover.CainiaoGlobalHandoverCancelAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
