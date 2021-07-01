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
type TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest struct {
    model.Params
    // 参数
    _param0   *UpdateResendStatusDTO
}

// 初始化TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest对象
func NewTaobaoRdcAligeniusWarehouseResendUpdateRequest() *TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest{
    return &TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.rdc.aligenius.warehouse.resend.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 参数
func (r *TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest) SetParam0(_param0 *UpdateResendStatusDTO) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest) GetParam0() *UpdateResendStatusDTO {
    return r._param0
}
