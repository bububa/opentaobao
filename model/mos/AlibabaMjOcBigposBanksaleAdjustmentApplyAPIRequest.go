package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMjOcBigposBanksaleAdjustmentApplyAPIRequest
大pos银行卡调账申请 API请求
alibaba.mj.oc.bigpos.banksale.adjustment.apply

大pos银行卡调账申请 */
type AlibabaMjOcBigposBanksaleAdjustmentApplyAPIRequest struct {
	model.Params
	// 门店号
	_storeNo string
	// 调账金额
	_amount int64
	// 卡号
	_cardNo string
	// 交易时间
	_operTime string
	// 收银员号
	_operator string
	// 调账收银机号
	_posNo string
}

// New
