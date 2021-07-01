package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

/* AlibabaMosOnsiteTradeRefund
退款
alibaba.mos.onsite.trade.refund

当交易发生之后一段时间内，由于消费者或者商户的原因需退款，商户可通过退款接口将支付款退还给消费者，喵街将在收到退款请求并验证成功后，按退款规则将支付款按原路退到消费者账号上。

1. 交易超过可退款时间（签约时设置的可退款时间）的订单无法进行退款。
2. 只支持全额退款。 */
func AlibabaMosOnsiteTradeRefund(clt *core.SDKClient, req *mos.AlibabaMosOnsiteTradeRefundAPIRequest, session string) (*mos.AlibabaMosOnsiteTradeRefundAPIResponse, error) {
	var resp mos.AlibabaMosOnsiteTradeRefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
