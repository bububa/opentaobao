package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkUmsReturnitemsGet 退货库位商品查询（退货出库辅助）-回流单
// alibaba.wdk.ums.returnitems.get
//
// 退货库位商品查询（退货出库辅助）-回流单
func AlibabaWdkUmsReturnitemsGet(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkUmsReturnitemsGetAPIRequest, resp *wdk.AlibabaWdkUmsReturnitemsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
