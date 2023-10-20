package ottpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ottpay"
)

// Youkuottpayorderauthpaywithprice 委托代扣可配定价服务
// youku.ott.pay.order.authpaywithprice
//
// 应用中心sdk连续包月委托代扣服务，次月可配置营销价
func Youkuottpayorderauthpaywithprice(clt *core.SDKClient, req *ottpay.YoukuottpayorderauthpaywithpriceAPIRequest, session string) (*ottpay.YoukuottpayorderauthpaywithpriceAPIResponse, error) {
	var resp ottpay.YoukuottpayorderauthpaywithpriceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
