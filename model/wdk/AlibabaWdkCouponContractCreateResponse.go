package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
营销券合同创建接口 APIResponse
alibaba.wdk.coupon.contract.create

营销券合同创建接口
*/
type AlibabaWdkCouponContractCreateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkCouponContractCreateResponse
}

type AlibabaWdkCouponContractCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_coupon_contract_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *AlibabaWdkCouponContractCreateApiResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
