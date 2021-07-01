package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniapppTemplateInstantiateAPIRequest
（已废弃）构建实例化应用 API请求
taobao.miniappp.template.instantiate

实例化saas化的小程序 */
type TaobaoMiniapppTemplateInstantiateAPIRequest struct {
	model.Params
	// 投放端,目前可投放： taobao(淘宝),tmall(天猫)
	_clients []string
	// 描述长度(9~200)
	_description string
	// schemadata, json字符串
	_extJson string
	// 小程序icon
	_icon string
	// 小程序名称
	_name string
	// 模板id
	_templateId string
	// 模板版本
	_templateVersion string
}

// New
