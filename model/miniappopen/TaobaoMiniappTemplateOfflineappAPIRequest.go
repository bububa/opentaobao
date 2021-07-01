package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappTemplateOfflineappAPIRequest
下线实例化应用 API请求
taobao.miniapp.template.offlineapp

对指定的实例化小程序进行下线,需要指定clients和app_version */
type TaobaoMiniappTemplateOfflineappAPIRequest struct {
	model.Params
	// 要下线的投放端,目前可投放： taobao(淘宝),tmall(天猫)
	_clients []string
	// 要下线的小程序app_id
	_appId string
	// 要下线的小程序版本号
	_appVersion string
	// 模板id
	_templateId string
}

// New
