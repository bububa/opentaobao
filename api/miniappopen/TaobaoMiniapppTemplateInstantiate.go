package miniappopen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniapppTemplateInstantiate （已废弃）构建实例化应用
// taobao.miniappp.template.instantiate
//
// 实例化saas化的小程序
func TaobaoMiniapppTemplateInstantiate(ctx context.Context, clt *core.SDKClient, req *miniappopen.TaobaoMiniapppTemplateInstantiateAPIRequest, resp *miniappopen.TaobaoMiniapppTemplateInstantiateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
