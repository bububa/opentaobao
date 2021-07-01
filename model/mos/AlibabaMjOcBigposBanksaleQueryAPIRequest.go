package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMjOcBigposBanksaleQueryAPIRequest
大pos银行卡查账接口 API请求
alibaba.mj.oc.bigpos.banksale.query

大pos银行卡查账接口，给收银员查询银行卡销售记录，便于调账 */
type AlibabaMjOcBigposBanksaleQueryAPIRequest struct {
	model.Params
	// 开始时间
	_startTime string
	// 卡号
	_cardNo string
	// 外部门店号
	_outStoreNo string
	// 结束时间
	_endTime string
}

// New
