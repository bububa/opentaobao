package servicecenter

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/servicecenter"
)

/* 
整车租赁风控模型回调 
tmall.car.lease.riskcallback

租赁公司回调风控结果
*/
func TmallCarLeaseRiskcallback(clt *core.SDKClient, req *servicecenter.TmallCarLeaseRiskcallbackRequest, session string) (*servicecenter.TmallCarLeaseRiskcallbackResponse, error) {
    var resp servicecenter.TmallCarLeaseRiskcallbackAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
