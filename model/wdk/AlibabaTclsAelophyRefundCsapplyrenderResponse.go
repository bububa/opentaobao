package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家代客售后逆向申请渲染获取 APIResponse
alibaba.tcls.aelophy.refund.csapplyrender

提供商家代客售后逆向申请渲染获取的接口
*/
type AlibabaTclsAelophyRefundCsapplyrenderAPIResponse struct {
    model.CommonResponse
    AlibabaTclsAelophyRefundCsapplyrenderResponse
}

type AlibabaTclsAelophyRefundCsapplyrenderResponse struct {
    XMLName xml.Name `xml:"alibaba_tcls_aelophy_refund_csapplyrender_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 响应结果
    
    ApiResult   *AlibabaTclsAelophyRefundCsapplyrenderApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`

    
}
