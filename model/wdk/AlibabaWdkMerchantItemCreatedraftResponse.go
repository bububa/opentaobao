package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新建商品草稿 APIResponse
alibaba.wdk.merchant.item.createdraft

新建商品草稿erp接口
*/
type AlibabaWdkMerchantItemCreatedraftAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMerchantItemCreatedraftResponse
}

type AlibabaWdkMerchantItemCreatedraftResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_merchant_item_createdraft_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaWdkMerchantItemCreatedraftResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
