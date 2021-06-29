package aeusergrowth

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
第三方平台推荐商品 APIResponse
aliexpress.usergrowth.recommend.items.get

第三方平台的推荐AE商品   场景：skin 、底部推荐等
*/
type AliexpressUsergrowthRecommendItemsGetAPIResponse struct {
    model.CommonResponse
    AliexpressUsergrowthRecommendItemsGetResponse
}

type AliexpressUsergrowthRecommendItemsGetResponse struct {
    XMLName xml.Name `xml:"aliexpress_usergrowth_recommend_items_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // response model
    
    Result   *AliexpressUsergrowthRecommendItemsGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
