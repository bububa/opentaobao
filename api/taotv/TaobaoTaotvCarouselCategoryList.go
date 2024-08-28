package taotv

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/taotv"
)

// TaobaoTaotvCarouselCategoryList 获取轮播分类列表
// taobao.taotv.carousel.category.list
//
// 获取轮播分类列表
func TaobaoTaotvCarouselCategoryList(ctx context.Context, clt *core.SDKClient, req *taotv.TaobaoTaotvCarouselCategoryListAPIRequest, resp *taotv.TaobaoTaotvCarouselCategoryListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
