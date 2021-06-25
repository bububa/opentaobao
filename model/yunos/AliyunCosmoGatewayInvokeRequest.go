package yunos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
alios cosmo服务调用 APIRequest
aliyun.cosmo.gateway.invoke

AliOS cosmo服务分发平台对外调用接口
*/
type AliyunCosmoGatewayInvokeRequest struct {
    model.Params

    // 请求上下文参数
    context   *RdamContext 

    // 请求对象
    rdamRequest   *RdamGenericRequest 

}

func NewAliyunCosmoGatewayInvokeRequest() *AliyunCosmoGatewayInvokeRequest{
    return &AliyunCosmoGatewayInvokeRequest{
        Params: model.NewParams(),
    }
}

func (r AliyunCosmoGatewayInvokeRequest) GetApiMethodName() string {
    return "aliyun.cosmo.gateway.invoke"
}

func (r AliyunCosmoGatewayInvokeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliyunCosmoGatewayInvokeRequest) SetContext(context *RdamContext) error {
    r.context = context
    r.Set("context", context)
    return nil
}

func (r AliyunCosmoGatewayInvokeRequest) GetContext() *RdamContext {
    return r.context
}

func (r *AliyunCosmoGatewayInvokeRequest) SetRdamRequest(rdamRequest *RdamGenericRequest) error {
    r.rdamRequest = rdamRequest
    r.Set("rdam_request", rdamRequest)
    return nil
}

func (r AliyunCosmoGatewayInvokeRequest) GetRdamRequest() *RdamGenericRequest {
    return r.rdamRequest
}

