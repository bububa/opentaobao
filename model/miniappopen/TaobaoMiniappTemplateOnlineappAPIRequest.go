package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiapptemplateonlineappAPIRequest 上线实例化应用 API请求
// taobao.miniapp.template.onlineapp
//
// 将指定的预览版本发布上线，预览版本号由构建实例化或更新实例化接口返回。
type TaobaominiapptemplateonlineappAPIRequest struct {
	model.Params
	// 要更新的投放端,目前可投放： taobao(淘宝),tmall(天猫)
	_clients []string
	// 小程序app_id
	_appId string
	// 模板id
	_templateId string
	// 模板版本
	_templateVersion string
	// 待上线的预览版本号
	_appVersion string
}

// NewTaobaominiapptemplateonlineappRequest 初始化TaobaominiapptemplateonlineappAPIRequest对象
func NewTaobaominiapptemplateonlineappRequest() *TaobaominiapptemplateonlineappAPIRequest {
	return &TaobaominiapptemplateonlineappAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiapptemplateonlineappAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.template.onlineapp"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiapptemplateonlineappAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiapptemplateonlineappAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClients is Clients Setter
// 要更新的投放端,目前可投放： taobao(淘宝),tmall(天猫)
func (r *TaobaominiapptemplateonlineappAPIRequest) SetClients(_clients []string) error {
	r._clients = _clients
	r.Set("clients", _clients)
	return nil
}

// GetClients Clients Getter
func (r TaobaominiapptemplateonlineappAPIRequest) GetClients() []string {
	return r._clients
}

// SetAppId is AppId Setter
// 小程序app_id
func (r *TaobaominiapptemplateonlineappAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r TaobaominiapptemplateonlineappAPIRequest) GetAppId() string {
	return r._appId
}

// SetTemplateId is TemplateId Setter
// 模板id
func (r *TaobaominiapptemplateonlineappAPIRequest) SetTemplateId(_templateId string) error {
	r._templateId = _templateId
	r.Set("template_id", _templateId)
	return nil
}

// GetTemplateId TemplateId Getter
func (r TaobaominiapptemplateonlineappAPIRequest) GetTemplateId() string {
	return r._templateId
}

// SetTemplateVersion is TemplateVersion Setter
// 模板版本
func (r *TaobaominiapptemplateonlineappAPIRequest) SetTemplateVersion(_templateVersion string) error {
	r._templateVersion = _templateVersion
	r.Set("template_version", _templateVersion)
	return nil
}

// GetTemplateVersion TemplateVersion Getter
func (r TaobaominiapptemplateonlineappAPIRequest) GetTemplateVersion() string {
	return r._templateVersion
}

// SetAppVersion is AppVersion Setter
// 待上线的预览版本号
func (r *TaobaominiapptemplateonlineappAPIRequest) SetAppVersion(_appVersion string) error {
	r._appVersion = _appVersion
	r.Set("app_version", _appVersion)
	return nil
}

// GetAppVersion AppVersion Getter
func (r TaobaominiapptemplateonlineappAPIRequest) GetAppVersion() string {
	return r._appVersion
}
