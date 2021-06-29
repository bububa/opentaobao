package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
营销券合同创建接口 API返回值 
alibaba.wdk.coupon.contract.create

营销券合同创建接口
*/
type AlibabaWdkCouponContractCreateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkCouponContractCreateResponse
}

// 营销券合同创建接口 成功返回结果
type AlibabaWdkCouponContractCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_coupon_contract_create_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *AlibabaWdkCouponContractCreateApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
