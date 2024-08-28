package miniappopen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniappWidgetTemplateInstanceQuery 小部件实例化版本查询
// taobao.miniapp.widget.template.instance.query
//
// 小部件实例化版本查询
func TaobaoMiniappWidgetTemplateInstanceQuery(ctx context.Context, clt *core.SDKClient, req *miniappopen.TaobaoMiniappWidgetTemplateInstanceQueryAPIRequest, resp *miniappopen.TaobaoMiniappWidgetTemplateInstanceQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
