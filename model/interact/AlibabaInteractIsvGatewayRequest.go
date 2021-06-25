package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
isv调用gateway APIRequest
alibaba.interact.isv.gateway

isv能够调用jae本身的server
*/
type AlibabaInteractIsvGatewayRequest struct {
    model.Params

}

func NewAlibabaInteractIsvGatewayRequest() *AlibabaInteractIsvGatewayRequest{
    return &AlibabaInteractIsvGatewayRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractIsvGatewayRequest) GetApiMethodName() string {
    return "alibaba.interact.isv.gateway"
}

func (r AlibabaInteractIsvGatewayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


