package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallMallitemcenterServiceproductQueryAPIRequest 天猫服务商服务产品查询 API请求
// tmall.mallitemcenter.serviceproduct.query
//
// 查询天猫服务的服务商发布的服务产品
type TmallMallitemcenterServiceproductQueryAPIRequest struct {
	model.Params
	// 服务名称
	_serviceCode string
	// 服务产品id
	_id int64
	// 产品状态
	_status int64
	// 产品类型
	_serviceProductType int64
}

// NewTmallMallitemcenterServiceproductQueryRequest 初始化TmallMallitemcenterServiceproductQueryAPIRequest对象
func NewTmallMallitemcenterServiceproductQueryRequest() *TmallMallitemcenterServiceproductQueryAPIRequest {
	return &TmallMallitemcenterServiceproductQueryAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallMallitemcenterServiceproductQueryAPIRequest) Reset() {
	r._serviceCode = ""
	r._id = 0
	r._status = 0
	r._serviceProductType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallMallitemcenterServiceproductQueryAPIRequest) GetApiMethodName() string {
	return "tmall.mallitemcenter.serviceproduct.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallMallitemcenterServiceproductQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallMallitemcenterServiceproductQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceCode is ServiceCode Setter
// 服务名称
func (r *TmallMallitemcenterServiceproductQueryAPIRequest) SetServiceCode(_serviceCode string) error {
	r._serviceCode = _serviceCode
	r.Set("service_code", _serviceCode)
	return nil
}

// GetServiceCode ServiceCode Getter
func (r TmallMallitemcenterServiceproductQueryAPIRequest) GetServiceCode() string {
	return r._serviceCode
}

// SetId is Id Setter
// 服务产品id
func (r *TmallMallitemcenterServiceproductQueryAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TmallMallitemcenterServiceproductQueryAPIRequest) GetId() int64 {
	return r._id
}

// SetStatus is Status Setter
// 产品状态
func (r *TmallMallitemcenterServiceproductQueryAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TmallMallitemcenterServiceproductQueryAPIRequest) GetStatus() int64 {
	return r._status
}

// SetServiceProductType is ServiceProductType Setter
// 产品类型
func (r *TmallMallitemcenterServiceproductQueryAPIRequest) SetServiceProductType(_serviceProductType int64) error {
	r._serviceProductType = _serviceProductType
	r.Set("service_product_type", _serviceProductType)
	return nil
}

// GetServiceProductType ServiceProductType Getter
func (r TmallMallitemcenterServiceproductQueryAPIRequest) GetServiceProductType() int64 {
	return r._serviceProductType
}

var poolTmallMallitemcenterServiceproductQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallMallitemcenterServiceproductQueryRequest()
	},
}

// GetTmallMallitemcenterServiceproductQueryRequest 从 sync.Pool 获取 TmallMallitemcenterServiceproductQueryAPIRequest
func GetTmallMallitemcenterServiceproductQueryAPIRequest() *TmallMallitemcenterServiceproductQueryAPIRequest {
	return poolTmallMallitemcenterServiceproductQueryAPIRequest.Get().(*TmallMallitemcenterServiceproductQueryAPIRequest)
}

// ReleaseTmallMallitemcenterServiceproductQueryAPIRequest 将 TmallMallitemcenterServiceproductQueryAPIRequest 放入 sync.Pool
func ReleaseTmallMallitemcenterServiceproductQueryAPIRequest(v *TmallMallitemcenterServiceproductQueryAPIRequest) {
	v.Reset()
	poolTmallMallitemcenterServiceproductQueryAPIRequest.Put(v)
}
