package uscesl

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// TaobaoUsceslBizItemLightUp 商品条码亮灯API
// taobao.uscesl.biz.item.light.up
//
// 亮灯API
func TaobaoUsceslBizItemLightUp(ctx context.Context, clt *core.SDKClient, req *uscesl.TaobaoUsceslBizItemLightUpAPIRequest, resp *uscesl.TaobaoUsceslBizItemLightUpAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
