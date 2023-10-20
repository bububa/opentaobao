package tmallcarenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallCarcenterVehicleinfoRegisterAPIRequest 车型数据更新 API请求
// tmall.carcenter.vehicleinfo.register
//
// 基本车型信息维护
type TmallCarcenterVehicleinfoRegisterAPIRequest struct {
	model.Params
	// 车型数据对象
	_vehicleInfo *OriginVehicleInfoDto
}

// NewTmallCarcenterVehicleinfoRegisterRequest 初始化TmallCarcenterVehicleinfoRegisterAPIRequest对象
func NewTmallCarcenterVehicleinfoRegisterRequest() *TmallCarcenterVehicleinfoRegisterAPIRequest {
	return &TmallCarcenterVehicleinfoRegisterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarcenterVehicleinfoRegisterAPIRequest) GetApiMethodName() string {
	return "tmall.carcenter.vehicleinfo.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarcenterVehicleinfoRegisterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCarcenterVehicleinfoRegisterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVehicleInfo is VehicleInfo Setter
// 车型数据对象
func (r *TmallCarcenterVehicleinfoRegisterAPIRequest) SetVehicleInfo(_vehicleInfo *OriginVehicleInfoDto) error {
	r._vehicleInfo = _vehicleInfo
	r.Set("vehicle_info", _vehicleInfo)
	return nil
}

// GetVehicleInfo VehicleInfo Getter
func (r TmallCarcenterVehicleinfoRegisterAPIRequest) GetVehicleInfo() *OriginVehicleInfoDto {
	return r._vehicleInfo
}
