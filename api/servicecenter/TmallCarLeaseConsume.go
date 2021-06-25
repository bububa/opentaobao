package servicecenter

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/servicecenter"
)

/* 
汽车租赁核销 
tmall.car.lease.consume

租赁公司回传信息，核销
*/
func TmallCarLeaseConsume(clt *core.SDKClient, req *servicecenter.TmallCarLeaseConsumeRequest, session string) (*servicecenter.TmallCarLeaseConsumeResponse, error) {
    var resp servicecenter.TmallCarLeaseConsumeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
