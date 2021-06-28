package taotv

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取轮播分类列表 APIResponse
taobao.taotv.carousel.category.list

获取轮播分类列表
*/
type TaobaoTaotvCarouselCategoryListAPIResponse struct {
    model.CommonResponse
    TaobaoTaotvCarouselCategoryListResponse
}

type TaobaoTaotvCarouselCategoryListResponse struct {
    XMLName xml.Name `xml:"taotv_carousel_category_list_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TaobaoTaotvCarouselCategoryListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
