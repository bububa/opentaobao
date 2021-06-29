package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
爱车车型数据同步 API请求
tmall.car.xcar.synchronize.car.model.data

爱车汽车车型数据同步到天猫
*/
type TmallCarXcarSynchronizeCarModelDataRequest struct {
    model.Params
    // 传入对象描述
    paramXCarSysModelDTO   *XCarSysModelDto
}

// 初始化TmallCarXcarSynchronizeCarModelDataRequest对象
func NewTmallCarXcarSynchronizeCarModelDataRequest() *TmallCarXcarSynchronizeCarModelDataRequest{
    return &TmallCarXcarSynchronizeCarModelDataRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarXcarSynchronizeCarModelDataRequest) GetApiMethodName() string {
    return "tmall.car.xcar.synchronize.car.model.data"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarXcarSynchronizeCarModelDataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamXCarSysModelDTO Setter
// 传入对象描述
func (r *TmallCarXcarSynchronizeCarModelDataRequest) SetParamXCarSysModelDTO(paramXCarSysModelDTO *XCarSysModelDto) error {
    r.paramXCarSysModelDTO = paramXCarSysModelDTO
    r.Set("param_x_car_sys_model_d_t_o", paramXCarSysModelDTO)
    return nil
}

// ParamXCarSysModelDTO Getter
func (r TmallCarXcarSynchronizeCarModelDataRequest) GetParamXCarSysModelDTO() *XCarSysModelDto {
    return r.paramXCarSysModelDTO
}
