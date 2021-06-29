package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
isv调用gateway API请求
alibaba.interact.isv.gateway

isv能够调用jae本身的server
*/
type AlibabaInteractIsvGatewayRequest struct {
    model.Params
}

// 初始化AlibabaInteractIsvGatewayRequest对象
func NewAlibabaInteractIsvGatewayRequest() *AlibabaInteractIsvGatewayRequest{
    return &AlibabaInteractIsvGatewayRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractIsvGatewayRequest) GetApiMethodName() string {
    return "alibaba.interact.isv.gateway"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractIsvGatewayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
