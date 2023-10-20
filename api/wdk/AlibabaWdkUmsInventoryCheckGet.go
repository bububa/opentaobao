package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkUmsInventoryCheckGet 盘点结果单-回流单
// alibaba.wdk.ums.inventory.check.get
//
// 盘点结果单-回流单
func AlibabaWdkUmsInventoryCheckGet(clt *core.SDKClient, req *wdk.AlibabaWdkUmsInventoryCheckGetAPIRequest, resp *wdk.AlibabaWdkUmsInventoryCheckGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
