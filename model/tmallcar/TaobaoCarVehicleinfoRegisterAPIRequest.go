package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCarVehicleinfoRegisterAPIRequest
全量车型导入 API请求
taobao.car.vehicleinfo.register

全量车型导入 */
type TaobaoCarVehicleinfoRegisterAPIRequest struct {
	model.Params
	// 参数集合
	_paramList []FullInfoCarModelDto
}

// NewTaobaoCarVehicleinfoRegisterRequest 初始化TaobaoCarVehicleinfoRegisterAPIRequest对象
func NewTaobaoCarVehicleinfoRegisterRequest() *TaobaoCarVehicleinfoRegisterAPIRequest {
	return &TaobaoCarVehicleinfoRegisterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCarVehicleinfoRegisterAPIRequest) GetApiMethodName() string {
	return "taobao.car.vehicleinfo.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCarVehicleinfoRegisterAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamList Setter
// 参数集合
func (r *TaobaoCarVehicleinfoRegisterAPIRequest) SetParamList(_paramList []FullInfoCarModelDto) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// Get ParamList Getter
func (r TaobaoCarVehicleinfoRegisterAPIRequest) GetParamList() []FullInfoCarModelDto {
	return r._paramList
}
