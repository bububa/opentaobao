package tmallcarenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
车型数据更新 APIRequest
tmall.carcenter.vehicleinfo.register

基本车型信息维护
*/
type TmallCarcenterVehicleinfoRegisterRequest struct {
    model.Params

    // 车型数据对象
    vehicleInfo   *OriginVehicleInfoDto 

}

func NewTmallCarcenterVehicleinfoRegisterRequest() *TmallCarcenterVehicleinfoRegisterRequest{
    return &TmallCarcenterVehicleinfoRegisterRequest{
        Params: model.NewParams(),
    }
}

func (r TmallCarcenterVehicleinfoRegisterRequest) GetApiMethodName() string {
    return "tmall.carcenter.vehicleinfo.register"
}

func (r TmallCarcenterVehicleinfoRegisterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallCarcenterVehicleinfoRegisterRequest) SetVehicleInfo(vehicleInfo *OriginVehicleInfoDto) error {
    r.vehicleInfo = vehicleInfo
    r.Set("vehicle_info", vehicleInfo)
    return nil
}

func (r TmallCarcenterVehicleinfoRegisterRequest) GetVehicleInfo() *OriginVehicleInfoDto {
    return r.vehicleInfo
}

