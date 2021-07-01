package alscmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscMerchantExtTicketcodeUseAPIRequest
外部核销服务 API请求
alibaba.alsc.merchant.ext.ticketcode.use

外部核销服务 */
type AlibabaAlscMerchantExtTicketcodeUseAPIRequest struct {
	model.Params
	// 外部券使用请求
	_useRequest *ExternalTicketUseRequest
}

// New
