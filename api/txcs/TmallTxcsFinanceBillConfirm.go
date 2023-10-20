package txcs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/txcs"
)

// TmallTxcsFinanceBillConfirm 供应商账单确认
// tmall.txcs.finance.bill.confirm
//
// 提供天猫超市外部合作商家：财务账单对账
func TmallTxcsFinanceBillConfirm(clt *core.SDKClient, req *txcs.TmallTxcsFinanceBillConfirmAPIRequest, resp *txcs.TmallTxcsFinanceBillConfirmAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
