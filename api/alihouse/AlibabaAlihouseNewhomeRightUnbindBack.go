package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeRightUnbindBack 权限回流-解绑
// alibaba.alihouse.newhome.right.unbind.back
//
// 权限回流-解绑
func AlibabaAlihouseNewhomeRightUnbindBack(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeRightUnbindBackAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeRightUnbindBackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
