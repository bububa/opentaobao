package tmallcarenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcarcentervehicleinforegisterAPIRequest 车型数据更新 API请求
// tmall.carcenter.vehicleinfo.register
//
// 基本车型信息维护
type TmallcarcentervehicleinforegisterAPIRequest struct {
	model.Params
	// 车型数据对象
	_vehicleInfo *OriginVehicleInfoDto
}

// NewTmallcarcentervehicleinforegisterRequest 初始化TmallcarcentervehicleinforegisterAPIRequest对象
func NewTmallcarcentervehicleinforegisterRequest() *TmallcarcentervehicleinforegisterAPIRequest {
	return &TmallcarcentervehicleinforegisterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcarcentervehicleinforegisterAPIRequest) GetApiMethodName() string {
	return "tmall.carcenter.vehicleinfo.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcarcentervehicleinforegisterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcarcentervehicleinforegisterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVehicleInfo is VehicleInfo Setter
// 车型数据对象
func (r *TmallcarcentervehicleinforegisterAPIRequest) SetVehicleInfo(_vehicleInfo *OriginVehicleInfoDto) error {
	r._vehicleInfo = _vehicleInfo
	r.Set("vehicle_info", _vehicleInfo)
	return nil
}

// GetVehicleInfo VehicleInfo Getter
func (r TmallcarcentervehicleinforegisterAPIRequest) GetVehicleInfo() *OriginVehicleInfoDto {
	return r._vehicleInfo
}
