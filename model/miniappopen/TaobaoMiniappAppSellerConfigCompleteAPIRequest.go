package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappAppSellerConfigCompleteAPIRequest 商家完成小程序相关配置 API请求
// taobao.miniapp.app.seller.config.complete
//
// 通过该接口告知平台商家已经完成小程序相关的必要设置，可进行后续操作。主要用于小部件、客服插件等场景。
type TaobaoMiniappAppSellerConfigCompleteAPIRequest struct {
	model.Params
	// 商家已完成配置的小部件/B端插件的appid
	_appId int64
	// 小部件必传，B端插件不用传。与app_id对应的已完成配置的版本号
	_version string
}

// NewTaobaoMiniappAppSellerConfigCompleteRequest 初始化TaobaoMiniappAppSellerConfigCompleteAPIRequest对象
func NewTaobaoMiniappAppSellerConfigCompleteRequest() *TaobaoMiniappAppSellerConfigCompleteAPIRequest {
	return &TaobaoMiniappAppSellerConfigCompleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappAppSellerConfigCompleteAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.app.seller.config.complete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappAppSellerConfigCompleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAppId is AppId Setter
// 商家已完成配置的小部件/B端插件的appid
func (r *TaobaoMiniappAppSellerConfigCompleteAPIRequest) SetAppId(_appId int64) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r TaobaoMiniappAppSellerConfigCompleteAPIRequest) GetAppId() int64 {
	return r._appId
}

// SetVersion is Version Setter
// 小部件必传，B端插件不用传。与app_id对应的已完成配置的版本号
func (r *TaobaoMiniappAppSellerConfigCompleteAPIRequest) SetVersion(_version string) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r TaobaoMiniappAppSellerConfigCompleteAPIRequest) GetVersion() string {
	return r._version
}
