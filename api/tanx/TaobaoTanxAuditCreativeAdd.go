package tanx

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tanx"
)

// TaobaoTanxAuditCreativeAdd 创意预审新增接口
// taobao.tanx.audit.creative.add
//
// 创意预审新增接口
func TaobaoTanxAuditCreativeAdd(ctx context.Context, clt *core.SDKClient, req *tanx.TaobaoTanxAuditCreativeAddAPIRequest, resp *tanx.TaobaoTanxAuditCreativeAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
