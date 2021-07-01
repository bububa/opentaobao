package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosFinanceBankinfoQuerybankAPIRequest
供应商银行账号查询 API请求
alibaba.mos.finance.bankinfo.querybank

查询供应商对应的银行账号信息 */
type AlibabaMosFinanceBankinfoQuerybankAPIRequest struct {
	model.Params
	// 供应商id
	_supplierId string
	// 门店号
	_storeNo string
	// 签约主体id
	_companyId string
}

// New
