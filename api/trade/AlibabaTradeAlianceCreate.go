package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// AlibabaTradeAlianceCreate 推客平台订单回流
// alibaba.trade.aliance.create
//
// 推客平台订单回流
func AlibabaTradeAlianceCreate(clt *core.SDKClient, req *trade.AlibabaTradeAlianceCreateAPIRequest, session string) (*trade.AlibabaTradeAlianceCreateAPIResponse, error) {
	var resp trade.AlibabaTradeAlianceCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
