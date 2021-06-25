package yunos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
alios cosmo服务调用 APIResponse
aliyun.cosmo.gateway.invoke

AliOS cosmo服务分发平台对外调用接口
*/
type AliyunCosmoGatewayInvokeAPIResponse struct {
    model.CommonResponse
    Response *AliyunCosmoGatewayInvokeResponse `json:"aliyun_cosmo_gateway_invoke_response,omitempty"`
}

type AliyunCosmoGatewayInvokeResponse struct {

    // result
    Result   *RdamResponse `json:"result,omitempty"`

}
