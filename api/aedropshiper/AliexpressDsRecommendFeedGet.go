package aedropshiper

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// AliexpressDsRecommendFeedGet 获取推荐商品信息流接口
// aliexpress.ds.recommend.feed.get
//
// 获取推荐商品信息流
func AliexpressDsRecommendFeedGet(ctx context.Context, clt *core.SDKClient, req *aedropshiper.AliexpressDsRecommendFeedGetAPIRequest, resp *aedropshiper.AliexpressDsRecommendFeedGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
