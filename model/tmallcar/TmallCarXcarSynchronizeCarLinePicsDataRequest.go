package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
爱卡车系图片数据接入 APIRequest
tmall.car.xcar.synchronize.car.line.pics.data

爱卡车系图片数据同步天猫汽车
*/
type TmallCarXcarSynchronizeCarLinePicsDataRequest struct {
    model.Params

    // 入参对象
    paramXCarSysLinePicsDTO   *XCarSysLinePicsDto 

}

func NewTmallCarXcarSynchronizeCarLinePicsDataRequest() *TmallCarXcarSynchronizeCarLinePicsDataRequest{
    return &TmallCarXcarSynchronizeCarLinePicsDataRequest{
        Params: model.NewParams(),
    }
}

func (r TmallCarXcarSynchronizeCarLinePicsDataRequest) GetApiMethodName() string {
    return "tmall.car.xcar.synchronize.car.line.pics.data"
}

func (r TmallCarXcarSynchronizeCarLinePicsDataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallCarXcarSynchronizeCarLinePicsDataRequest) SetParamXCarSysLinePicsDTO(paramXCarSysLinePicsDTO *XCarSysLinePicsDto) error {
    r.paramXCarSysLinePicsDTO = paramXCarSysLinePicsDTO
    r.Set("param_x_car_sys_line_pics_d_t_o", paramXCarSysLinePicsDTO)
    return nil
}

func (r TmallCarXcarSynchronizeCarLinePicsDataRequest) GetParamXCarSysLinePicsDTO() *XCarSysLinePicsDto {
    return r.paramXCarSysLinePicsDTO
}

