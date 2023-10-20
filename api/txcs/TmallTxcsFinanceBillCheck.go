package txcs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/txcs"
)

// TmallTxcsFinanceBillCheck 天猫超市外部商家财务账单对账
// tmall.txcs.finance.bill.check
//
// 提供天猫超市外部合作商家财务账单对账
func TmallTxcsFinanceBillCheck(clt *core.SDKClient, req *txcs.TmallTxcsFinanceBillCheckAPIRequest, resp *txcs.TmallTxcsFinanceBillCheckAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
