package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniappTemplateOfflineapp 下线实例化应用
// taobao.miniapp.template.offlineapp
//
// 对指定的实例化小程序进行下线,需要指定clients和app_version
func TaobaoMiniappTemplateOfflineapp(clt *core.SDKClient, req *miniappopen.TaobaoMiniappTemplateOfflineappAPIRequest, resp *miniappopen.TaobaoMiniappTemplateOfflineappAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
