package tuike

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTuikeOfferGetAPIRequest
推广商品查询接口 API请求
alibaba.tuike.offer.get

查询1688推客平台卖家推广中的商品信息 */
type AlibabaTuikeOfferGetAPIRequest struct {
	model.Params
	// 标识调用方
	_isvCode string
	// 搜索查询参数(json)
	_queryString string
}

// NewAlibabaTuikeOfferGetRequest 初始化AlibabaTuikeOfferGetAPIRequest对象
func NewAlibabaTuikeOfferGetRequest() *AlibabaTuikeOfferGetAPIRequest {
	return &AlibabaTuikeOfferGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTuikeOfferGetAPIRequest) GetApiMethodName() string {
	return "alibaba.tuike.offer.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTuikeOfferGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is IsvCode Setter
// 标识调用方
func (r *AlibabaTuikeOfferGetAPIRequest) SetIsvCode(_isvCode string) error {
	r._isvCode = _isvCode
	r.Set("isv_code", _isvCode)
	return nil
}

// Get IsvCode Getter
func (r AlibabaTuikeOfferGetAPIRequest) GetIsvCode() string {
	return r._isvCode
}

// Set is QueryString Setter
// 搜索查询参数(json)
func (r *AlibabaTuikeOfferGetAPIRequest) SetQueryString(_queryString string) error {
	r._queryString = _queryString
	r.Set("query_string", _queryString)
	return nil
}

// Get QueryString Getter
func (r AlibabaTuikeOfferGetAPIRequest) GetQueryString() string {
	return r._queryString
}
