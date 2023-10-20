package icbuproduct

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuProductTypeAvailableGetAPIRequest) Reset() {
	r._typeRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductTypeAvailableGetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.type.available.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductTypeAvailableGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuProductTypeAvailableGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaIcbuProductTypeAvailableGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuProductTypeAvailableGetRequest()
	},
}

// GetAlibabaIcbuProductTypeAvailableGetRequest 从 sync.Pool 获取 AlibabaIcbuProductTypeAvailableGetAPIRequest
func GetAlibabaIcbuProductTypeAvailableGetAPIRequest() *AlibabaIcbuProductTypeAvailableGetAPIRequest {
	return poolAlibabaIcbuProductTypeAvailableGetAPIRequest.Get().(*AlibabaIcbuProductTypeAvailableGetAPIRequest)
}

// ReleaseAlibabaIcbuProductTypeAvailableGetAPIRequest 将 AlibabaIcbuProductTypeAvailableGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuProductTypeAvailableGetAPIRequest(v *AlibabaIcbuProductTypeAvailableGetAPIRequest) {
	v.Reset()
	poolAlibabaIcbuProductTypeAvailableGetAPIRequest.Put(v)
}
