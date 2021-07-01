package taotv

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取轮播分类列表 API返回值 
taobao.taotv.carousel.category.list

获取轮播分类列表
*/
type TaobaoTaotvCarouselCategoryListAPIResponse struct {
    model.CommonResponse
    TaobaoTaotvCarouselCategoryListAPIResponseModel
}

// 获取轮播分类列表 成功返回结果
type TaobaoTaotvCarouselCategoryListAPIResponseModel struct {
    XMLName xml.Name `xml:"taotv_carousel_category_list_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TaobaoTaotvCarouselCategoryListResult `json:"result,omitempty" xml:"result,omitempty"`
}
