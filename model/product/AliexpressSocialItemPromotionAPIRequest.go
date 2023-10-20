package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssocialitempromotionAPIRequest 获取推广链接 API请求
// aliexpress.social.item.promotion
//
// 获取商品社交推广链接
type AliexpresssocialitempromotionAPIRequest struct {
	model.Params
	// 推广的商品链接
	_targetUrl string
	// 子渠道号
	_af string
	// campaign Id
	_cn string
	// creative id
	_cv string
	// click id
	_dp string
}

// NewAliexpresssocialitempromotionRequest 初始化AliexpresssocialitempromotionAPIRequest对象
func NewAliexpresssocialitempromotionRequest() *AliexpresssocialitempromotionAPIRequest {
	return &AliexpresssocialitempromotionAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssocialitempromotionAPIRequest) GetApiMethodName() string {
	return "aliexpress.social.item.promotion"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssocialitempromotionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssocialitempromotionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTargetUrl is TargetUrl Setter
// 推广的商品链接
func (r *AliexpresssocialitempromotionAPIRequest) SetTargetUrl(_targetUrl string) error {
	r._targetUrl = _targetUrl
	r.Set("target_url", _targetUrl)
	return nil
}

// GetTargetUrl TargetUrl Getter
func (r AliexpresssocialitempromotionAPIRequest) GetTargetUrl() string {
	return r._targetUrl
}

// SetAf is Af Setter
// 子渠道号
func (r *AliexpresssocialitempromotionAPIRequest) SetAf(_af string) error {
	r._af = _af
	r.Set("af", _af)
	return nil
}

// GetAf Af Getter
func (r AliexpresssocialitempromotionAPIRequest) GetAf() string {
	return r._af
}

// SetCn is Cn Setter
// campaign Id
func (r *AliexpresssocialitempromotionAPIRequest) SetCn(_cn string) error {
	r._cn = _cn
	r.Set("cn", _cn)
	return nil
}

// GetCn Cn Getter
func (r AliexpresssocialitempromotionAPIRequest) GetCn() string {
	return r._cn
}

// SetCv is Cv Setter
// creative id
func (r *AliexpresssocialitempromotionAPIRequest) SetCv(_cv string) error {
	r._cv = _cv
	r.Set("cv", _cv)
	return nil
}

// GetCv Cv Getter
func (r AliexpresssocialitempromotionAPIRequest) GetCv() string {
	return r._cv
}

// SetDp is Dp Setter
// click id
func (r *AliexpresssocialitempromotionAPIRequest) SetDp(_dp string) error {
	r._dp = _dp
	r.Set("dp", _dp)
	return nil
}

// GetDp Dp Getter
func (r AliexpresssocialitempromotionAPIRequest) GetDp() string {
	return r._dp
}
