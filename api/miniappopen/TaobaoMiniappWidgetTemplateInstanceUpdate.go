package miniappopen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniappWidgetTemplateInstanceUpdate 小部件实例化版本更新
// taobao.miniapp.widget.template.instance.update
//
// 小部件版本更新
func TaobaoMiniappWidgetTemplateInstanceUpdate(ctx context.Context, clt *core.SDKClient, req *miniappopen.TaobaoMiniappWidgetTemplateInstanceUpdateAPIRequest, resp *miniappopen.TaobaoMiniappWidgetTemplateInstanceUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
