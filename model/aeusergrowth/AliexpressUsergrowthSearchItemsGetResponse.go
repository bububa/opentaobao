package aeusergrowth

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
第三方平台搜索AE商品 APIResponse
aliexpress.usergrowth.search.items.get

第三方平台的搜索服务   获取AE商品list
*/
type AliexpressUsergrowthSearchItemsGetAPIResponse struct {
    model.CommonResponse
    AliexpressUsergrowthSearchItemsGetResponse
}

type AliexpressUsergrowthSearchItemsGetResponse struct {
    XMLName xml.Name `xml:"aliexpress_usergrowth_search_items_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // response model
    
    Result   *AliexpressUsergrowthSearchItemsGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
