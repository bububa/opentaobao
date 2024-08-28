package baodian

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baodian"
)

// TaobaoBaodianDepositGet 宝点用户帐户查询（已迁移）
// taobao.baodian.deposit.get
//
// 查询用户宝点帐户信息及当前宝点价格
func TaobaoBaodianDepositGet(ctx context.Context, clt *core.SDKClient, req *baodian.TaobaoBaodianDepositGetAPIRequest, resp *baodian.TaobaoBaodianDepositGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
