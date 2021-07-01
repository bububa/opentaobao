package exchange

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallExchangeRefusereasonGetAPIRequest
获取拒绝换货原因列表 API请求
tmall.exchange.refusereason.get

获取拒绝换货原因列表 */
type TmallExchangeRefusereasonGetAPIRequest struct {
	model.Params
	// 换货单号ID
	_disputeId int64
	// 返回字段
	_fields []string
	// 换货申请类型：0-任意类型；1-售中；2-售后
	_disputeType int64
}

// New
