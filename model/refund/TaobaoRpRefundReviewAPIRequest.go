package refund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRpRefundReviewAPIRequest
审核退款单 API请求
taobao.rp.refund.review

审核退款单，标志是否可用于批量退款，目前仅支持天猫订单。 */
type TaobaoRpRefundReviewAPIRequest struct {
	model.Params
	// 退款单编号
	_refundId int64
	// 审核人姓名
	_operator string
	// 退款阶段，可选值：售中：onsale，售后：aftersale
	_refundPhase string
	// 退款最后更新时间，以时间戳的方式表示
	_refundVersion int64
	// 审核是否可用于批量退款，可选值：true（审核通过），false（审核不通过或反审核）
	_result bool
	// 审核留言
	_message string
}

// New
