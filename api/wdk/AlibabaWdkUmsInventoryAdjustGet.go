package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkUmsInventoryAdjustGet 库调单-回流单
// alibaba.wdk.ums.inventory.adjust.get
//
// 库调单-回流单
func AlibabaWdkUmsInventoryAdjustGet(clt *core.SDKClient, req *wdk.AlibabaWdkUmsInventoryAdjustGetAPIRequest, resp *wdk.AlibabaWdkUmsInventoryAdjustGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
