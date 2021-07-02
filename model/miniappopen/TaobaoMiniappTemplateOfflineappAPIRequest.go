package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappTemplateOfflineappAPIRequest 下线实例化应用 API请求
// taobao.miniapp.template.offlineapp
//
// 对指定的实例化小程序进行下线,需要指定clients和app_version
type TaobaoMiniappTemplateOfflineappAPIRequest struct {
	model.Params
	// 要下线的投放端,目前可投放： taobao(淘宝),tmall(天猫)
	_clients []string
	// 要下线的小程序app_id
	_appId string
	// 要下线的小程序版本号
	_appVersion string
	// 模板id
	_templateId string
}

// NewTaobaoMiniappTemplateOfflineappRequest 初始化TaobaoMiniappTemplateOfflineappAPIRequest对象
func NewTaobaoMiniappTemplateOfflineappRequest() *TaobaoMiniappTemplateOfflineappAPIRequest {
	return &TaobaoMiniappTemplateOfflineappAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappTemplateOfflineappAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.template.offlineapp"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappTemplateOfflineappAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetClients is Clients Setter
// 要下线的投放端,目前可投放： taobao(淘宝),tmall(天猫)
func (r *TaobaoMiniappTemplateOfflineappAPIRequest) SetClients(_clients []string) error {
	r._clients = _clients
	r.Set("clients", _clients)
	return nil
}

// GetClients Clients Getter
func (r TaobaoMiniappTemplateOfflineappAPIRequest) GetClients() []string {
	return r._clients
}

// SetAppId is AppId Setter
// 要下线的小程序app_id
func (r *TaobaoMiniappTemplateOfflineappAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r TaobaoMiniappTemplateOfflineappAPIRequest) GetAppId() string {
	return r._appId
}

// SetAppVersion is AppVersion Setter
// 要下线的小程序版本号
func (r *TaobaoMiniappTemplateOfflineappAPIRequest) SetAppVersion(_appVersion string) error {
	r._appVersion = _appVersion
	r.Set("app_version", _appVersion)
	return nil
}

// GetAppVersion AppVersion Getter
func (r TaobaoMiniappTemplateOfflineappAPIRequest) GetAppVersion() string {
	return r._appVersion
}

// SetTemplateId is TemplateId Setter
// 模板id
func (r *TaobaoMiniappTemplateOfflineappAPIRequest) SetTemplateId(_templateId string) error {
	r._templateId = _templateId
	r.Set("template_id", _templateId)
	return nil
}

// GetTemplateId TemplateId Getter
func (r TaobaoMiniappTemplateOfflineappAPIRequest) GetTemplateId() string {
	return r._templateId
}
