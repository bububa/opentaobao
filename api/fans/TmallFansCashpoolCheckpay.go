package fans

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fans"
)

// TmallFansCashpoolCheckpay 检查资金池付款状态
// tmall.fans.cashpool.checkpay
//
// 检查资金池付款状态
func TmallFansCashpoolCheckpay(clt *core.SDKClient, req *fans.TmallFansCashpoolCheckpayAPIRequest, resp *fans.TmallFansCashpoolCheckpayAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
