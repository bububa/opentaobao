package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamjocpay POS收银成功后订单同步
// alibaba.mj.oc.pay
//
// 此API用于在银泰商场中，消费者在收银台收银/退款时， POS系统在收银或退款成功后，调用此接口进行订单同步
func Alibabamjocpay(clt *core.SDKClient, req *mos.AlibabamjocpayAPIRequest, session string) (*mos.AlibabamjocpayAPIResponse, error) {
	var resp mos.AlibabamjocpayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
