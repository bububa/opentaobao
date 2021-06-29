package tmallcarenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车EPC版本压缩库新增接口 APIRequest
tmall.carcenter.vehicle.version.insert

汽车EPC版本压缩库新增接口
*/
type TmallCarcenterVehicleVersionInsertRequest struct {
    model.Params

    // 版本压缩库入参
    dto   *VersionVehicleInfoOriginalDto 

}

func NewTmallCarcenterVehicleVersionInsertRequest() *TmallCarcenterVehicleVersionInsertRequest{
    return &TmallCarcenterVehicleVersionInsertRequest{
        Params: model.NewParams(),
    }
}

func (r TmallCarcenterVehicleVersionInsertRequest) GetApiMethodName() string {
    return "tmall.carcenter.vehicle.version.insert"
}

func (r TmallCarcenterVehicleVersionInsertRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallCarcenterVehicleVersionInsertRequest) SetDto(dto *VersionVehicleInfoOriginalDto) error {
    r.dto = dto
    r.Set("dto", dto)
    return nil
}

func (r TmallCarcenterVehicleVersionInsertRequest) GetDto() *VersionVehicleInfoOriginalDto {
    return r.dto
}

