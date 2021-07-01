package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMjMosFundModifybillbankaccountAPIRequest
修改付款单的银行账户信息 API请求
alibaba.mj.mos.fund.modifybillbankaccount

修改付款单的银行账户信息 */
type AlibabaMjMosFundModifybillbankaccountAPIRequest struct {
	model.Params
	// 修改入参
	_modifyDto *ModifyBillDto
}

// New
