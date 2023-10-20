package topoaid

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopPackageAuthInfoGetAPIRequest 淘宝末端包裹信息获取 API请求
// taobao.top.package.auth.info.get
//
// 淘宝末端包裹信息获取
type TaobaoTopPackageAuthInfoGetAPIRequest struct {
	model.Params
	// 查询类
	_isPrivacyPackageRequest *IsPrivacyPackageRequest
}

// NewTaobaoTopPackageAuthInfoGetRequest 初始化TaobaoTopPackageAuthInfoGetAPIRequest对象
func NewTaobaoTopPackageAuthInfoGetRequest() *TaobaoTopPackageAuthInfoGetAPIRequest {
	return &TaobaoTopPackageAuthInfoGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTopPackageAuthInfoGetAPIRequest) Reset() {
	r._isPrivacyPackageRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTopPackageAuthInfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.top.package.auth.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTopPackageAuthInfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTopPackageAuthInfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIsPrivacyPackageRequest is IsPrivacyPackageRequest Setter
// 查询类
func (r *TaobaoTopPackageAuthInfoGetAPIRequest) SetIsPrivacyPackageRequest(_isPrivacyPackageRequest *IsPrivacyPackageRequest) error {
	r._isPrivacyPackageRequest = _isPrivacyPackageRequest
	r.Set("is_privacy_package_request", _isPrivacyPackageRequest)
	return nil
}

// GetIsPrivacyPackageRequest IsPrivacyPackageRequest Getter
func (r TaobaoTopPackageAuthInfoGetAPIRequest) GetIsPrivacyPackageRequest() *IsPrivacyPackageRequest {
	return r._isPrivacyPackageRequest
}

var poolTaobaoTopPackageAuthInfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTopPackageAuthInfoGetRequest()
	},
}

// GetTaobaoTopPackageAuthInfoGetRequest 从 sync.Pool 获取 TaobaoTopPackageAuthInfoGetAPIRequest
func GetTaobaoTopPackageAuthInfoGetAPIRequest() *TaobaoTopPackageAuthInfoGetAPIRequest {
	return poolTaobaoTopPackageAuthInfoGetAPIRequest.Get().(*TaobaoTopPackageAuthInfoGetAPIRequest)
}

// ReleaseTaobaoTopPackageAuthInfoGetAPIRequest 将 TaobaoTopPackageAuthInfoGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTopPackageAuthInfoGetAPIRequest(v *TaobaoTopPackageAuthInfoGetAPIRequest) {
	v.Reset()
	poolTaobaoTopPackageAuthInfoGetAPIRequest.Put(v)
}
