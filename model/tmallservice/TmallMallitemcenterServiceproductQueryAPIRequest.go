package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallmallitemcenterserviceproductqueryAPIRequest 天猫服务商服务产品查询 API请求
// tmall.mallitemcenter.serviceproduct.query
//
// 查询天猫服务的服务商发布的服务产品
type TmallmallitemcenterserviceproductqueryAPIRequest struct {
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

// NewTmallmallitemcenterserviceproductqueryRequest 初始化TmallmallitemcenterserviceproductqueryAPIRequest对象
func NewTmallmallitemcenterserviceproductqueryRequest() *TmallmallitemcenterserviceproductqueryAPIRequest {
	return &TmallmallitemcenterserviceproductqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallmallitemcenterserviceproductqueryAPIRequest) GetApiMethodName() string {
	return "tmall.mallitemcenter.serviceproduct.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallmallitemcenterserviceproductqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallmallitemcenterserviceproductqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceCode is ServiceCode Setter
// 服务名称
func (r *TmallmallitemcenterserviceproductqueryAPIRequest) SetServiceCode(_serviceCode string) error {
	r._serviceCode = _serviceCode
	r.Set("service_code", _serviceCode)
	return nil
}

// GetServiceCode ServiceCode Getter
func (r TmallmallitemcenterserviceproductqueryAPIRequest) GetServiceCode() string {
	return r._serviceCode
}

// SetId is Id Setter
// 服务产品id
func (r *TmallmallitemcenterserviceproductqueryAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TmallmallitemcenterserviceproductqueryAPIRequest) GetId() int64 {
	return r._id
}

// SetStatus is Status Setter
// 产品状态
func (r *TmallmallitemcenterserviceproductqueryAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TmallmallitemcenterserviceproductqueryAPIRequest) GetStatus() int64 {
	return r._status
}

// SetServiceProductType is ServiceProductType Setter
// 产品类型
func (r *TmallmallitemcenterserviceproductqueryAPIRequest) SetServiceProductType(_serviceProductType int64) error {
	r._serviceProductType = _serviceProductType
	r.Set("service_product_type", _serviceProductType)
	return nil
}

// GetServiceProductType ServiceProductType Getter
func (r TmallmallitemcenterserviceproductqueryAPIRequest) GetServiceProductType() int64 {
	return r._serviceProductType
}
