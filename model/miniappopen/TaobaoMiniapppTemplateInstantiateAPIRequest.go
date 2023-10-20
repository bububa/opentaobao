package miniappopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniapppTemplateInstantiateAPIRequest （已废弃）构建实例化应用 API请求
// taobao.miniappp.template.instantiate
//
// 实例化saas化的小程序
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
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMiniapppTemplateInstantiateAPIRequest) Reset() {
	r._clients = r._clients[:0]
	r._description = ""
	r._extJson = ""
	r._icon = ""
	r._name = ""
	r._templateId = ""
	r._templateVersion = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniapppTemplateInstantiateAPIRequest) GetApiMethodName() string {
	return "taobao.miniappp.template.instantiate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniapppTemplateInstantiateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniapppTemplateInstantiateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClients is Clients Setter
// 投放端,目前可投放： taobao(淘宝),tmall(天猫)
func (r *TaobaoMiniapppTemplateInstantiateAPIRequest) SetClients(_clients []string) error {
	r._clients = _clients
	r.Set("clients", _clients)
	return nil
}

// GetClients Clients Getter
func (r TaobaoMiniapppTemplateInstantiateAPIRequest) GetClients() []string {
	return r._clients
}

// SetDescription is Description Setter
// 描述长度(9~200)
func (r *TaobaoMiniapppTemplateInstantiateAPIRequest) SetDescription(_description string) error {
	r._description = _description
	r.Set("description", _description)
	return nil
}

// GetDescription Description Getter
func (r TaobaoMiniapppTemplateInstantiateAPIRequest) GetDescription() string {
	return r._description
}

// SetExtJson is ExtJson Setter
// schemadata, json字符串
func (r *TaobaoMiniapppTemplateInstantiateAPIRequest) SetExtJson(_extJson string) error {
	r._extJson = _extJson
	r.Set("ext_json", _extJson)
	return nil
}

// GetExtJson ExtJson Getter
func (r TaobaoMiniapppTemplateInstantiateAPIRequest) GetExtJson() string {
	return r._extJson
}

// SetIcon is Icon Setter
// 小程序icon
func (r *TaobaoMiniapppTemplateInstantiateAPIRequest) SetIcon(_icon string) error {
	r._icon = _icon
	r.Set("icon", _icon)
	return nil
}

// GetIcon Icon Getter
func (r TaobaoMiniapppTemplateInstantiateAPIRequest) GetIcon() string {
	return r._icon
}

// SetName is Name Setter
// 小程序名称
func (r *TaobaoMiniapppTemplateInstantiateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoMiniapppTemplateInstantiateAPIRequest) GetName() string {
	return r._name
}

// SetTemplateId is TemplateId Setter
// 模板id
func (r *TaobaoMiniapppTemplateInstantiateAPIRequest) SetTemplateId(_templateId string) error {
	r._templateId = _templateId
	r.Set("template_id", _templateId)
	return nil
}

// GetTemplateId TemplateId Getter
func (r TaobaoMiniapppTemplateInstantiateAPIRequest) GetTemplateId() string {
	return r._templateId
}

// SetTemplateVersion is TemplateVersion Setter
// 模板版本
func (r *TaobaoMiniapppTemplateInstantiateAPIRequest) SetTemplateVersion(_templateVersion string) error {
	r._templateVersion = _templateVersion
	r.Set("template_version", _templateVersion)
	return nil
}

// GetTemplateVersion TemplateVersion Getter
func (r TaobaoMiniapppTemplateInstantiateAPIRequest) GetTemplateVersion() string {
	return r._templateVersion
}

var poolTaobaoMiniapppTemplateInstantiateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMiniapppTemplateInstantiateRequest()
	},
}

// GetTaobaoMiniapppTemplateInstantiateRequest 从 sync.Pool 获取 TaobaoMiniapppTemplateInstantiateAPIRequest
func GetTaobaoMiniapppTemplateInstantiateAPIRequest() *TaobaoMiniapppTemplateInstantiateAPIRequest {
	return poolTaobaoMiniapppTemplateInstantiateAPIRequest.Get().(*TaobaoMiniapppTemplateInstantiateAPIRequest)
}

// ReleaseTaobaoMiniapppTemplateInstantiateAPIRequest 将 TaobaoMiniapppTemplateInstantiateAPIRequest 放入 sync.Pool
func ReleaseTaobaoMiniapppTemplateInstantiateAPIRequest(v *TaobaoMiniapppTemplateInstantiateAPIRequest) {
	v.Reset()
	poolTaobaoMiniapppTemplateInstantiateAPIRequest.Put(v)
}
