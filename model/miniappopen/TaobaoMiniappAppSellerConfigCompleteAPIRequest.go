package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappappsellerconfigcompleteAPIRequest 商家完成小程序相关配置 API请求
// taobao.miniapp.app.seller.config.complete
//
// 通过该接口告知平台商家已经完成小程序相关的必要设置，可进行后续操作。主要用于小部件、客服插件等场景。
type TaobaominiappappsellerconfigcompleteAPIRequest struct {
	model.Params
	// 小部件必传，B端插件不用传。与app_id对应的已完成配置的版本号
	_version string
	// 商家已完成配置的小部件/B端插件的appid
	_appId int64
}

// NewTaobaominiappappsellerconfigcompleteRequest 初始化TaobaominiappappsellerconfigcompleteAPIRequest对象
func NewTaobaominiappappsellerconfigcompleteRequest() *TaobaominiappappsellerconfigcompleteAPIRequest {
	return &TaobaominiappappsellerconfigcompleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappappsellerconfigcompleteAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.app.seller.config.complete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappappsellerconfigcompleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappappsellerconfigcompleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVersion is Version Setter
// 小部件必传，B端插件不用传。与app_id对应的已完成配置的版本号
func (r *TaobaominiappappsellerconfigcompleteAPIRequest) SetVersion(_version string) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r TaobaominiappappsellerconfigcompleteAPIRequest) GetVersion() string {
	return r._version
}

// SetAppId is AppId Setter
// 商家已完成配置的小部件/B端插件的appid
func (r *TaobaominiappappsellerconfigcompleteAPIRequest) SetAppId(_appId int64) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r TaobaominiappappsellerconfigcompleteAPIRequest) GetAppId() int64 {
	return r._appId
}
