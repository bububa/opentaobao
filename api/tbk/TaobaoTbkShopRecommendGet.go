package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkShopRecommendGet 淘宝客-公用-店铺关联推荐
// taobao.tbk.shop.recommend.get
//
// 入参卖家id，可推荐与此店铺相关联的相关店铺。
func TaobaoTbkShopRecommendGet(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkShopRecommendGetAPIRequest, resp *tbk.TaobaoTbkShopRecommendGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
