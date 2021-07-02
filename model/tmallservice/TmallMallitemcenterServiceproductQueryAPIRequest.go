package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallMallitemcenterServiceproductQueryAPIRequest 天猫服务商服务产品查询 API请求
// tmall.mallitemcenter.serviceproduct.query
//
// 查询天猫服务的服务商发布的服务产品
type TmallMallitemcenterServiceproductQueryAPIRequest struct {
	model.Params
	// 服务产品id
	_id int64
	// 产品状态
	_status int64
	// 服务名称
	_serviceCode string
	// 产品类型
	_serviceProductType int64
}

// NewTmallMallitemcenterServiceproductQueryRequest 初始化TmallMallitemcenterServiceproductQueryAPIRequest对象
func NewTmallMallitemcenterServiceproductQueryRequest() *TmallMallitemcenterServiceproductQueryAPIRequest {
	return &TmallMallitemcenterServiceproductQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallMallitemcenterServiceproductQueryAPIRequest) GetApiMethodName() string {
	return "tmall.mallitemcenter.serviceproduct.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallMallitemcenterServiceproductQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Id Setter
// 服务产品id
func (r *TmallMallitemcenterServiceproductQueryAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// Get Id Getter
func (r TmallMallitemcenterServiceproductQueryAPIRequest) GetId() int64 {
	return r._id
}

// Set is Status Setter
// 产品状态
func (r *TmallMallitemcenterServiceproductQueryAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r TmallMallitemcenterServiceproductQueryAPIRequest) GetStatus() int64 {
	return r._status
}

// Set is ServiceCode Setter
// 服务名称
func (r *TmallMallitemcenterServiceproductQueryAPIRequest) SetServiceCode(_serviceCode string) error {
	r._serviceCode = _serviceCode
	r.Set("service_code", _serviceCode)
	return nil
}

// Get ServiceCode Getter
func (r TmallMallitemcenterServiceproductQueryAPIRequest) GetServiceCode() string {
	return r._serviceCode
}

// Set is ServiceProductType Setter
// 产品类型
func (r *TmallMallitemcenterServiceproductQueryAPIRequest) SetServiceProductType(_serviceProductType int64) error {
	r._serviceProductType = _serviceProductType
	r.Set("service_product_type", _serviceProductType)
	return nil
}

// Get ServiceProductType Getter
func (r TmallMallitemcenterServiceproductQueryAPIRequest) GetServiceProductType() int64 {
	return r._serviceProductType
}
