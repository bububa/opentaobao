package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销退单事件回传接口 API请求
taobao.rdc.aligenius.warehouse.reverse.event.update

用于erp回传销退单相关信息到平台
*/
type TaobaoRdcAligeniusWarehouseReverseEventUpdateRequest struct {
    model.Params
    // 参数
    _param0   *ReverseEventInfoDto
}

// 初始化TaobaoRdcAligeniusWarehouseReverseEventUpdateRequest对象
func NewTaobaoRdcAligeniusWarehouseReverseEventUpdateRequest() *TaobaoRdcAligeniusWarehouseReverseEventUpdateRequest{
    return &TaobaoRdcAligeniusWarehouseReverseEventUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRdcAligeniusWarehouseReverseEventUpdateRequest) GetApiMethodName() string {
    return "taobao.rdc.aligenius.warehouse.reverse.event.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRdcAligeniusWarehouseReverseEventUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 参数
func (r *TaobaoRdcAligeniusWarehouseReverseEventUpdateRequest) SetParam0(_param0 *ReverseEventInfoDto) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TaobaoRdcAligeniusWarehouseReverseEventUpdateRequest) GetParam0() *ReverseEventInfoDto {
    return r._param0
}
