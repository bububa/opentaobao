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

// NewTaobaoMiniapppTemplateInstantiateRequest 初始化TaobaoMiniapppTemplateInstantiateAPIRequest对象
func NewTaobaoMiniapppTemplateInstantiateRequest() *TaobaoMiniapppTemplateInstantiateAPIRequest {
	return &TaobaoMiniapppTemplateInstantiateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniapppTemplateInstantiateAPIRequest) GetApiMethodName() string {
	return "taobao.miniappp.template.instantiate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniapppTemplateInstantiateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Clients Setter
// 投放端,目前可投放： taobao(淘宝),tmall(天猫)
func (r *TaobaoMiniapppTemplateInstantiateAPIRequest) SetClients(_clients []string) error {
	r._clients = _clients
	r.Set("clients", _clients)
	return nil
}

// Get Clients Getter
func (r TaobaoMiniapppTemplateInstantiateAPIRequest) GetClients() []string {
	return r._clients
}

// Set is Description Setter
// 描述长度(9~200)
func (r *TaobaoMiniapppTemplateInstantiateAPIRequest) SetDescription(_description string) error {
	r._description = _description
	r.Set("description", _description)
	return nil
}

// Get Description Getter
func (r TaobaoMiniapppTemplateInstantiateAPIRequest) GetDescription() string {
	return r._description
}

// Set is ExtJson Setter
// schemadata, json字符串
func (r *TaobaoMiniapppTemplateInstantiateAPIRequest) SetExtJson(_extJson string) error {
	r._extJson = _extJson
	r.Set("ext_json", _extJson)
	return nil
}

// Get ExtJson Getter
func (r TaobaoMiniapppTemplateInstantiateAPIRequest) GetExtJson() string {
	return r._extJson
}

// Set is Icon Setter
// 小程序icon
func (r *TaobaoMiniapppTemplateInstantiateAPIRequest) SetIcon(_icon string) error {
	r._icon = _icon
	r.Set("icon", _icon)
	return nil
}

// Get Icon Getter
func (r TaobaoMiniapppTemplateInstantiateAPIRequest) GetIcon() string {
	return r._icon
}

// Set is Name Setter
// 小程序名称
func (r *TaobaoMiniapppTemplateInstantiateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// Get Name Getter
func (r TaobaoMiniapppTemplateInstantiateAPIRequest) GetName() string {
	return r._name
}

// Set is TemplateId Setter
// 模板id
func (r *TaobaoMiniapppTemplateInstantiateAPIRequest) SetTemplateId(_templateId string) error {
	r._templateId = _templateId
	r.Set("template_id", _templateId)
	return nil
}

// Get TemplateId Getter
func (r TaobaoMiniapppTemplateInstantiateAPIRequest) GetTemplateId() string {
	return r._templateId
}

// Set is TemplateVersion Setter
// 模板版本
func (r *TaobaoMiniapppTemplateInstantiateAPIRequest) SetTemplateVersion(_templateVersion string) error {
	r._templateVersion = _templateVersion
	r.Set("template_version", _templateVersion)
	return nil
}

// Get TemplateVersion Getter
func (r TaobaoMiniapppTemplateInstantiateAPIRequest) GetTemplateVersion() string {
	return r._templateVersion
}
