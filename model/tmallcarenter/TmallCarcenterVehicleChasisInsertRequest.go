package tmallcarenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
EPC车型底盘压缩库新增接口 API请求
tmall.carcenter.vehicle.chasis.insert

EPC车型底盘压缩库新增接口
*/
type TmallCarcenterVehicleChasisInsertRequest struct {
    model.Params
    // 底盘压缩库入参
    dto   *ChasisVehicleInfoOriginalDto
}

// 初始化TmallCarcenterVehicleChasisInsertRequest对象
func NewTmallCarcenterVehicleChasisInsertRequest() *TmallCarcenterVehicleChasisInsertRequest{
    return &TmallCarcenterVehicleChasisInsertRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarcenterVehicleChasisInsertRequest) GetApiMethodName() string {
    return "tmall.carcenter.vehicle.chasis.insert"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarcenterVehicleChasisInsertRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Dto Setter
// 底盘压缩库入参
func (r *TmallCarcenterVehicleChasisInsertRequest) SetDto(dto *ChasisVehicleInfoOriginalDto) error {
    r.dto = dto
    r.Set("dto", dto)
    return nil
}

// Dto Getter
func (r TmallCarcenterVehicleChasisInsertRequest) GetDto() *ChasisVehicleInfoOriginalDto {
    return r.dto
}
