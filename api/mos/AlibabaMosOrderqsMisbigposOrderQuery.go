package mos

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mos"
)

/* 
大pos新选单退 
alibaba.mos.orderqs.misbigpos.order.query

大pos新选单退
*/
func AlibabaMosOrderqsMisbigposOrderQuery(clt *core.SDKClient, req *mos.AlibabaMosOrderqsMisbigposOrderQueryAPIRequest, session string) (*mos.AlibabaMosOrderqsMisbigposOrderQueryAPIResponse, error) {
    var resp mos.AlibabaMosOrderqsMisbigposOrderQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
