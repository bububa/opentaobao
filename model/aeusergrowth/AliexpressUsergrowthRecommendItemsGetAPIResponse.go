package aeusergrowth

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
第三方平台推荐商品 API返回值 
aliexpress.usergrowth.recommend.items.get

第三方平台的推荐AE商品   场景：skin 、底部推荐等
*/
type AliexpressUsergrowthRecommendItemsGetAPIResponse struct {
    model.CommonResponse
    AliexpressUsergrowthRecommendItemsGetAPIResponseModel
}

// 第三方平台推荐商品 成功返回结果
type AliexpressUsergrowthRecommendItemsGetAPIResponseModel struct {
    XMLName xml.Name `xml:"aliexpress_usergrowth_recommend_items_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // response model
    Result   *AliexpressUsergrowthRecommendItemsGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
