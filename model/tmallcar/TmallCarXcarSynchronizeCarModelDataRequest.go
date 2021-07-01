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
type TmallCarXcarSynchronizeCarModelDataAPIRequest struct {
    model.Params
    // 传入对象描述
    _paramXCarSysModelDTO   *XCarSysModelDTO
}

// 初始化TmallCarXcarSynchronizeCarModelDataAPIRequest对象
func NewTmallCarXcarSynchronizeCarModelDataRequest() *TmallCarXcarSynchronizeCarModelDataAPIRequest{
    return &TmallCarXcarSynchronizeCarModelDataAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarXcarSynchronizeCarModelDataAPIRequest) GetApiMethodName() string {
    return "tmall.car.xcar.synchronize.car.model.data"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarXcarSynchronizeCarModelDataAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamXCarSysModelDTO Setter
// 传入对象描述
func (r *TmallCarXcarSynchronizeCarModelDataAPIRequest) SetParamXCarSysModelDTO(_paramXCarSysModelDTO *XCarSysModelDTO) error {
    r._paramXCarSysModelDTO = _paramXCarSysModelDTO
    r.Set("param_x_car_sys_model_d_t_o", _paramXCarSysModelDTO)
    return nil
}

// ParamXCarSysModelDTO Getter
func (r TmallCarXcarSynchronizeCarModelDataAPIRequest) GetParamXCarSysModelDTO() *XCarSysModelDTO {
    return r._paramXCarSysModelDTO
}
