package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniappWidgetTemplateInstanceQuery 小部件实例化版本查询
// taobao.miniapp.widget.template.instance.query
//
// 小部件实例化版本查询
func TaobaoMiniappWidgetTemplateInstanceQuery(clt *core.SDKClient, req *miniappopen.TaobaoMiniappWidgetTemplateInstanceQueryAPIRequest, session string) (*miniappopen.TaobaoMiniappWidgetTemplateInstanceQueryAPIResponse, error) {
	var resp miniappopen.TaobaoMiniappWidgetTemplateInstanceQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
