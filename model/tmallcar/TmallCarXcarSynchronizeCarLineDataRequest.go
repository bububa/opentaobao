package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
我的爱卡车型配置数据 API请求
tmall.car.xcar.synchronize.car.line.data

同步我的爱卡车系配置数据到天猫汽车
*/
type TmallCarXcarSynchronizeCarLineDataRequest struct {
    model.Params
    // 入参对象
    _paramXCarSysLineDTO   *XCarSysLineDto
}

// 初始化TmallCarXcarSynchronizeCarLineDataRequest对象
func NewTmallCarXcarSynchronizeCarLineDataRequest() *TmallCarXcarSynchronizeCarLineDataRequest{
    return &TmallCarXcarSynchronizeCarLineDataRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarXcarSynchronizeCarLineDataRequest) GetApiMethodName() string {
    return "tmall.car.xcar.synchronize.car.line.data"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarXcarSynchronizeCarLineDataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamXCarSysLineDTO Setter
// 入参对象
func (r *TmallCarXcarSynchronizeCarLineDataRequest) SetParamXCarSysLineDTO(_paramXCarSysLineDTO *XCarSysLineDto) error {
    r._paramXCarSysLineDTO = _paramXCarSysLineDTO
    r.Set("param_x_car_sys_line_d_t_o", _paramXCarSysLineDTO)
    return nil
}

// ParamXCarSysLineDTO Getter
func (r TmallCarXcarSynchronizeCarLineDataRequest) GetParamXCarSysLineDTO() *XCarSysLineDto {
    return r._paramXCarSysLineDTO
}
