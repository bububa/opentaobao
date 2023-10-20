package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniappTemplateQueryapp 查询实例化应用版本
// taobao.miniapp.template.queryapp
//
// 根据模板id和商家信息，查询实例化小程序版本查询
func TaobaoMiniappTemplateQueryapp(clt *core.SDKClient, req *miniappopen.TaobaoMiniappTemplateQueryappAPIRequest, resp *miniappopen.TaobaoMiniappTemplateQueryappAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
