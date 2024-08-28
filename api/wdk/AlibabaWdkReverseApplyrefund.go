package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkReverseApplyrefund 逆向申请接口
// alibaba.wdk.reverse.applyrefund
//
// 逆向渲染
func AlibabaWdkReverseApplyrefund(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkReverseApplyrefundAPIRequest, resp *wdk.AlibabaWdkReverseApplyrefundAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
