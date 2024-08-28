package miniappopen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniappTemplateUpdate （已废弃）更新实例化应用
// taobao.miniapp.template.update
//
// 商家应用c端模板实例化小程序更新
func TaobaoMiniappTemplateUpdate(ctx context.Context, clt *core.SDKClient, req *miniappopen.TaobaoMiniappTemplateUpdateAPIRequest, resp *miniappopen.TaobaoMiniappTemplateUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
