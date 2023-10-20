package refund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/refund"
)

// Taobaorefundsreceiveget 查询卖家收到的退款列表
// taobao.refunds.receive.get
//
// 查询卖家收到的退款列表
func Taobaorefundsreceiveget(clt *core.SDKClient, req *refund.TaobaorefundsreceivegetAPIRequest, session string) (*refund.TaobaorefundsreceivegetAPIResponse, error) {
	var resp refund.TaobaorefundsreceivegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
