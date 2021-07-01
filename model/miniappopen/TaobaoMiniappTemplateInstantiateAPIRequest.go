package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappTemplateInstantiateAPIRequest
构建实例化应用 API请求
taobao.miniapp.template.instantiate

实例化saas化的小程序 */
type TaobaoMiniappTemplateInstantiateAPIRequest struct {
	model.Params
	// 投放端,目前可投放： taobao(淘宝),tmall(天猫)，taobao为必填，需要模板在这些端上已经发布上线
	_clients []string
	// 描述长度(9~200)
	_description string
	// 扩展信息，JSON格式。
	_extJson string
	// 小程序icon
	_icon string
	// 小程序名称按平台规则自动生成，该字段仅做兜底使用。
	_name string
	// 模板id
	_templateId string
	// 模板版本
	_templateVersion string
	// 小程序简称【1-16】字符，可重名
	_alias string
}

// New
