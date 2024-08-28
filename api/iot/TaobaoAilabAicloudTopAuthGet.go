package iot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopAuthGet 登陆
// taobao.ailab.aicloud.top.auth.get
//
// 登陆
func TaobaoAilabAicloudTopAuthGet(ctx context.Context, clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopAuthGetAPIRequest, resp *iot.TaobaoAilabAicloudTopAuthGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
