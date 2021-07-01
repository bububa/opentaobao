package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMjMosFundCreatebillAPIRequest
创建一个付款单 API请求
alibaba.mj.mos.fund.createbill

创建一个付款单 */
type AlibabaMjMosFundCreatebillAPIRequest struct {
	model.Params
	// 创建付款单入参
	_billDto *CreateBillDto
}

// New
