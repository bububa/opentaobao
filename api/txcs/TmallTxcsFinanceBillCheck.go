package txcs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/txcs"
)

// Tmalltxcsfinancebillcheck 天猫超市外部商家财务账单对账
// tmall.txcs.finance.bill.check
//
// 提供天猫超市外部合作商家财务账单对账
func Tmalltxcsfinancebillcheck(clt *core.SDKClient, req *txcs.TmalltxcsfinancebillcheckAPIRequest, session string) (*txcs.TmalltxcsfinancebillcheckAPIResponse, error) {
	var resp txcs.TmalltxcsfinancebillcheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
