package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
分销商合约生产 
alibaba.wtt.order.contract.subscribe

分销商合约生产
*/
func AlibabaWttOrderContractSubscribe(clt *core.SDKClient, req *alicom.AlibabaWttOrderContractSubscribeAPIRequest, session string) (*alicom.AlibabaWttOrderContractSubscribeAPIResponse, error) {
    var resp alicom.AlibabaWttOrderContractSubscribeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
