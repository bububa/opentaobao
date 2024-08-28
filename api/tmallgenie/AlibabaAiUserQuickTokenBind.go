package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAiUserQuickTokenBind 人工智能实验室精灵用户绑定第三方Token接口
// alibaba.ai.user.quick.token.bind
//
// 人工智能实验室精灵用户绑定第三方Token接口
func AlibabaAiUserQuickTokenBind(ctx context.Context, clt *core.SDKClient, req *tmallgenie.AlibabaAiUserQuickTokenBindAPIRequest, resp *tmallgenie.AlibabaAiUserQuickTokenBindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
