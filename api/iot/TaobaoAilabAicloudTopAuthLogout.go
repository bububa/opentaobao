package iot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopAuthLogout 登出
// taobao.ailab.aicloud.top.auth.logout
//
// 登出
func TaobaoAilabAicloudTopAuthLogout(ctx context.Context, clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopAuthLogoutAPIRequest, resp *iot.TaobaoAilabAicloudTopAuthLogoutAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
