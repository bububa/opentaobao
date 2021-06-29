package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店商品信心查询 APIResponse
alibaba.wdk.merchant.storeitem.query

门店商品信心查询
*/
type AlibabaWdkMerchantStoreitemQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMerchantStoreitemQueryResponse
}

type AlibabaWdkMerchantStoreitemQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_merchant_storeitem_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaWdkMerchantStoreitemQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
