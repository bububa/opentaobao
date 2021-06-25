package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
补发单状态回传 APIRequest
taobao.rdc.aligenius.warehouse.resend.update

补发单状态回传接口
*/
type TaobaoRdcAligeniusWarehouseResendUpdateRequest struct {
    model.Params

    // 参数
    param0   *UpdateResendStatusDto 

}

func NewTaobaoRdcAligeniusWarehouseResendUpdateRequest() *TaobaoRdcAligeniusWarehouseResendUpdateRequest{
    return &TaobaoRdcAligeniusWarehouseResendUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRdcAligeniusWarehouseResendUpdateRequest) GetApiMethodName() string {
    return "taobao.rdc.aligenius.warehouse.resend.update"
}

func (r TaobaoRdcAligeniusWarehouseResendUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRdcAligeniusWarehouseResendUpdateRequest) SetParam0(param0 *UpdateResendStatusDto) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoRdcAligeniusWarehouseResendUpdateRequest) GetParam0() *UpdateResendStatusDto {
    return r.param0
}

