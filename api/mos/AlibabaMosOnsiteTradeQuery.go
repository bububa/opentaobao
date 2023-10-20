package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamosonsitetradequery 新商场当面付交易查询
// alibaba.mos.onsite.trade.query
//
// 本接口提供新商场当面付订单的查询的功能，商户可以通过本接口主动查询订单状态，完成下一步的业务逻辑。
// 商户系统应在两种场景下调用此接口: 商户POS系统应该在调用[条码支付请求接口]并成功返回后，调用此接口查询订单的处理状态。
func Alibabamosonsitetradequery(clt *core.SDKClient, req *mos.AlibabamosonsitetradequeryAPIRequest, session string) (*mos.AlibabamosonsitetradequeryAPIResponse, error) {
	var resp mos.AlibabamosonsitetradequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
