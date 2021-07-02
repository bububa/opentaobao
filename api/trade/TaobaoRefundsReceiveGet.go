package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TaobaoRefundsReceiveGet 查询卖家收到的退款列表
// taobao.refunds.receive.get
//
// 查询卖家收到的退款列表
func TaobaoRefundsReceiveGet(clt *core.SDKClient, req *trade.TaobaoRefundsReceiveGetAPIRequest, session string) (*trade.TaobaoRefundsReceiveGetAPIResponse, error) {
	var resp trade.TaobaoRefundsReceiveGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
