package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家商品查询 APIResponse
alibaba.wdk.merchant.item.query

商家商品查询
*/
type AlibabaWdkMerchantItemQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMerchantItemQueryResponse
}

type AlibabaWdkMerchantItemQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_merchant_item_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaWdkMerchantItemQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
