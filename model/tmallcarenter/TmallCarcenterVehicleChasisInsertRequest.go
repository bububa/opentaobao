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
type TmallCarcenterVehicleChasisInsertAPIRequest struct {
    model.Params
    // 底盘压缩库入参
    _dto   *ChasisVehicleInfoOriginalDTO
}

// 初始化TmallCarcenterVehicleChasisInsertAPIRequest对象
func NewTmallCarcenterVehicleChasisInsertRequest() *TmallCarcenterVehicleChasisInsertAPIRequest{
    return &TmallCarcenterVehicleChasisInsertAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarcenterVehicleChasisInsertAPIRequest) GetApiMethodName() string {
    return "tmall.carcenter.vehicle.chasis.insert"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarcenterVehicleChasisInsertAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Dto Setter
// 底盘压缩库入参
func (r *TmallCarcenterVehicleChasisInsertAPIRequest) SetDto(_dto *ChasisVehicleInfoOriginalDTO) error {
    r._dto = _dto
    r.Set("dto", _dto)
    return nil
}

// Dto Getter
func (r TmallCarcenterVehicleChasisInsertAPIRequest) GetDto() *ChasisVehicleInfoOriginalDTO {
    return r._dto
}
