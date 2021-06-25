package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销退单事件回传接口 APIRequest
taobao.rdc.aligenius.warehouse.reverse.event.update

用于erp回传销退单相关信息到平台
*/
type TaobaoRdcAligeniusWarehouseReverseEventUpdateRequest struct {
    model.Params

    // 参数
    param0   *ReverseEventInfoDto 

}

func NewTaobaoRdcAligeniusWarehouseReverseEventUpdateRequest() *TaobaoRdcAligeniusWarehouseReverseEventUpdateRequest{
    return &TaobaoRdcAligeniusWarehouseReverseEventUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRdcAligeniusWarehouseReverseEventUpdateRequest) GetApiMethodName() string {
    return "taobao.rdc.aligenius.warehouse.reverse.event.update"
}

func (r TaobaoRdcAligeniusWarehouseReverseEventUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRdcAligeniusWarehouseReverseEventUpdateRequest) SetParam0(param0 *ReverseEventInfoDto) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoRdcAligeniusWarehouseReverseEventUpdateRequest) GetParam0() *ReverseEventInfoDto {
    return r.param0
}

