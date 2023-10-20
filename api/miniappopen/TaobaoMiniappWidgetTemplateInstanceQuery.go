package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniappWidgetTemplateInstanceQuery 小部件实例化版本查询
// taobao.miniapp.widget.template.instance.query
//
// 小部件实例化版本查询
func TaobaoMiniappWidgetTemplateInstanceQuery(clt *core.SDKClient, req *miniappopen.TaobaoMiniappWidgetTemplateInstanceQueryAPIRequest, resp *miniappopen.TaobaoMiniappWidgetTemplateInstanceQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
