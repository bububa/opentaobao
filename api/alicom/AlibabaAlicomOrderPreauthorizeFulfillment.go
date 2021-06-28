package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
履约结果 
alibaba.alicom.order.preauthorize.fulfillment

预授权-履约结果
*/
func AlibabaAlicomOrderPreauthorizeFulfillment(clt *core.SDKClient, req *alicom.AlibabaAlicomOrderPreauthorizeFulfillmentRequest, session string) (*alicom.AlibabaAlicomOrderPreauthorizeFulfillmentAPIResponse, error) {
    var resp alicom.AlibabaAlicomOrderPreauthorizeFulfillmentAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
