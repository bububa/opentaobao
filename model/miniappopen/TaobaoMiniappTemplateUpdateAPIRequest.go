package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiapptemplateupdateAPIRequest （已废弃）更新实例化应用 API请求
// taobao.miniapp.template.update
//
// 商家应用c端模板实例化小程序更新
type TaobaominiapptemplateupdateAPIRequest struct {
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

// NewTaobaominiapptemplateupdateRequest 初始化TaobaominiapptemplateupdateAPIRequest对象
func NewTaobaominiapptemplateupdateRequest() *TaobaominiapptemplateupdateAPIRequest {
	return &TaobaominiapptemplateupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiapptemplateupdateAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.template.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiapptemplateupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiapptemplateupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClients is Clients Setter
// 要更新的投放端,目前可投放： taobao(淘宝),tmall(天猫)
func (r *TaobaominiapptemplateupdateAPIRequest) SetClients(_clients []string) error {
	r._clients = _clients
	r.Set("clients", _clients)
	return nil
}

// GetClients Clients Getter
func (r TaobaominiapptemplateupdateAPIRequest) GetClients() []string {
	return r._clients
}

// SetId is Id Setter
// 应用id
func (r *TaobaominiapptemplateupdateAPIRequest) SetId(_id string) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaominiapptemplateupdateAPIRequest) GetId() string {
	return r._id
}

// SetExtJson is ExtJson Setter
// schema信息，不填且 应用线上版本使用的templateId与传入的templateId不一致，则会报错; 一致，则复用线上版本的schema。
func (r *TaobaominiapptemplateupdateAPIRequest) SetExtJson(_extJson string) error {
	r._extJson = _extJson
	r.Set("ext_json", _extJson)
	return nil
}

// GetExtJson ExtJson Getter
func (r TaobaominiapptemplateupdateAPIRequest) GetExtJson() string {
	return r._extJson
}

// SetTemplateId is TemplateId Setter
// 模板id
func (r *TaobaominiapptemplateupdateAPIRequest) SetTemplateId(_templateId string) error {
	r._templateId = _templateId
	r.Set("template_id", _templateId)
	return nil
}

// GetTemplateId TemplateId Getter
func (r TaobaominiapptemplateupdateAPIRequest) GetTemplateId() string {
	return r._templateId
}

// SetTemplateVersion is TemplateVersion Setter
// 模板版本
func (r *TaobaominiapptemplateupdateAPIRequest) SetTemplateVersion(_templateVersion string) error {
	r._templateVersion = _templateVersion
	r.Set("template_version", _templateVersion)
	return nil
}

// GetTemplateVersion TemplateVersion Getter
func (r TaobaominiapptemplateupdateAPIRequest) GetTemplateVersion() string {
	return r._templateVersion
}
