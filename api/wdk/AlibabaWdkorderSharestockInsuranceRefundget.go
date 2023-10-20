package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkorderSharestockInsuranceRefundget 共享库存投保业务售后逆向订单数据获取
// alibaba.wdkorder.sharestock.insurance.refundget
//
// 共享库存投保业务售后逆向订单数据获取
func AlibabaWdkorderSharestockInsuranceRefundget(clt *core.SDKClient, req *wdk.AlibabaWdkorderSharestockInsuranceRefundgetAPIRequest, resp *wdk.AlibabaWdkorderSharestockInsuranceRefundgetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
