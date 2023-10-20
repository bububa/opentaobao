package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMosFinanceBankinfoQuerybank 供应商银行账号查询
// alibaba.mos.finance.bankinfo.querybank
//
// 查询供应商对应的银行账号信息
func AlibabaMosFinanceBankinfoQuerybank(clt *core.SDKClient, req *mos.AlibabaMosFinanceBankinfoQuerybankAPIRequest, resp *mos.AlibabaMosFinanceBankinfoQuerybankAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
