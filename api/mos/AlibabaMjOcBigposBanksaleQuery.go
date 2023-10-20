package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMjOcBigposBanksaleQuery 大pos银行卡查账接口
// alibaba.mj.oc.bigpos.banksale.query
//
// 大pos银行卡查账接口，给收银员查询银行卡销售记录，便于调账
func AlibabaMjOcBigposBanksaleQuery(clt *core.SDKClient, req *mos.AlibabaMjOcBigposBanksaleQueryAPIRequest, resp *mos.AlibabaMjOcBigposBanksaleQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
