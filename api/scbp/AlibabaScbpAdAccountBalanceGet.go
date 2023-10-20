package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdAccountBalanceGet 查询账户余额
// alibaba.scbp.ad.account.balance.get
//
// 查询推广账户余额
func AlibabaScbpAdAccountBalanceGet(clt *core.SDKClient, req *scbp.AlibabaScbpAdAccountBalanceGetAPIRequest, resp *scbp.AlibabaScbpAdAccountBalanceGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
