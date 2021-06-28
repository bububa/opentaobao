package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
分销商提交受理订单 
alibaba.tianji.distributor.order.submit

分销商提交受理订单，如合约订购、充值受理等
*/
func AlibabaTianjiDistributorOrderSubmit(clt *core.SDKClient, req *alicom.AlibabaTianjiDistributorOrderSubmitRequest, session string) (*alicom.AlibabaTianjiDistributorOrderSubmitAPIResponse, error) {
    var resp alicom.AlibabaTianjiDistributorOrderSubmitAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
