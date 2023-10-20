package icbuproduct

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuproducttypeavailablegetAPIRequest 商家发品类型查询 API请求
// alibaba.icbu.product.type.available.get
//
// 查询商家发品权限
type AlibabaicbuproducttypeavailablegetAPIRequest struct {
	model.Params
	// -
	_typeRequest *ProductTopRequest
}

// NewAlibabaicbuproducttypeavailablegetRequest 初始化AlibabaicbuproducttypeavailablegetAPIRequest对象
func NewAlibabaicbuproducttypeavailablegetRequest() *AlibabaicbuproducttypeavailablegetAPIRequest {
	return &AlibabaicbuproducttypeavailablegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbuproducttypeavailablegetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.type.available.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbuproducttypeavailablegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbuproducttypeavailablegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTypeRequest is TypeRequest Setter
// -
func (r *AlibabaicbuproducttypeavailablegetAPIRequest) SetTypeRequest(_typeRequest *ProductTopRequest) error {
	r._typeRequest = _typeRequest
	r.Set("type_request", _typeRequest)
	return nil
}

// GetTypeRequest TypeRequest Getter
func (r AlibabaicbuproducttypeavailablegetAPIRequest) GetTypeRequest() *ProductTopRequest {
	return r._typeRequest
}
