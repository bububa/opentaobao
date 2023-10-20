package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMjOcPay POS收银成功后订单同步
// alibaba.mj.oc.pay
//
// 此API用于在银泰商场中，消费者在收银台收银/退款时， POS系统在收银或退款成功后，调用此接口进行订单同步
func AlibabaMjOcPay(clt *core.SDKClient, req *mos.AlibabaMjOcPayAPIRequest, resp *mos.AlibabaMjOcPayAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
