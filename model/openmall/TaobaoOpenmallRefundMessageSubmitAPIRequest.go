package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenmallRefundMessageSubmitAPIRequest
提交退款单留言 API请求
taobao.openmall.refund.message.submit

OpenMall业务提交退款单留言 */
type TaobaoOpenmallRefundMessageSubmitAPIRequest struct {
	model.Params
	// 分销者身份
	_distributor string
	// 退款单ID
	_refundId int64
	// 提交留言结构
	_refundMessage *RefundMessage
}

// New
