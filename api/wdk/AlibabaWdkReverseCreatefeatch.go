package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkReverseCreatefeatch 逆向取货
// alibaba.wdk.reverse.createfeatch
//
// 发起逆向取货
func AlibabaWdkReverseCreatefeatch(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkReverseCreatefeatchAPIRequest, resp *wdk.AlibabaWdkReverseCreatefeatchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
