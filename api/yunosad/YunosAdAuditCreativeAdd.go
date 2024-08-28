package yunosad

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosad"
)

// YunosAdAuditCreativeAdd 单个创意预审接口
// yunos.ad.audit.creative.add
//
// YunOS广告业务第三方DSP单个创意预审接口
func YunosAdAuditCreativeAdd(ctx context.Context, clt *core.SDKClient, req *yunosad.YunosAdAuditCreativeAddAPIRequest, resp *yunosad.YunosAdAuditCreativeAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
