package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiapptemplaterollbackAPIRequest 回滚实例化应用 API请求
// taobao.miniapp.template.rollback
//
// 将实例化小程序回滚到指定版本
type TaobaominiapptemplaterollbackAPIRequest struct {
	model.Params
	// 要回滚的投放端,目前可投放： taobao,tmall
	_clients []string
	// 小程序app_id
	_appId string
	// 回到到该app_version
	_appVersion string
	// 实例化小程序对应的模板id
	_templateId string
	// 与app_version对应的模板版本
	_templateVersion string
}

// NewTaobaominiapptemplaterollbackRequest 初始化TaobaominiapptemplaterollbackAPIRequest对象
func NewTaobaominiapptemplaterollbackRequest() *TaobaominiapptemplaterollbackAPIRequest {
	return &TaobaominiapptemplaterollbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiapptemplaterollbackAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.template.rollback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiapptemplaterollbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiapptemplaterollbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClients is Clients Setter
// 要回滚的投放端,目前可投放： taobao,tmall
func (r *TaobaominiapptemplaterollbackAPIRequest) SetClients(_clients []string) error {
	r._clients = _clients
	r.Set("clients", _clients)
	return nil
}

// GetClients Clients Getter
func (r TaobaominiapptemplaterollbackAPIRequest) GetClients() []string {
	return r._clients
}

// SetAppId is AppId Setter
// 小程序app_id
func (r *TaobaominiapptemplaterollbackAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r TaobaominiapptemplaterollbackAPIRequest) GetAppId() string {
	return r._appId
}

// SetAppVersion is AppVersion Setter
// 回到到该app_version
func (r *TaobaominiapptemplaterollbackAPIRequest) SetAppVersion(_appVersion string) error {
	r._appVersion = _appVersion
	r.Set("app_version", _appVersion)
	return nil
}

// GetAppVersion AppVersion Getter
func (r TaobaominiapptemplaterollbackAPIRequest) GetAppVersion() string {
	return r._appVersion
}

// SetTemplateId is TemplateId Setter
// 实例化小程序对应的模板id
func (r *TaobaominiapptemplaterollbackAPIRequest) SetTemplateId(_templateId string) error {
	r._templateId = _templateId
	r.Set("template_id", _templateId)
	return nil
}

// GetTemplateId TemplateId Getter
func (r TaobaominiapptemplaterollbackAPIRequest) GetTemplateId() string {
	return r._templateId
}

// SetTemplateVersion is TemplateVersion Setter
// 与app_version对应的模板版本
func (r *TaobaominiapptemplaterollbackAPIRequest) SetTemplateVersion(_templateVersion string) error {
	r._templateVersion = _templateVersion
	r.Set("template_version", _templateVersion)
	return nil
}

// GetTemplateVersion TemplateVersion Getter
func (r TaobaominiapptemplaterollbackAPIRequest) GetTemplateVersion() string {
	return r._templateVersion
}
