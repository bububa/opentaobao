package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenmallRefundGetAPIRequest
获取OpenMall退款单详情 API请求
taobao.openmall.refund.get

获取OpenMall退款单详情 */
type TaobaoOpenmallRefundGetAPIRequest struct {
	model.Params
	// 渠道商身份
	_distributor string
	// 退款单ID
	_refundId int64
}

// New
