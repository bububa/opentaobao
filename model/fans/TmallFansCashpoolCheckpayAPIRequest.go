package fans

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallFansCashpoolCheckpayAPIRequest
检查资金池付款状态 API请求
tmall.fans.cashpool.checkpay

检查资金池付款状态 */
type TmallFansCashpoolCheckpayAPIRequest struct {
	model.Params
	// 资金池列表
	_cashPoolList []int64
}

// New
