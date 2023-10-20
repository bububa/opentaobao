package refund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/refund"
)

// TaobaoRefundsReceiveGet 查询卖家收到的退款列表
// taobao.refunds.receive.get
//
// 查询卖家收到的退款列表
func TaobaoRefundsReceiveGet(clt *core.SDKClient, req *refund.TaobaoRefundsReceiveGetAPIRequest, session string) (*refund.TaobaoRefundsReceiveGetAPIResponse, error) {
	var resp refund.TaobaoRefundsReceiveGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
