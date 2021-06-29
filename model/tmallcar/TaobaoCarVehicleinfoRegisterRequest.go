package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全量车型导入 API请求
taobao.car.vehicleinfo.register

全量车型导入
*/
type TaobaoCarVehicleinfoRegisterRequest struct {
    model.Params
    // 参数集合
    _paramList   []FullInfoCarModelDTO
}

// 初始化TaobaoCarVehicleinfoRegisterRequest对象
func NewTaobaoCarVehicleinfoRegisterRequest() *TaobaoCarVehicleinfoRegisterRequest{
    return &TaobaoCarVehicleinfoRegisterRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCarVehicleinfoRegisterRequest) GetApiMethodName() string {
    return "taobao.car.vehicleinfo.register"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCarVehicleinfoRegisterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamList Setter
// 参数集合
func (r *TaobaoCarVehicleinfoRegisterRequest) SetParamList(_paramList []FullInfoCarModelDTO) error {
    r._paramList = _paramList
    r.Set("param_list", _paramList)
    return nil
}

// ParamList Getter
func (r TaobaoCarVehicleinfoRegisterRequest) GetParamList() []FullInfoCarModelDTO {
    return r._paramList
}
