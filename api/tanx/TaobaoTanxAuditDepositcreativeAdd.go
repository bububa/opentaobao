package tanx

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tanx"
)

// TaobaoTanxAuditDepositcreativeAdd dsp托管创意新增接口
// taobao.tanx.audit.depositcreative.add
//
// dsp托管创意新增接口
func TaobaoTanxAuditDepositcreativeAdd(ctx context.Context, clt *core.SDKClient, req *tanx.TaobaoTanxAuditDepositcreativeAddAPIRequest, resp *tanx.TaobaoTanxAuditDepositcreativeAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
