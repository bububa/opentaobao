package fans

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fans"
)

// TmallFansCashpoolCheckpay 检查资金池付款状态
// tmall.fans.cashpool.checkpay
//
// 检查资金池付款状态
func TmallFansCashpoolCheckpay(ctx context.Context, clt *core.SDKClient, req *fans.TmallFansCashpoolCheckpayAPIRequest, resp *fans.TmallFansCashpoolCheckpayAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
