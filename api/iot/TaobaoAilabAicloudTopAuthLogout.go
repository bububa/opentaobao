package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopAuthLogout 登出
// taobao.ailab.aicloud.top.auth.logout
//
// 登出
func TaobaoAilabAicloudTopAuthLogout(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopAuthLogoutAPIRequest, resp *iot.TaobaoAilabAicloudTopAuthLogoutAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
