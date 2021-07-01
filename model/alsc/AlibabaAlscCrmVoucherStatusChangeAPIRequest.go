package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmVoucherStatusChangeAPIRequest
优惠券状态更改 API请求
alibaba.alsc.crm.voucher.status.change

核销优惠券 */
type AlibabaAlscCrmVoucherStatusChangeAPIRequest struct {
	model.Params
	// 参数
	_paramVoucherStatusChangeOpenReq *VoucherStatusChangeOpenReq
}

// New
