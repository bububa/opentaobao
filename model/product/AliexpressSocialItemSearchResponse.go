package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
AE社交选品 APIResponse
aliexpress.social.item.search

AE社交选品,通过各种筛选条件对社交商品池进行筛选
*/
type AliexpressSocialItemSearchAPIResponse struct {
    model.CommonResponse
    AliexpressSocialItemSearchResponse
}

type AliexpressSocialItemSearchResponse struct {
    XMLName xml.Name `xml:"aliexpress_social_item_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 报类型
    
    Result   *ItemPickPagingResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
