package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkshopqueryAPIRequest 门店查询接口 API请求
// alibaba.wdk.shop.query
//
// 根据门店code查询门店信息
type AlibabawdkshopqueryAPIRequest struct {
	model.Params
	// 如果不传，返回所有
	_ouCode string
}

// NewAlibabawdkshopqueryRequest 初始化AlibabawdkshopqueryAPIRequest对象
func NewAlibabawdkshopqueryRequest() *AlibabawdkshopqueryAPIRequest {
	return &AlibabawdkshopqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkshopqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.shop.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkshopqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkshopqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuCode is OuCode Setter
// 如果不传，返回所有
func (r *AlibabawdkshopqueryAPIRequest) SetOuCode(_ouCode string) error {
	r._ouCode = _ouCode
	r.Set("ou_code", _ouCode)
	return nil
}

// GetOuCode OuCode Getter
func (r AlibabawdkshopqueryAPIRequest) GetOuCode() string {
	return r._ouCode
}
