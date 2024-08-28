package tanx

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tanx"
)

// TaobaoTanxCreativeGet 获取单个审核创意状态
// taobao.tanx.creative.get
//
// 获取单个审核创意状态
func TaobaoTanxCreativeGet(ctx context.Context, clt *core.SDKClient, req *tanx.TaobaoTanxCreativeGetAPIRequest, resp *tanx.TaobaoTanxCreativeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
