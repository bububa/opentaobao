package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappTemplateUpdateappAPIRequest
更新实例化应用 API请求
taobao.miniapp.template.updateapp

商家应用c端模板实例化小程序更新，生成新的版本，但不会自动上线新版本 */
type TaobaoMiniappTemplateUpdateappAPIRequest struct {
	model.Params
	// 要更新的投放端,目前可投放： taobao(淘宝),tmall(天猫)
	_clients []string
	// 应用id，如果应用在对应端上已有的最新版本所使用的模板id,模板version和extjson，与本次入参一致，则认为不需要更新，返回已有的版本。
	_appId string
	// 扩展信息。线上版本使用的template_id与传入的template_id一致时，可不填。
	_extJson string
	// 模板id
	_templateId string
	// 模板版本
	_templateVersion string
	// Logo更新，1月5次
	_icon string
	// 描述更新，1年5次
	_desc string
	// 简称更新，1年5次
	_alias string
}

// New
