package yunos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
alios cosmo服务调用 API返回值 
aliyun.cosmo.gateway.invoke

AliOS cosmo服务分发平台对外调用接口
*/
type AliyunCosmoGatewayInvokeAPIResponse struct {
    model.CommonResponse
    AliyunCosmoGatewayInvokeResponse
}

// alios cosmo服务调用 成功返回结果
type AliyunCosmoGatewayInvokeResponse struct {
    XMLName xml.Name `xml:"aliyun_cosmo_gateway_invoke_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *RdamResponse `json:"result,omitempty" xml:"result,omitempty"`
}
