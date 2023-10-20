package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeApartmentLineAPIRequest 公寓上下架 API请求
// alibaba.alihouse.newhome.apartment.line
//
// 公寓上下架
type AlibabaAlihouseNewhomeApartmentLineAPIRequest struct {
	model.Params
	// 外部公寓id
	_outerId string
	// 0-下架 1-上架
	_type *model.File
}

// NewAlibabaAlihouseNewhomeApartmentLineRequest 初始化AlibabaAlihouseNewhomeApartmentLineAPIRequest对象
func NewAlibabaAlihouseNewhomeApartmentLineRequest() *AlibabaAlihouseNewhomeApartmentLineAPIRequest {
	return &AlibabaAlihouseNewhomeApartmentLineAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeApartmentLineAPIRequest) Reset() {
	r._outerId = ""
	r._type = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeApartmentLineAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.apartment.line"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeApartmentLineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeApartmentLineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 外部公寓id
func (r *AlibabaAlihouseNewhomeApartmentLineAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaAlihouseNewhomeApartmentLineAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetType is Type Setter
// 0-下架 1-上架
func (r *AlibabaAlihouseNewhomeApartmentLineAPIRequest) SetType(_type *model.File) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaAlihouseNewhomeApartmentLineAPIRequest) GetType() *model.File {
	return r._type
}

var poolAlibabaAlihouseNewhomeApartmentLineAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeApartmentLineRequest()
	},
}

// GetAlibabaAlihouseNewhomeApartmentLineRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeApartmentLineAPIRequest
func GetAlibabaAlihouseNewhomeApartmentLineAPIRequest() *AlibabaAlihouseNewhomeApartmentLineAPIRequest {
	return poolAlibabaAlihouseNewhomeApartmentLineAPIRequest.Get().(*AlibabaAlihouseNewhomeApartmentLineAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeApartmentLineAPIRequest 将 AlibabaAlihouseNewhomeApartmentLineAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeApartmentLineAPIRequest(v *AlibabaAlihouseNewhomeApartmentLineAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeApartmentLineAPIRequest.Put(v)
}
