package alihealthoutflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
支付宝短链跳转三方h5通用接口 APIResponse
alibaba.alihealth.short.url.get

支付宝短链跳转三方h5通用接口
*/
type AlibabaAlihealthShortUrlGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthShortUrlGetResponse
}

type AlibabaAlihealthShortUrlGetResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_short_url_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    ServiceResult   *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`

    
}
