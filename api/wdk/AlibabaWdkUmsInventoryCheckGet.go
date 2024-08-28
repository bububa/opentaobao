package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkUmsInventoryCheckGet 盘点结果单-回流单
// alibaba.wdk.ums.inventory.check.get
//
// 盘点结果单-回流单
func AlibabaWdkUmsInventoryCheckGet(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkUmsInventoryCheckGetAPIRequest, resp *wdk.AlibabaWdkUmsInventoryCheckGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
