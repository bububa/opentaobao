package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmVoucherSendAPIRequest
发送券给指定用户 API请求
alibaba.alsc.crm.voucher.send

发送券给指定用户 */
type AlibabaAlscCrmVoucherSendAPIRequest struct {
	model.Params
	// 请求参数
	_paramVoucherSendOpenReq *VoucherSendOpenReq
}

// New
