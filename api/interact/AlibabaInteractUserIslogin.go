package interact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractUserIslogin 校验用户是否已经登录
// alibaba.interact.user.islogin
//
// API的功能是校验用户是否登录，ISV调用接口的时候，通过此接口映射到mtop.interact.user.islogin接口上，因此接口只是做一个给ISV注册调用api的入口，没有发生真实的RPC
func AlibabaInteractUserIslogin(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractUserIsloginAPIRequest, resp *interact.AlibabaInteractUserIsloginAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
