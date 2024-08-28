package travel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// TaobaoAlitripTravelItemSingleQuery 【API3.0】度假单个商品查询接口
// taobao.alitrip.travel.item.single.query
//
// 旅行度假新商品查询接口（单个商品查询） 第三版
func TaobaoAlitripTravelItemSingleQuery(ctx context.Context, clt *core.SDKClient, req *travel.TaobaoAlitripTravelItemSingleQueryAPIRequest, resp *travel.TaobaoAlitripTravelItemSingleQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
