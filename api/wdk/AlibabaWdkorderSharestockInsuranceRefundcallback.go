package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkorderSharestockInsuranceRefundcallback 共享库存逆向订单理赔单回传
// alibaba.wdkorder.sharestock.insurance.refundcallback
//
// 共享库存逆向订单理赔单回传
func AlibabaWdkorderSharestockInsuranceRefundcallback(clt *core.SDKClient, req *wdk.AlibabaWdkorderSharestockInsuranceRefundcallbackAPIRequest, resp *wdk.AlibabaWdkorderSharestockInsuranceRefundcallbackAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
