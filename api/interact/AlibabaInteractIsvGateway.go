package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
isv调用gateway 
alibaba.interact.isv.gateway

isv能够调用jae本身的server
*/
func AlibabaInteractIsvGateway(clt *core.SDKClient, req *interact.AlibabaInteractIsvGatewayAPIRequest, session string) (*interact.AlibabaInteractIsvGatewayAPIResponse, error) {
    var resp interact.AlibabaInteractIsvGatewayAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
