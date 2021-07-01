package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappTemplateOnlineappAPIRequest
上线实例化应用 API请求
taobao.miniapp.template.onlineapp

将指定的预览版本发布上线，预览版本号由构建实例化或更新实例化接口返回。 */
type TaobaoMiniappTemplateOnlineappAPIRequest struct {
	model.Params
	// 要更新的投放端,目前可投放： taobao(淘宝),tmall(天猫)
	_clients []string
	// 小程序app_id
	_appId string
	// 模板id
	_templateId string
	// 模板版本
	_templateVersion string
	// 待上线的预览版本号
	_appVersion string
}

// New
