package tmallcarenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
EPC车型底盘压缩库新增接口 APIRequest
tmall.carcenter.vehicle.chasis.insert

EPC车型底盘压缩库新增接口
*/
type TmallCarcenterVehicleChasisInsertRequest struct {
    model.Params

    // 底盘压缩库入参
    dto   *ChasisVehicleInfoOriginalDto 

}

func NewTmallCarcenterVehicleChasisInsertRequest() *TmallCarcenterVehicleChasisInsertRequest{
    return &TmallCarcenterVehicleChasisInsertRequest{
        Params: model.NewParams(),
    }
}

func (r TmallCarcenterVehicleChasisInsertRequest) GetApiMethodName() string {
    return "tmall.carcenter.vehicle.chasis.insert"
}

func (r TmallCarcenterVehicleChasisInsertRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallCarcenterVehicleChasisInsertRequest) SetDto(dto *ChasisVehicleInfoOriginalDto) error {
    r.dto = dto
    r.Set("dto", dto)
    return nil
}

func (r TmallCarcenterVehicleChasisInsertRequest) GetDto() *ChasisVehicleInfoOriginalDto {
    return r.dto
}

