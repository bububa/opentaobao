package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiapptemplateinstantiateAPIRequest 构建实例化应用 API请求
// taobao.miniapp.template.instantiate
//
// 实例化saas化的小程序
type TaobaominiapptemplateinstantiateAPIRequest struct {
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

// NewTaobaominiapptemplateinstantiateRequest 初始化TaobaominiapptemplateinstantiateAPIRequest对象
func NewTaobaominiapptemplateinstantiateRequest() *TaobaominiapptemplateinstantiateAPIRequest {
	return &TaobaominiapptemplateinstantiateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiapptemplateinstantiateAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.template.instantiate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiapptemplateinstantiateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiapptemplateinstantiateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClients is Clients Setter
// 投放端,目前可投放： taobao(淘宝),tmall(天猫)，taobao为必填，需要模板在这些端上已经发布上线
func (r *TaobaominiapptemplateinstantiateAPIRequest) SetClients(_clients []string) error {
	r._clients = _clients
	r.Set("clients", _clients)
	return nil
}

// GetClients Clients Getter
func (r TaobaominiapptemplateinstantiateAPIRequest) GetClients() []string {
	return r._clients
}

// SetDescription is Description Setter
// 描述长度(9~200)
func (r *TaobaominiapptemplateinstantiateAPIRequest) SetDescription(_description string) error {
	r._description = _description
	r.Set("description", _description)
	return nil
}

// GetDescription Description Getter
func (r TaobaominiapptemplateinstantiateAPIRequest) GetDescription() string {
	return r._description
}

// SetExtJson is ExtJson Setter
// 扩展信息，JSON格式。
func (r *TaobaominiapptemplateinstantiateAPIRequest) SetExtJson(_extJson string) error {
	r._extJson = _extJson
	r.Set("ext_json", _extJson)
	return nil
}

// GetExtJson ExtJson Getter
func (r TaobaominiapptemplateinstantiateAPIRequest) GetExtJson() string {
	return r._extJson
}

// SetIcon is Icon Setter
// 小程序icon
func (r *TaobaominiapptemplateinstantiateAPIRequest) SetIcon(_icon string) error {
	r._icon = _icon
	r.Set("icon", _icon)
	return nil
}

// GetIcon Icon Getter
func (r TaobaominiapptemplateinstantiateAPIRequest) GetIcon() string {
	return r._icon
}

// SetName is Name Setter
// 小程序名称按平台规则自动生成，该字段仅做兜底使用。
func (r *TaobaominiapptemplateinstantiateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaominiapptemplateinstantiateAPIRequest) GetName() string {
	return r._name
}

// SetTemplateId is TemplateId Setter
// 模板id
func (r *TaobaominiapptemplateinstantiateAPIRequest) SetTemplateId(_templateId string) error {
	r._templateId = _templateId
	r.Set("template_id", _templateId)
	return nil
}

// GetTemplateId TemplateId Getter
func (r TaobaominiapptemplateinstantiateAPIRequest) GetTemplateId() string {
	return r._templateId
}

// SetTemplateVersion is TemplateVersion Setter
// 模板版本
func (r *TaobaominiapptemplateinstantiateAPIRequest) SetTemplateVersion(_templateVersion string) error {
	r._templateVersion = _templateVersion
	r.Set("template_version", _templateVersion)
	return nil
}

// GetTemplateVersion TemplateVersion Getter
func (r TaobaominiapptemplateinstantiateAPIRequest) GetTemplateVersion() string {
	return r._templateVersion
}

// SetAlias is Alias Setter
// 小程序简称【1-16】字符，可重名
func (r *TaobaominiapptemplateinstantiateAPIRequest) SetAlias(_alias string) error {
	r._alias = _alias
	r.Set("alias", _alias)
	return nil
}

// GetAlias Alias Getter
func (r TaobaominiapptemplateinstantiateAPIRequest) GetAlias() string {
	return r._alias
}
