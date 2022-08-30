package icbuproduct

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductTypeAvailableGetAPIRequest 商家发品类型查询 API请求
// alibaba.icbu.product.type.available.get
//
// 查询商家发品权限
type AlibabaIcbuProductTypeAvailableGetAPIRequest struct {
	model.Params
	// -
	_typeRequest *ProductTopRequest
}

// NewAlibabaIcbuProductTypeAvailableGetRequest 初始化AlibabaIcbuProductTypeAvailableGetAPIRequest对象
func NewAlibabaIcbuProductTypeAvailableGetRequest() *AlibabaIcbuProductTypeAvailableGetAPIRequest {
	return &AlibabaIcbuProductTypeAvailableGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductTypeAvailableGetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.type.available.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductTypeAvailableGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTypeRequest is TypeRequest Setter
// -
func (r *AlibabaIcbuProductTypeAvailableGetAPIRequest) SetTypeRequest(_typeRequest *ProductTopRequest) error {
	r._typeRequest = _typeRequest
	r.Set("type_request", _typeRequest)
	return nil
}

// GetTypeRequest TypeRequest Getter
func (r AlibabaIcbuProductTypeAvailableGetAPIRequest) GetTypeRequest() *ProductTopRequest {
	return r._typeRequest
}
