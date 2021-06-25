package taotv

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取轮播分类列表 APIResponse
taobao.taotv.carousel.category.list

获取轮播分类列表
*/
type TaobaoTaotvCarouselCategoryListAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTaotvCarouselCategoryListResponse `json:"taobao_taotv_carousel_category_list_response,omitempty"`
}

type TaobaoTaotvCarouselCategoryListResponse struct {

    // result
    Result   *TaobaoTaotvCarouselCategoryListResult `json:"result,omitempty"`

}
