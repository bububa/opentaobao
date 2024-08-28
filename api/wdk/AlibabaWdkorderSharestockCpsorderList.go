package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkorderSharestockCpsorderList cps正向分销订单批量回流
// alibaba.wdkorder.sharestock.cpsorder.list
//
// cps正向分销订单批量回流
func AlibabaWdkorderSharestockCpsorderList(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkorderSharestockCpsorderListAPIRequest, resp *wdk.AlibabaWdkorderSharestockCpsorderListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
