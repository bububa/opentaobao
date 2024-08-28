package damai

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenPushitem 大麦换验平台-第三方对外开放-票品接口pushItem
// alibaba.damai.mev.open.pushitem
//
// 开放接口 推送票品
func AlibabaDamaiMevOpenPushitem(ctx context.Context, clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenPushitemAPIRequest, resp *damai.AlibabaDamaiMevOpenPushitemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
