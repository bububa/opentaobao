package exchange

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/exchange"
)

// TmallExchangeGet 获取单笔换货详情
// tmall.exchange.get
//
// 获取单笔换货详情
func TmallExchangeGet(clt *core.SDKClient, req *exchange.TmallExchangeGetAPIRequest, session string) (*exchange.TmallExchangeGetAPIResponse, error) {
	var resp exchange.TmallExchangeGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
