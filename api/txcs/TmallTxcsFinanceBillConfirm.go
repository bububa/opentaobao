package txcs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/txcs"
)

// Tmalltxcsfinancebillconfirm 供应商账单确认
// tmall.txcs.finance.bill.confirm
//
// 提供天猫超市外部合作商家：财务账单对账
func Tmalltxcsfinancebillconfirm(clt *core.SDKClient, req *txcs.TmalltxcsfinancebillconfirmAPIRequest, session string) (*txcs.TmalltxcsfinancebillconfirmAPIResponse, error) {
	var resp txcs.TmalltxcsfinancebillconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
