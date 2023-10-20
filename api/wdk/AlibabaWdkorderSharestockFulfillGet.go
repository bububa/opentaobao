package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkorderSharestockFulfillGet 商户订单履约数据获取
// alibaba.wdkorder.sharestock.fulfill.get
//
// 商户订单履约数据获取
func AlibabaWdkorderSharestockFulfillGet(clt *core.SDKClient, req *wdk.AlibabaWdkorderSharestockFulfillGetAPIRequest, resp *wdk.AlibabaWdkorderSharestockFulfillGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
