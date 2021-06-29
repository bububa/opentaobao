package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
爱卡车系图片数据接入 API请求
tmall.car.xcar.synchronize.car.line.pics.data

爱卡车系图片数据同步天猫汽车
*/
type TmallCarXcarSynchronizeCarLinePicsDataRequest struct {
    model.Params
    // 入参对象
    _paramXCarSysLinePicsDTO   *XCarSysLinePicsDTO
}

// 初始化TmallCarXcarSynchronizeCarLinePicsDataRequest对象
func NewTmallCarXcarSynchronizeCarLinePicsDataRequest() *TmallCarXcarSynchronizeCarLinePicsDataRequest{
    return &TmallCarXcarSynchronizeCarLinePicsDataRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarXcarSynchronizeCarLinePicsDataRequest) GetApiMethodName() string {
    return "tmall.car.xcar.synchronize.car.line.pics.data"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarXcarSynchronizeCarLinePicsDataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamXCarSysLinePicsDTO Setter
// 入参对象
func (r *TmallCarXcarSynchronizeCarLinePicsDataRequest) SetParamXCarSysLinePicsDTO(_paramXCarSysLinePicsDTO *XCarSysLinePicsDTO) error {
    r._paramXCarSysLinePicsDTO = _paramXCarSysLinePicsDTO
    r.Set("param_x_car_sys_line_pics_d_t_o", _paramXCarSysLinePicsDTO)
    return nil
}

// ParamXCarSysLinePicsDTO Getter
func (r TmallCarXcarSynchronizeCarLinePicsDataRequest) GetParamXCarSysLinePicsDTO() *XCarSysLinePicsDTO {
    return r._paramXCarSysLinePicsDTO
}
