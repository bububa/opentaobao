package yunosad

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosad"
)

// YunosAdAuditCreativeGetlist 批量获取创意审核状态
// yunos.ad.audit.creative.getlist
//
// 批量获取创意审核状态
func YunosAdAuditCreativeGetlist(ctx context.Context, clt *core.SDKClient, req *yunosad.YunosAdAuditCreativeGetlistAPIRequest, resp *yunosad.YunosAdAuditCreativeGetlistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
