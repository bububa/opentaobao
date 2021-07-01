package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

/* AlibabaMosFinanceBankinfoQuerybank
供应商银行账号查询
alibaba.mos.finance.bankinfo.querybank

查询供应商对应的银行账号信息 */
func AlibabaMosFinanceBankinfoQuerybank(clt *core.SDKClient, req *mos.AlibabaMosFinanceBankinfoQuerybankAPIRequest, session string) (*mos.AlibabaMosFinanceBankinfoQuerybankAPIResponse, error) {
	var resp mos.AlibabaMosFinanceBankinfoQuerybankAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
