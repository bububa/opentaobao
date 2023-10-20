package topoaid

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotoppackageauthinfogetAPIRequest 淘宝末端包裹信息获取 API请求
// taobao.top.package.auth.info.get
//
// 淘宝末端包裹信息获取
type TaobaotoppackageauthinfogetAPIRequest struct {
	model.Params
	// 查询类
	_isPrivacyPackageRequest *IsPrivacyPackageRequest
}

// NewTaobaotoppackageauthinfogetRequest 初始化TaobaotoppackageauthinfogetAPIRequest对象
func NewTaobaotoppackageauthinfogetRequest() *TaobaotoppackageauthinfogetAPIRequest {
	return &TaobaotoppackageauthinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotoppackageauthinfogetAPIRequest) GetApiMethodName() string {
	return "taobao.top.package.auth.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotoppackageauthinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotoppackageauthinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIsPrivacyPackageRequest is IsPrivacyPackageRequest Setter
// 查询类
func (r *TaobaotoppackageauthinfogetAPIRequest) SetIsPrivacyPackageRequest(_isPrivacyPackageRequest *IsPrivacyPackageRequest) error {
	r._isPrivacyPackageRequest = _isPrivacyPackageRequest
	r.Set("is_privacy_package_request", _isPrivacyPackageRequest)
	return nil
}

// GetIsPrivacyPackageRequest IsPrivacyPackageRequest Getter
func (r TaobaotoppackageauthinfogetAPIRequest) GetIsPrivacyPackageRequest() *IsPrivacyPackageRequest {
	return r._isPrivacyPackageRequest
}
