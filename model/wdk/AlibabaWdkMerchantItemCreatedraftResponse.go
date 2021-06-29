package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新建商品草稿 API返回值 
alibaba.wdk.merchant.item.createdraft

新建商品草稿erp接口
*/
type AlibabaWdkMerchantItemCreatedraftAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMerchantItemCreatedraftResponse
}

// 新建商品草稿 成功返回结果
type AlibabaWdkMerchantItemCreatedraftResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_merchant_item_createdraft_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AlibabaWdkMerchantItemCreatedraftResult `json:"result,omitempty" xml:"result,omitempty"`
}
