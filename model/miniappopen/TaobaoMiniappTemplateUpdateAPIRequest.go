package miniappopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappTemplateUpdateAPIRequest （已废弃）更新实例化应用 API请求
// taobao.miniapp.template.update
//
// 商家应用c端模板实例化小程序更新
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

// NewTaobaoMiniappTemplateUpdateRequest 初始化TaobaoMiniappTemplateUpdateAPIRequest对象
func NewTaobaoMiniappTemplateUpdateRequest() *TaobaoMiniappTemplateUpdateAPIRequest {
	return &TaobaoMiniappTemplateUpdateAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMiniappTemplateUpdateAPIRequest) Reset() {
	r._clients = r._clients[:0]
	r._id = ""
	r._extJson = ""
	r._templateId = ""
	r._templateVersion = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappTemplateUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.template.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappTemplateUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappTemplateUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClients is Clients Setter
// 要更新的投放端,目前可投放： taobao(淘宝),tmall(天猫)
func (r *TaobaoMiniappTemplateUpdateAPIRequest) SetClients(_clients []string) error {
	r._clients = _clients
	r.Set("clients", _clients)
	return nil
}

// GetClients Clients Getter
func (r TaobaoMiniappTemplateUpdateAPIRequest) GetClients() []string {
	return r._clients
}

// SetId is Id Setter
// 应用id
func (r *TaobaoMiniappTemplateUpdateAPIRequest) SetId(_id string) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaoMiniappTemplateUpdateAPIRequest) GetId() string {
	return r._id
}

// SetExtJson is ExtJson Setter
// schema信息，不填且 应用线上版本使用的templateId与传入的templateId不一致，则会报错; 一致，则复用线上版本的schema。
func (r *TaobaoMiniappTemplateUpdateAPIRequest) SetExtJson(_extJson string) error {
	r._extJson = _extJson
	r.Set("ext_json", _extJson)
	return nil
}

// GetExtJson ExtJson Getter
func (r TaobaoMiniappTemplateUpdateAPIRequest) GetExtJson() string {
	return r._extJson
}

// SetTemplateId is TemplateId Setter
// 模板id
func (r *TaobaoMiniappTemplateUpdateAPIRequest) SetTemplateId(_templateId string) error {
	r._templateId = _templateId
	r.Set("template_id", _templateId)
	return nil
}

// GetTemplateId TemplateId Getter
func (r TaobaoMiniappTemplateUpdateAPIRequest) GetTemplateId() string {
	return r._templateId
}

// SetTemplateVersion is TemplateVersion Setter
// 模板版本
func (r *TaobaoMiniappTemplateUpdateAPIRequest) SetTemplateVersion(_templateVersion string) error {
	r._templateVersion = _templateVersion
	r.Set("template_version", _templateVersion)
	return nil
}

// GetTemplateVersion TemplateVersion Getter
func (r TaobaoMiniappTemplateUpdateAPIRequest) GetTemplateVersion() string {
	return r._templateVersion
}

var poolTaobaoMiniappTemplateUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMiniappTemplateUpdateRequest()
	},
}

// GetTaobaoMiniappTemplateUpdateRequest 从 sync.Pool 获取 TaobaoMiniappTemplateUpdateAPIRequest
func GetTaobaoMiniappTemplateUpdateAPIRequest() *TaobaoMiniappTemplateUpdateAPIRequest {
	return poolTaobaoMiniappTemplateUpdateAPIRequest.Get().(*TaobaoMiniappTemplateUpdateAPIRequest)
}

// ReleaseTaobaoMiniappTemplateUpdateAPIRequest 将 TaobaoMiniappTemplateUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoMiniappTemplateUpdateAPIRequest(v *TaobaoMiniappTemplateUpdateAPIRequest) {
	v.Reset()
	poolTaobaoMiniappTemplateUpdateAPIRequest.Put(v)
}
