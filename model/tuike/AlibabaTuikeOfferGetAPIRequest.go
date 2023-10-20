package tuike

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatuikeoffergetAPIRequest 推广商品查询接口 API请求
// alibaba.tuike.offer.get
//
// 查询1688推客平台卖家推广中的商品信息
type AlibabatuikeoffergetAPIRequest struct {
	model.Params
	// 标识调用方
	_isvCode string
	// 搜索查询参数(json)
	_queryString string
}

// NewAlibabatuikeoffergetRequest 初始化AlibabatuikeoffergetAPIRequest对象
func NewAlibabatuikeoffergetRequest() *AlibabatuikeoffergetAPIRequest {
	return &AlibabatuikeoffergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatuikeoffergetAPIRequest) GetApiMethodName() string {
	return "alibaba.tuike.offer.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatuikeoffergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatuikeoffergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIsvCode is IsvCode Setter
// 标识调用方
func (r *AlibabatuikeoffergetAPIRequest) SetIsvCode(_isvCode string) error {
	r._isvCode = _isvCode
	r.Set("isv_code", _isvCode)
	return nil
}

// GetIsvCode IsvCode Getter
func (r AlibabatuikeoffergetAPIRequest) GetIsvCode() string {
	return r._isvCode
}

// SetQueryString is QueryString Setter
// 搜索查询参数(json)
func (r *AlibabatuikeoffergetAPIRequest) SetQueryString(_queryString string) error {
	r._queryString = _queryString
	r.Set("query_string", _queryString)
	return nil
}

// GetQueryString QueryString Getter
func (r AlibabatuikeoffergetAPIRequest) GetQueryString() string {
	return r._queryString
}
