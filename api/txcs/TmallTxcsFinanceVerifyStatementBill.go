package txcs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/txcs"
)

// Tmalltxcsfinanceverifystatementbill 供应商核销单录入
// tmall.txcs.finance.verify.statement.bill
//
// 供应商核销单录入
func Tmalltxcsfinanceverifystatementbill(clt *core.SDKClient, req *txcs.TmalltxcsfinanceverifystatementbillAPIRequest, session string) (*txcs.TmalltxcsfinanceverifystatementbillAPIResponse, error) {
	var resp txcs.TmalltxcsfinanceverifystatementbillAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
