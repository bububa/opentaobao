package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyRefundCsapplyrender 商家代客售后逆向申请渲染获取
// alibaba.tcls.aelophy.refund.csapplyrender
//
// 提供商家代客售后逆向申请渲染获取的接口
func AlibabaTclsAelophyRefundCsapplyrender(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyRefundCsapplyrenderAPIRequest, resp *wdk.AlibabaTclsAelophyRefundCsapplyrenderAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
