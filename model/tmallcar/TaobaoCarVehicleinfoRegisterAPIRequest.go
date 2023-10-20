package tmallcar

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCarVehicleinfoRegisterAPIRequest 全量车型导入 API请求
// taobao.car.vehicleinfo.register
//
// 全量车型导入
type TaobaoCarVehicleinfoRegisterAPIRequest struct {
	model.Params
	// 参数集合
	_paramList []FullInfoCarModelDto
}

// NewTaobaoCarVehicleinfoRegisterRequest 初始化TaobaoCarVehicleinfoRegisterAPIRequest对象
func NewTaobaoCarVehicleinfoRegisterRequest() *TaobaoCarVehicleinfoRegisterAPIRequest {
	return &TaobaoCarVehicleinfoRegisterAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCarVehicleinfoRegisterAPIRequest) Reset() {
	r._paramList = r._paramList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCarVehicleinfoRegisterAPIRequest) GetApiMethodName() string {
	return "taobao.car.vehicleinfo.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCarVehicleinfoRegisterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCarVehicleinfoRegisterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamList is ParamList Setter
// 参数集合
func (r *TaobaoCarVehicleinfoRegisterAPIRequest) SetParamList(_paramList []FullInfoCarModelDto) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r TaobaoCarVehicleinfoRegisterAPIRequest) GetParamList() []FullInfoCarModelDto {
	return r._paramList
}

var poolTaobaoCarVehicleinfoRegisterAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCarVehicleinfoRegisterRequest()
	},
}

// GetTaobaoCarVehicleinfoRegisterRequest 从 sync.Pool 获取 TaobaoCarVehicleinfoRegisterAPIRequest
func GetTaobaoCarVehicleinfoRegisterAPIRequest() *TaobaoCarVehicleinfoRegisterAPIRequest {
	return poolTaobaoCarVehicleinfoRegisterAPIRequest.Get().(*TaobaoCarVehicleinfoRegisterAPIRequest)
}

// ReleaseTaobaoCarVehicleinfoRegisterAPIRequest 将 TaobaoCarVehicleinfoRegisterAPIRequest 放入 sync.Pool
func ReleaseTaobaoCarVehicleinfoRegisterAPIRequest(v *TaobaoCarVehicleinfoRegisterAPIRequest) {
	v.Reset()
	poolTaobaoCarVehicleinfoRegisterAPIRequest.Put(v)
}
