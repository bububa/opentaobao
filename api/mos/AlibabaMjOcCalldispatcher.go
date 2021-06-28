package mos

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mos"
)

/* 
呼叫运力 
alibaba.mj.oc.calldispatcher

定时达呼叫运力接口
*/
func AlibabaMjOcCalldispatcher(clt *core.SDKClient, req *mos.AlibabaMjOcCalldispatcherRequest, session string) (*mos.AlibabaMjOcCalldispatcherAPIResponse, error) {
    var resp mos.AlibabaMjOcCalldispatcherAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
