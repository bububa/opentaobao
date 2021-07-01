package alscmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscMerchantExtTicketRefundAPIRequest
口碑凭证售后退 API请求
alibaba.alsc.merchant.ext.ticket.refund

口碑凭证售后退 */
type AlibabaAlscMerchantExtTicketRefundAPIRequest struct {
	model.Params
	// 券核销流水号，针对该次核销发起售后退操作
	_transId string
	// 外部请求号，支持英文字母和数字，由开发者自行定义（不允许重复）
	_ticketRequestId string
	// 凭证码，包括内部凭证码和外部凭证码，内部凭证码为12位，纯数字，且唯一不重复
	_ticketCode string
}

// New
