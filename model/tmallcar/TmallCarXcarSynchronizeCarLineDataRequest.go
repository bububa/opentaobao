package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
我的爱卡车型配置数据 APIRequest
tmall.car.xcar.synchronize.car.line.data

同步我的爱卡车系配置数据到天猫汽车
*/
type TmallCarXcarSynchronizeCarLineDataRequest struct {
    model.Params

    // 入参对象
    paramXCarSysLineDTO   *XCarSysLineDto 

}

func NewTmallCarXcarSynchronizeCarLineDataRequest() *TmallCarXcarSynchronizeCarLineDataRequest{
    return &TmallCarXcarSynchronizeCarLineDataRequest{
        Params: model.NewParams(),
    }
}

func (r TmallCarXcarSynchronizeCarLineDataRequest) GetApiMethodName() string {
    return "tmall.car.xcar.synchronize.car.line.data"
}

func (r TmallCarXcarSynchronizeCarLineDataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallCarXcarSynchronizeCarLineDataRequest) SetParamXCarSysLineDTO(paramXCarSysLineDTO *XCarSysLineDto) error {
    r.paramXCarSysLineDTO = paramXCarSysLineDTO
    r.Set("param_x_car_sys_line_d_t_o", paramXCarSysLineDTO)
    return nil
}

func (r TmallCarXcarSynchronizeCarLineDataRequest) GetParamXCarSysLineDTO() *XCarSysLineDto {
    return r.paramXCarSysLineDTO
}

