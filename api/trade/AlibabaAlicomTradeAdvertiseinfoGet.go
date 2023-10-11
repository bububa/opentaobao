package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// AlibabaAlicomTradeAdvertiseinfoGet 获取订单上的在信息流投放信息
// alibaba.alicom.trade.advertiseinfo.get
//
// 获取订单上的在信息流投放信息
func AlibabaAlicomTradeAdvertiseinfoGet(clt *core.SDKClient, req *trade.AlibabaAlicomTradeAdvertiseinfoGetAPIRequest, session string) (*trade.AlibabaAlicomTradeAdvertiseinfoGetAPIResponse, error) {
	var resp trade.AlibabaAlicomTradeAdvertiseinfoGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
