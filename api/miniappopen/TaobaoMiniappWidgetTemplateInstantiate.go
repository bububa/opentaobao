package miniappopen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniappWidgetTemplateInstantiate 小部件实例化接口
// taobao.miniapp.widget.template.instantiate
//
// 小部件实例化接口
func TaobaoMiniappWidgetTemplateInstantiate(ctx context.Context, clt *core.SDKClient, req *miniappopen.TaobaoMiniappWidgetTemplateInstantiateAPIRequest, resp *miniappopen.TaobaoMiniappWidgetTemplateInstantiateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
