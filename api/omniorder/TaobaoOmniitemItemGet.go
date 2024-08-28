package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniitemItemGet 获取全渠道门店商品
// taobao.omniitem.item.get
//
// 通过门店id/类目id/商品id单个或多个参数组合查询全渠道门店商品信息
func TaobaoOmniitemItemGet(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoOmniitemItemGetAPIRequest, resp *omniorder.TaobaoOmniitemItemGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
