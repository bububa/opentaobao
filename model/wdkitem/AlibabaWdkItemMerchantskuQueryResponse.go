package wdkitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家商品信息查询 APIResponse
alibaba.wdk.item.merchantsku.query

商家商品信息查询
*/
type AlibabaWdkItemMerchantskuQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkItemMerchantskuQueryResponse
}

type AlibabaWdkItemMerchantskuQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_item_merchantsku_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaWdkItemMerchantskuQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
