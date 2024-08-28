package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkReverseLoadFeatchorder 取货详情
// alibaba.wdk.reverse.load.featchorder
//
// 取货详情
func AlibabaWdkReverseLoadFeatchorder(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkReverseLoadFeatchorderAPIRequest, resp *wdk.AlibabaWdkReverseLoadFeatchorderAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
