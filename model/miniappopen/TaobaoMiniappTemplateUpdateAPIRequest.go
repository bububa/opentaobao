package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappTemplateUpdateAPIRequest
（已废弃）更新实例化应用 API请求
taobao.miniapp.template.update

商家应用c端模板实例化小程序更新 */
type TaobaoMiniappTemplateUpdateAPIRequest struct {
	model.Params
	// 要更新的投放端,目前可投放： taobao(淘宝),tmall(天猫)
	_clients []string
	// 应用id
	_id string
	// schema信息，不填且 应用线上版本使用的templateId与传入的templateId不一致，则会报错; 一致，则复用线上版本的schema。
	_extJson string
	// 模板id
	_templateId string
	// 模板版本
	_templateVersion string
}

// New
