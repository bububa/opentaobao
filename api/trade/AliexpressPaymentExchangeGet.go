package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Aliexpresspaymentexchangeget getExchange
// aliexpress.payment.exchange.get
//
// 提供国际汇率服务
func Aliexpresspaymentexchangeget(clt *core.SDKClient, req *trade.AliexpresspaymentexchangegetAPIRequest, session string) (*trade.AliexpresspaymentexchangegetAPIResponse, error) {
	var resp trade.AliexpresspaymentexchangegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
