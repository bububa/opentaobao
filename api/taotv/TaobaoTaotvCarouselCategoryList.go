package taotv

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/taotv"
)

/* 
获取轮播分类列表 
taobao.taotv.carousel.category.list

获取轮播分类列表
*/
func TaobaoTaotvCarouselCategoryList(clt *core.SDKClient, req *taotv.TaobaoTaotvCarouselCategoryListAPIRequest, session string) (*taotv.TaobaoTaotvCarouselCategoryListAPIResponse, error) {
    var resp taotv.TaobaoTaotvCarouselCategoryListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
