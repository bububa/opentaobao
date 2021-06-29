package wdkitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店商品信息查询 APIResponse
alibaba.wdk.item.storesku.query

门店商品信息查询
*/
type AlibabaWdkItemStoreskuQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkItemStoreskuQueryResponse
}

type AlibabaWdkItemStoreskuQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_item_storesku_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaWdkItemStoreskuQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
