package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappTemplateRollbackAPIRequest
回滚实例化应用 API请求
taobao.miniapp.template.rollback

将实例化小程序回滚到指定版本 */
type TaobaoMiniappTemplateRollbackAPIRequest struct {
	model.Params
	// 要回滚的投放端,目前可投放： taobao,tmall
	_clients []string
	// 小程序app_id
	_appId string
	// 回到到该app_version
	_appVersion string
	// 实例化小程序对应的模板id
	_templateId string
	// 与app_version对应的模板版本
	_templateVersion string
}

// New
