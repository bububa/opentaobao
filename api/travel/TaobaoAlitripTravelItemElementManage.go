package travel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// TaobaoAlitripTravelItemElementManage 【API3.0】资源元素管理接口
// taobao.alitrip.travel.item.element.manage
//
// 资源元素管理接口：提供商家管理（增删改）基本资源元素信息。基本资源元素可供多个商品共享
func TaobaoAlitripTravelItemElementManage(ctx context.Context, clt *core.SDKClient, req *travel.TaobaoAlitripTravelItemElementManageAPIRequest, resp *travel.TaobaoAlitripTravelItemElementManageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
