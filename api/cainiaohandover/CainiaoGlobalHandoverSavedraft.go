package cainiaohandover

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/cainiaohandover"
)

/* 
创建交接单草稿 
cainiao.global.handover.savedraft

提供给ISV通过该接口创建交接单草稿
*/
func CainiaoGlobalHandoverSavedraft(clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalHandoverSavedraftRequest, session string) (*cainiaohandover.CainiaoGlobalHandoverSavedraftAPIResponse, error) {
    var resp cainiaohandover.CainiaoGlobalHandoverSavedraftAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
