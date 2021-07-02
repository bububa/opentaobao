package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniappTemplateUpdate （已废弃）更新实例化应用
// taobao.miniapp.template.update
//
// 商家应用c端模板实例化小程序更新
func TaobaoMiniappTemplateUpdate(clt *core.SDKClient, req *miniappopen.TaobaoMiniappTemplateUpdateAPIRequest, session string) (*miniappopen.TaobaoMiniappTemplateUpdateAPIResponse, error) {
	var resp miniappopen.TaobaoMiniappTemplateUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
