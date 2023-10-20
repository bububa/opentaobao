package miniappopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappTemplateOnlineappAPIRequest 上线实例化应用 API请求
// taobao.miniapp.template.onlineapp
//
// 将指定的预览版本发布上线，预览版本号由构建实例化或更新实例化接口返回。
type TaobaoMiniappTemplateOnlineappAPIRequest struct {
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

// NewTaobaoMiniappTemplateOnlineappRequest 初始化TaobaoMiniappTemplateOnlineappAPIRequest对象
func NewTaobaoMiniappTemplateOnlineappRequest() *TaobaoMiniappTemplateOnlineappAPIRequest {
	return &TaobaoMiniappTemplateOnlineappAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMiniappTemplateOnlineappAPIRequest) Reset() {
	r._clients = r._clients[:0]
	r._appId = ""
	r._templateId = ""
	r._templateVersion = ""
	r._appVersion = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappTemplateOnlineappAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.template.onlineapp"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappTemplateOnlineappAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappTemplateOnlineappAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClients is Clients Setter
// 要更新的投放端,目前可投放： taobao(淘宝),tmall(天猫)
func (r *TaobaoMiniappTemplateOnlineappAPIRequest) SetClients(_clients []string) error {
	r._clients = _clients
	r.Set("clients", _clients)
	return nil
}

// GetClients Clients Getter
func (r TaobaoMiniappTemplateOnlineappAPIRequest) GetClients() []string {
	return r._clients
}

// SetAppId is AppId Setter
// 小程序app_id
func (r *TaobaoMiniappTemplateOnlineappAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r TaobaoMiniappTemplateOnlineappAPIRequest) GetAppId() string {
	return r._appId
}

// SetTemplateId is TemplateId Setter
// 模板id
func (r *TaobaoMiniappTemplateOnlineappAPIRequest) SetTemplateId(_templateId string) error {
	r._templateId = _templateId
	r.Set("template_id", _templateId)
	return nil
}

// GetTemplateId TemplateId Getter
func (r TaobaoMiniappTemplateOnlineappAPIRequest) GetTemplateId() string {
	return r._templateId
}

// SetTemplateVersion is TemplateVersion Setter
// 模板版本
func (r *TaobaoMiniappTemplateOnlineappAPIRequest) SetTemplateVersion(_templateVersion string) error {
	r._templateVersion = _templateVersion
	r.Set("template_version", _templateVersion)
	return nil
}

// GetTemplateVersion TemplateVersion Getter
func (r TaobaoMiniappTemplateOnlineappAPIRequest) GetTemplateVersion() string {
	return r._templateVersion
}

// SetAppVersion is AppVersion Setter
// 待上线的预览版本号
func (r *TaobaoMiniappTemplateOnlineappAPIRequest) SetAppVersion(_appVersion string) error {
	r._appVersion = _appVersion
	r.Set("app_version", _appVersion)
	return nil
}

// GetAppVersion AppVersion Getter
func (r TaobaoMiniappTemplateOnlineappAPIRequest) GetAppVersion() string {
	return r._appVersion
}

var poolTaobaoMiniappTemplateOnlineappAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMiniappTemplateOnlineappRequest()
	},
}

// GetTaobaoMiniappTemplateOnlineappRequest 从 sync.Pool 获取 TaobaoMiniappTemplateOnlineappAPIRequest
func GetTaobaoMiniappTemplateOnlineappAPIRequest() *TaobaoMiniappTemplateOnlineappAPIRequest {
	return poolTaobaoMiniappTemplateOnlineappAPIRequest.Get().(*TaobaoMiniappTemplateOnlineappAPIRequest)
}

// ReleaseTaobaoMiniappTemplateOnlineappAPIRequest 将 TaobaoMiniappTemplateOnlineappAPIRequest 放入 sync.Pool
func ReleaseTaobaoMiniappTemplateOnlineappAPIRequest(v *TaobaoMiniappTemplateOnlineappAPIRequest) {
	v.Reset()
	poolTaobaoMiniappTemplateOnlineappAPIRequest.Put(v)
}
