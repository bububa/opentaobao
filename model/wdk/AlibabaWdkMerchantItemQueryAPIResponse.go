package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家商品查询 API返回值 
alibaba.wdk.merchant.item.query

商家商品查询
*/
type AlibabaWdkMerchantItemQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMerchantItemQueryAPIResponseModel
}

// 商家商品查询 成功返回结果
type AlibabaWdkMerchantItemQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_wdk_merchant_item_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AlibabaWdkMerchantItemQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
