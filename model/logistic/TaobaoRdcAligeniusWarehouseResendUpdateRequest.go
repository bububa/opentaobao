package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
补发单状态回传 API请求
taobao.rdc.aligenius.warehouse.resend.update

补发单状态回传接口
*/
type TaobaoRdcAligeniusWarehouseResendUpdateRequest struct {
    model.Params
    // 参数
    _param0   *UpdateResendStatusDto
}

// 初始化TaobaoRdcAligeniusWarehouseResendUpdateRequest对象
func NewTaobaoRdcAligeniusWarehouseResendUpdateRequest() *TaobaoRdcAligeniusWarehouseResendUpdateRequest{
    return &TaobaoRdcAligeniusWarehouseResendUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRdcAligeniusWarehouseResendUpdateRequest) GetApiMethodName() string {
    return "taobao.rdc.aligenius.warehouse.resend.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRdcAligeniusWarehouseResendUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 参数
func (r *TaobaoRdcAligeniusWarehouseResendUpdateRequest) SetParam0(_param0 *UpdateResendStatusDto) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TaobaoRdcAligeniusWarehouseResendUpdateRequest) GetParam0() *UpdateResendStatusDto {
    return r._param0
}
