package ottpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ottpay"
)

// Youkuottpayorderauthpay 委托代扣服务
// youku.ott.pay.order.authpay
//
// 应用中心sdk连续包月委托代扣服务
func Youkuottpayorderauthpay(clt *core.SDKClient, req *ottpay.YoukuottpayorderauthpayAPIRequest, session string) (*ottpay.YoukuottpayorderauthpayAPIResponse, error) {
	var resp ottpay.YoukuottpayorderauthpayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
