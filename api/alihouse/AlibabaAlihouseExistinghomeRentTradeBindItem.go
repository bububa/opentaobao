package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeRentTradeBindItem 交易绑定商品
// alibaba.alihouse.existinghome.rent.trade.bind.item
//
// 交易绑定商品
func AlibabaAlihouseExistinghomeRentTradeBindItem(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeRentTradeBindItemAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeRentTradeBindItemAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeRentTradeBindItemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
