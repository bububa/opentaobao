package tuike

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTuikeOfferGetProAPIRequest 推广商品查询接口2.0 API请求
// alibaba.tuike.offer.get.pro
//
// 查询1688推客平台卖家推广中的商品信息
type AlibabaTuikeOfferGetProAPIRequest struct {
	model.Params
	// 用户ID
	_loginId string
	// 标识调用方
	_isvCode string
	// 搜索查询参数(json)
	_queryString string
}

// NewAlibabaTuikeOfferGetProRequest 初始化AlibabaTuikeOfferGetProAPIRequest对象
func NewAlibabaTuikeOfferGetProRequest() *AlibabaTuikeOfferGetProAPIRequest {
	return &AlibabaTuikeOfferGetProAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTuikeOfferGetProAPIRequest) GetApiMethodName() string {
	return "alibaba.tuike.offer.get.pro"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTuikeOfferGetProAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetLoginId is LoginId Setter
// 用户ID
func (r *AlibabaTuikeOfferGetProAPIRequest) SetLoginId(_loginId string) error {
	r._loginId = _loginId
	r.Set("login_id", _loginId)
	return nil
}

// GetLoginId LoginId Getter
func (r AlibabaTuikeOfferGetProAPIRequest) GetLoginId() string {
	return r._loginId
}

// SetIsvCode is IsvCode Setter
// 标识调用方
func (r *AlibabaTuikeOfferGetProAPIRequest) SetIsvCode(_isvCode string) error {
	r._isvCode = _isvCode
	r.Set("isv_code", _isvCode)
	return nil
}

// GetIsvCode IsvCode Getter
func (r AlibabaTuikeOfferGetProAPIRequest) GetIsvCode() string {
	return r._isvCode
}

// SetQueryString is QueryString Setter
// 搜索查询参数(json)
func (r *AlibabaTuikeOfferGetProAPIRequest) SetQueryString(_queryString string) error {
	r._queryString = _queryString
	r.Set("query_string", _queryString)
	return nil
}

// GetQueryString QueryString Getter
func (r AlibabaTuikeOfferGetProAPIRequest) GetQueryString() string {
	return r._queryString
}
