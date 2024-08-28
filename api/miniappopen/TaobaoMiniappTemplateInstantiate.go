package miniappopen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniappTemplateInstantiate 构建实例化应用
// taobao.miniapp.template.instantiate
//
// 实例化saas化的小程序
func TaobaoMiniappTemplateInstantiate(ctx context.Context, clt *core.SDKClient, req *miniappopen.TaobaoMiniappTemplateInstantiateAPIRequest, resp *miniappopen.TaobaoMiniappTemplateInstantiateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
