package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocarvehicleinforegisterAPIRequest 全量车型导入 API请求
// taobao.car.vehicleinfo.register
//
// 全量车型导入
type TaobaocarvehicleinforegisterAPIRequest struct {
	model.Params
	// 参数集合
	_paramList []FullInfoCarModelDto
}

// NewTaobaocarvehicleinforegisterRequest 初始化TaobaocarvehicleinforegisterAPIRequest对象
func NewTaobaocarvehicleinforegisterRequest() *TaobaocarvehicleinforegisterAPIRequest {
	return &TaobaocarvehicleinforegisterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocarvehicleinforegisterAPIRequest) GetApiMethodName() string {
	return "taobao.car.vehicleinfo.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocarvehicleinforegisterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocarvehicleinforegisterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamList is ParamList Setter
// 参数集合
func (r *TaobaocarvehicleinforegisterAPIRequest) SetParamList(_paramList []FullInfoCarModelDto) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r TaobaocarvehicleinforegisterAPIRequest) GetParamList() []FullInfoCarModelDto {
	return r._paramList
}
