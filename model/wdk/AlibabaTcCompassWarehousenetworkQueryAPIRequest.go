package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTcCompassWarehousenetworkQueryAPIRequest 按仓维度来查询鸟潮网络 API请求
// alibaba.tc.compass.warehousenetwork.query
//
// 按仓维度来查询鸟潮网络
type AlibabaTcCompassWarehousenetworkQueryAPIRequest struct {
	model.Params
	// 仓商家编码
	_warehouseMerchantCode string
	// 仓编码
	_warehouseCode string
	// WAVE_ARRIVE-波次达。ONE_HOUR-小时达
	_serviceType string
}

// NewAlibabaTcCompassWarehousenetworkQueryRequest 初始化AlibabaTcCompassWarehousenetworkQueryAPIRequest对象
func NewAlibabaTcCompassWarehousenetworkQueryRequest() *AlibabaTcCompassWarehousenetworkQueryAPIRequest {
	return &AlibabaTcCompassWarehousenetworkQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTcCompassWarehousenetworkQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.tc.compass.warehousenetwork.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTcCompassWarehousenetworkQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTcCompassWarehousenetworkQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseMerchantCode is WarehouseMerchantCode Setter
// 仓商家编码
func (r *AlibabaTcCompassWarehousenetworkQueryAPIRequest) SetWarehouseMerchantCode(_warehouseMerchantCode string) error {
	r._warehouseMerchantCode = _warehouseMerchantCode
	r.Set("warehouse_merchant_code", _warehouseMerchantCode)
	return nil
}

// GetWarehouseMerchantCode WarehouseMerchantCode Getter
func (r AlibabaTcCompassWarehousenetworkQueryAPIRequest) GetWarehouseMerchantCode() string {
	return r._warehouseMerchantCode
}

// SetWarehouseCode is WarehouseCode Setter
// 仓编码
func (r *AlibabaTcCompassWarehousenetworkQueryAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r AlibabaTcCompassWarehousenetworkQueryAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}

// SetServiceType is ServiceType Setter
// WAVE_ARRIVE-波次达。ONE_HOUR-小时达
func (r *AlibabaTcCompassWarehousenetworkQueryAPIRequest) SetServiceType(_serviceType string) error {
	r._serviceType = _serviceType
	r.Set("service_type", _serviceType)
	return nil
}

// GetServiceType ServiceType Getter
func (r AlibabaTcCompassWarehousenetworkQueryAPIRequest) GetServiceType() string {
	return r._serviceType
}
