package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappptemplateinstantiateAPIRequest （已废弃）构建实例化应用 API请求
// taobao.miniappp.template.instantiate
//
// 实例化saas化的小程序
type TaobaominiappptemplateinstantiateAPIRequest struct {
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

// NewTaobaominiappptemplateinstantiateRequest 初始化TaobaominiappptemplateinstantiateAPIRequest对象
func NewTaobaominiappptemplateinstantiateRequest() *TaobaominiappptemplateinstantiateAPIRequest {
	return &TaobaominiappptemplateinstantiateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappptemplateinstantiateAPIRequest) GetApiMethodName() string {
	return "taobao.miniappp.template.instantiate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappptemplateinstantiateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappptemplateinstantiateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClients is Clients Setter
// 投放端,目前可投放： taobao(淘宝),tmall(天猫)
func (r *TaobaominiappptemplateinstantiateAPIRequest) SetClients(_clients []string) error {
	r._clients = _clients
	r.Set("clients", _clients)
	return nil
}

// GetClients Clients Getter
func (r TaobaominiappptemplateinstantiateAPIRequest) GetClients() []string {
	return r._clients
}

// SetDescription is Description Setter
// 描述长度(9~200)
func (r *TaobaominiappptemplateinstantiateAPIRequest) SetDescription(_description string) error {
	r._description = _description
	r.Set("description", _description)
	return nil
}

// GetDescription Description Getter
func (r TaobaominiappptemplateinstantiateAPIRequest) GetDescription() string {
	return r._description
}

// SetExtJson is ExtJson Setter
// schemadata, json字符串
func (r *TaobaominiappptemplateinstantiateAPIRequest) SetExtJson(_extJson string) error {
	r._extJson = _extJson
	r.Set("ext_json", _extJson)
	return nil
}

// GetExtJson ExtJson Getter
func (r TaobaominiappptemplateinstantiateAPIRequest) GetExtJson() string {
	return r._extJson
}

// SetIcon is Icon Setter
// 小程序icon
func (r *TaobaominiappptemplateinstantiateAPIRequest) SetIcon(_icon string) error {
	r._icon = _icon
	r.Set("icon", _icon)
	return nil
}

// GetIcon Icon Getter
func (r TaobaominiappptemplateinstantiateAPIRequest) GetIcon() string {
	return r._icon
}

// SetName is Name Setter
// 小程序名称
func (r *TaobaominiappptemplateinstantiateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaominiappptemplateinstantiateAPIRequest) GetName() string {
	return r._name
}

// SetTemplateId is TemplateId Setter
// 模板id
func (r *TaobaominiappptemplateinstantiateAPIRequest) SetTemplateId(_templateId string) error {
	r._templateId = _templateId
	r.Set("template_id", _templateId)
	return nil
}

// GetTemplateId TemplateId Getter
func (r TaobaominiappptemplateinstantiateAPIRequest) GetTemplateId() string {
	return r._templateId
}

// SetTemplateVersion is TemplateVersion Setter
// 模板版本
func (r *TaobaominiappptemplateinstantiateAPIRequest) SetTemplateVersion(_templateVersion string) error {
	r._templateVersion = _templateVersion
	r.Set("template_version", _templateVersion)
	return nil
}

// GetTemplateVersion TemplateVersion Getter
func (r TaobaominiappptemplateinstantiateAPIRequest) GetTemplateVersion() string {
	return r._templateVersion
}
