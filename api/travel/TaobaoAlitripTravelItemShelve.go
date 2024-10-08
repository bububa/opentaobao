package travel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// TaobaoAlitripTravelItemShelve 【API3.0】度假线路商品上下架接口
// taobao.alitrip.travel.item.shelve
//
// 旅行度假新商品发布接口 第三版：度假商品上下架接口
// 注意：定时上下架功能，目前只支持接送、租车类目。
func TaobaoAlitripTravelItemShelve(ctx context.Context, clt *core.SDKClient, req *travel.TaobaoAlitripTravelItemShelveAPIRequest, resp *travel.TaobaoAlitripTravelItemShelveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
