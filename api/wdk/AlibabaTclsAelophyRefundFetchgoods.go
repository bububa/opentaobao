package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyRefundFetchgoods saas 售后逆向 商户发起逆向取货
// alibaba.tcls.aelophy.refund.fetchgoods
//
// saas 售后逆向 商户发起逆向取货
func AlibabaTclsAelophyRefundFetchgoods(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyRefundFetchgoodsAPIRequest, resp *wdk.AlibabaTclsAelophyRefundFetchgoodsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
