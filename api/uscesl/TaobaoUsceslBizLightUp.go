package uscesl

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// TaobaoUsceslBizLightUp 价签LED等点亮
// taobao.uscesl.biz.light.up
//
// 价签LED等点亮
func TaobaoUsceslBizLightUp(ctx context.Context, clt *core.SDKClient, req *uscesl.TaobaoUsceslBizLightUpAPIRequest, resp *uscesl.TaobaoUsceslBizLightUpAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
