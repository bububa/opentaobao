package ottpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ottpay"
)

// YoukuOttPayOrderAuthpaywithprice 委托代扣可配定价服务
// youku.ott.pay.order.authpaywithprice
//
// 应用中心sdk连续包月委托代扣服务，次月可配置营销价
func YoukuOttPayOrderAuthpaywithprice(clt *core.SDKClient, req *ottpay.YoukuOttPayOrderAuthpaywithpriceAPIRequest, resp *ottpay.YoukuOttPayOrderAuthpaywithpriceAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
