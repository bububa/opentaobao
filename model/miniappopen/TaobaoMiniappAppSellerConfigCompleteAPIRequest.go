package miniappopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappAppSellerConfigCompleteAPIRequest 商家完成小程序相关配置 API请求
// taobao.miniapp.app.seller.config.complete
//
// 通过该接口告知平台商家已经完成小程序相关的必要设置，可进行后续操作。主要用于小部件、客服插件等场景。
type TaobaoMiniappAppSellerConfigCompleteAPIRequest struct {
	model.Params
	// 小部件必传，B端插件不用传。与app_id对应的已完成配置的版本号
	_version string
	// 商家已完成配置的小部件/B端插件的appid
	_appId int64
}

// NewTaobaoMiniappAppSellerConfigCompleteRequest 初始化TaobaoMiniappAppSellerConfigCompleteAPIRequest对象
func NewTaobaoMiniappAppSellerConfigCompleteRequest() *TaobaoMiniappAppSellerConfigCompleteAPIRequest {
	return &TaobaoMiniappAppSellerConfigCompleteAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMiniappAppSellerConfigCompleteAPIRequest) Reset() {
	r._version = ""
	r._appId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappAppSellerConfigCompleteAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.app.seller.config.complete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappAppSellerConfigCompleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappAppSellerConfigCompleteAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoMiniappAppSellerConfigCompleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMiniappAppSellerConfigCompleteRequest()
	},
}

// GetTaobaoMiniappAppSellerConfigCompleteRequest 从 sync.Pool 获取 TaobaoMiniappAppSellerConfigCompleteAPIRequest
func GetTaobaoMiniappAppSellerConfigCompleteAPIRequest() *TaobaoMiniappAppSellerConfigCompleteAPIRequest {
	return poolTaobaoMiniappAppSellerConfigCompleteAPIRequest.Get().(*TaobaoMiniappAppSellerConfigCompleteAPIRequest)
}

// ReleaseTaobaoMiniappAppSellerConfigCompleteAPIRequest 将 TaobaoMiniappAppSellerConfigCompleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoMiniappAppSellerConfigCompleteAPIRequest(v *TaobaoMiniappAppSellerConfigCompleteAPIRequest) {
	v.Reset()
	poolTaobaoMiniappAppSellerConfigCompleteAPIRequest.Put(v)
}
