package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMjMosFundCancelbillAPIRequest
取消付款单 API请求
alibaba.mj.mos.fund.cancelbill

取消付款单 */
type AlibabaMjMosFundCancelbillAPIRequest struct {
	model.Params
	// 取消入参
	_cancelBillDTO *CancelBillDto
}

// New
