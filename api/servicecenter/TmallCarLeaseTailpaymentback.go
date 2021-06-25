package servicecenter

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/servicecenter"
)

/* 
尾款处置方案回传 
tmall.car.lease.tailpaymentback

尾款处置方案回传
*/
func TmallCarLeaseTailpaymentback(clt *core.SDKClient, req *servicecenter.TmallCarLeaseTailpaymentbackRequest, session string) (*servicecenter.TmallCarLeaseTailpaymentbackResponse, error) {
    var resp servicecenter.TmallCarLeaseTailpaymentbackAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
