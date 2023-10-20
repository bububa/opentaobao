package txcs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/txcs"
)

// TmallTxcsFinanceBillQuery 天猫超市外部商家财务账单信息查询
// tmall.txcs.finance.bill.query
//
// 提供天猫超市外部合作商家财务账单对账
func TmallTxcsFinanceBillQuery(clt *core.SDKClient, req *txcs.TmallTxcsFinanceBillQueryAPIRequest, resp *txcs.TmallTxcsFinanceBillQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
