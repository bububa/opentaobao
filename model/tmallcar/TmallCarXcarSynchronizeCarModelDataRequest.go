package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
爱车车型数据同步 APIRequest
tmall.car.xcar.synchronize.car.model.data

爱车汽车车型数据同步到天猫
*/
type TmallCarXcarSynchronizeCarModelDataRequest struct {
    model.Params

    // 传入对象描述
    paramXCarSysModelDTO   *XCarSysModelDto 

}

func NewTmallCarXcarSynchronizeCarModelDataRequest() *TmallCarXcarSynchronizeCarModelDataRequest{
    return &TmallCarXcarSynchronizeCarModelDataRequest{
        Params: model.NewParams(),
    }
}

func (r TmallCarXcarSynchronizeCarModelDataRequest) GetApiMethodName() string {
    return "tmall.car.xcar.synchronize.car.model.data"
}

func (r TmallCarXcarSynchronizeCarModelDataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallCarXcarSynchronizeCarModelDataRequest) SetParamXCarSysModelDTO(paramXCarSysModelDTO *XCarSysModelDto) error {
    r.paramXCarSysModelDTO = paramXCarSysModelDTO
    r.Set("param_x_car_sys_model_d_t_o", paramXCarSysModelDTO)
    return nil
}

func (r TmallCarXcarSynchronizeCarModelDataRequest) GetParamXCarSysModelDTO() *XCarSysModelDto {
    return r.paramXCarSysModelDTO
}

