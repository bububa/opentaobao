package wdkitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
类目查询接口 APIResponse
alibaba.wdk.item.category.query

类目查询接口
*/
type AlibabaWdkItemCategoryQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkItemCategoryQueryResponse
}

type AlibabaWdkItemCategoryQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_item_category_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaWdkItemCategoryQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
