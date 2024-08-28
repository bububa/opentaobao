package travel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// TaobaoAlitripTravelItemNewQuery 【API3.0】新版度假单个商品查询接口
// taobao.alitrip.travel.item.new.query
//
// 新版旅行度假新商品查询接口（单个商品查询）
func TaobaoAlitripTravelItemNewQuery(ctx context.Context, clt *core.SDKClient, req *travel.TaobaoAlitripTravelItemNewQueryAPIRequest, resp *travel.TaobaoAlitripTravelItemNewQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
