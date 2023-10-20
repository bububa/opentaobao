package exchange

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/exchange"
)

// TmallExchangeRefusereasonGet 获取拒绝换货原因列表
// tmall.exchange.refusereason.get
//
// 获取拒绝换货原因列表
func TmallExchangeRefusereasonGet(clt *core.SDKClient, req *exchange.TmallExchangeRefusereasonGetAPIRequest, session string) (*exchange.TmallExchangeRefusereasonGetAPIResponse, error) {
	var resp exchange.TmallExchangeRefusereasonGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
