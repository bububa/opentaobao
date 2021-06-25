package mos

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mos"
)

/* 
创建一个付款单 
alibaba.mj.mos.fund.createbill

创建一个付款单
*/
func AlibabaMjMosFundCreatebill(clt *core.SDKClient, req *mos.AlibabaMjMosFundCreatebillRequest, session string) (*mos.AlibabaMjMosFundCreatebillResponse, error) {
    var resp mos.AlibabaMjMosFundCreatebillAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
