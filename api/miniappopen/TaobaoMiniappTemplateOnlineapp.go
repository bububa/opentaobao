package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniappTemplateOnlineapp 上线实例化应用
// taobao.miniapp.template.onlineapp
//
// 将指定的预览版本发布上线，预览版本号由构建实例化或更新实例化接口返回。
func TaobaoMiniappTemplateOnlineapp(clt *core.SDKClient, req *miniappopen.TaobaoMiniappTemplateOnlineappAPIRequest, resp *miniappopen.TaobaoMiniappTemplateOnlineappAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
