package mos

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mos"
)

/* 
根据手机号码领券 
alibaba.mj.moscarnival.receivecoupon

根据手机号码领券
*/
func AlibabaMjMoscarnivalReceivecoupon(clt *core.SDKClient, req *mos.AlibabaMjMoscarnivalReceivecouponRequest, session string) (*mos.AlibabaMjMoscarnivalReceivecouponResponse, error) {
    var resp mos.AlibabaMjMoscarnivalReceivecouponAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
