package alihealthoutflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
支付宝短链跳转三方h5通用接口 API返回值 
alibaba.alihealth.short.url.get

支付宝短链跳转三方h5通用接口
*/
type AlibabaAlihealthShortUrlGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthShortUrlGetResponse
}

// 支付宝短链跳转三方h5通用接口 成功返回结果
type AlibabaAlihealthShortUrlGetResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_short_url_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // alinkappserver系统返回的通用结果类
    ServiceResult   *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}
