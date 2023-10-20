package miniappopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappTemplateInstantiateAPIRequest 构建实例化应用 API请求
// taobao.miniapp.template.instantiate
//
// 实例化saas化的小程序
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

// NewTaobaoMiniappTemplateInstantiateRequest 初始化TaobaoMiniappTemplateInstantiateAPIRequest对象
func NewTaobaoMiniappTemplateInstantiateRequest() *TaobaoMiniappTemplateInstantiateAPIRequest {
	return &TaobaoMiniappTemplateInstantiateAPIRequest{
		Params: model.NewParams(8),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMiniappTemplateInstantiateAPIRequest) Reset() {
	r._clients = r._clients[:0]
	r._description = ""
	r._extJson = ""
	r._icon = ""
	r._name = ""
	r._templateId = ""
	r._templateVersion = ""
	r._alias = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappTemplateInstantiateAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.template.instantiate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappTemplateInstantiateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappTemplateInstantiateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClients is Clients Setter
// 投放端,目前可投放： taobao(淘宝),tmall(天猫)，taobao为必填，需要模板在这些端上已经发布上线
func (r *TaobaoMiniappTemplateInstantiateAPIRequest) SetClients(_clients []string) error {
	r._clients = _clients
	r.Set("clients", _clients)
	return nil
}

// GetClients Clients Getter
func (r TaobaoMiniappTemplateInstantiateAPIRequest) GetClients() []string {
	return r._clients
}

// SetDescription is Description Setter
// 描述长度(9~200)
func (r *TaobaoMiniappTemplateInstantiateAPIRequest) SetDescription(_description string) error {
	r._description = _description
	r.Set("description", _description)
	return nil
}

// GetDescription Description Getter
func (r TaobaoMiniappTemplateInstantiateAPIRequest) GetDescription() string {
	return r._description
}

// SetExtJson is ExtJson Setter
// 扩展信息，JSON格式。
func (r *TaobaoMiniappTemplateInstantiateAPIRequest) SetExtJson(_extJson string) error {
	r._extJson = _extJson
	r.Set("ext_json", _extJson)
	return nil
}

// GetExtJson ExtJson Getter
func (r TaobaoMiniappTemplateInstantiateAPIRequest) GetExtJson() string {
	return r._extJson
}

// SetIcon is Icon Setter
// 小程序icon
func (r *TaobaoMiniappTemplateInstantiateAPIRequest) SetIcon(_icon string) error {
	r._icon = _icon
	r.Set("icon", _icon)
	return nil
}

// GetIcon Icon Getter
func (r TaobaoMiniappTemplateInstantiateAPIRequest) GetIcon() string {
	return r._icon
}

// SetName is Name Setter
// 小程序名称按平台规则自动生成，该字段仅做兜底使用。
func (r *TaobaoMiniappTemplateInstantiateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoMiniappTemplateInstantiateAPIRequest) GetName() string {
	return r._name
}

// SetTemplateId is TemplateId Setter
// 模板id
func (r *TaobaoMiniappTemplateInstantiateAPIRequest) SetTemplateId(_templateId string) error {
	r._templateId = _templateId
	r.Set("template_id", _templateId)
	return nil
}

// GetTemplateId TemplateId Getter
func (r TaobaoMiniappTemplateInstantiateAPIRequest) GetTemplateId() string {
	return r._templateId
}

// SetTemplateVersion is TemplateVersion Setter
// 模板版本
func (r *TaobaoMiniappTemplateInstantiateAPIRequest) SetTemplateVersion(_templateVersion string) error {
	r._templateVersion = _templateVersion
	r.Set("template_version", _templateVersion)
	return nil
}

// GetTemplateVersion TemplateVersion Getter
func (r TaobaoMiniappTemplateInstantiateAPIRequest) GetTemplateVersion() string {
	return r._templateVersion
}

// SetAlias is Alias Setter
// 小程序简称【1-16】字符，可重名
func (r *TaobaoMiniappTemplateInstantiateAPIRequest) SetAlias(_alias string) error {
	r._alias = _alias
	r.Set("alias", _alias)
	return nil
}

// GetAlias Alias Getter
func (r TaobaoMiniappTemplateInstantiateAPIRequest) GetAlias() string {
	return r._alias
}

var poolTaobaoMiniappTemplateInstantiateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMiniappTemplateInstantiateRequest()
	},
}

// GetTaobaoMiniappTemplateInstantiateRequest 从 sync.Pool 获取 TaobaoMiniappTemplateInstantiateAPIRequest
func GetTaobaoMiniappTemplateInstantiateAPIRequest() *TaobaoMiniappTemplateInstantiateAPIRequest {
	return poolTaobaoMiniappTemplateInstantiateAPIRequest.Get().(*TaobaoMiniappTemplateInstantiateAPIRequest)
}

// ReleaseTaobaoMiniappTemplateInstantiateAPIRequest 将 TaobaoMiniappTemplateInstantiateAPIRequest 放入 sync.Pool
func ReleaseTaobaoMiniappTemplateInstantiateAPIRequest(v *TaobaoMiniappTemplateInstantiateAPIRequest) {
	v.Reset()
	poolTaobaoMiniappTemplateInstantiateAPIRequest.Put(v)
}
