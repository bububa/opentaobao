package mos

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mos"
)

/* 
根据加密手机号领券 
alibaba.mj.moscarnival.receiveencrypt

根据加密手机号领券
*/
func AlibabaMjMoscarnivalReceiveencrypt(clt *core.SDKClient, req *mos.AlibabaMjMoscarnivalReceiveencryptAPIRequest, session string) (*mos.AlibabaMjMoscarnivalReceiveencryptAPIResponse, error) {
    var resp mos.AlibabaMjMoscarnivalReceiveencryptAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
