package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌查询接口 APIResponse
alibaba.wdk.merchant.brand.query

三江erp对接时，提供品牌查询的接口
*/
type AlibabaWdkMerchantBrandQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMerchantBrandQueryResponse
}

type AlibabaWdkMerchantBrandQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_merchant_brand_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaWdkMerchantBrandQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
