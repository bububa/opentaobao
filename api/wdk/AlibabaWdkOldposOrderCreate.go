package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkOldposOrderCreate 淘鲜达外部商户老pos机产生的订单同步进淘鲜达
// alibaba.wdk.oldpos.order.create
//
// 淘鲜达外部商户老pos机产生的订单同步进淘鲜达
func AlibabaWdkOldposOrderCreate(clt *core.SDKClient, req *wdk.AlibabaWdkOldposOrderCreateAPIRequest, resp *wdk.AlibabaWdkOldposOrderCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
