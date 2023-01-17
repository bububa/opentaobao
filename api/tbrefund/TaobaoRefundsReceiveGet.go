package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoRefundsReceiveGet 查询卖家收到的退款列表
// taobao.refunds.receive.get
//
// 查询卖家收到的退款列表
func TaobaoRefundsReceiveGet(clt *core.SDKClient, req *tbrefund.TaobaoRefundsReceiveGetAPIRequest, session string) (*tbrefund.TaobaoRefundsReceiveGetAPIResponse, error) {
	var resp tbrefund.TaobaoRefundsReceiveGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
