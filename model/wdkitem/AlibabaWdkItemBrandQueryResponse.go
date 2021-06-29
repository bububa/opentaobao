package wdkitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌信息查询 APIResponse
alibaba.wdk.item.brand.query

品牌信息查询
*/
type AlibabaWdkItemBrandQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkItemBrandQueryResponse
}

type AlibabaWdkItemBrandQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_item_brand_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaWdkItemBrandQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
