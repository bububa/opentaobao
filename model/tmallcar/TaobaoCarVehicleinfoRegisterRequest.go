package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全量车型导入 APIRequest
taobao.car.vehicleinfo.register

全量车型导入
*/
type TaobaoCarVehicleinfoRegisterRequest struct {
    model.Params

    // 参数集合
    paramList   []FullInfoCarModelDTO 

}

func NewTaobaoCarVehicleinfoRegisterRequest() *TaobaoCarVehicleinfoRegisterRequest{
    return &TaobaoCarVehicleinfoRegisterRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCarVehicleinfoRegisterRequest) GetApiMethodName() string {
    return "taobao.car.vehicleinfo.register"
}

func (r TaobaoCarVehicleinfoRegisterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoCarVehicleinfoRegisterRequest) SetParamList(paramList []FullInfoCarModelDTO) error {
    r.paramList = paramList
    r.Set("param_list", paramList)
    return nil
}

func (r TaobaoCarVehicleinfoRegisterRequest) GetParamList() []FullInfoCarModelDTO {
    return r.paramList
}

