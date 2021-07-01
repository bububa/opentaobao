package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
补发单物流信息回传 API请求
taobao.rdc.aligenius.warehouse.resend.logistics.msg.post

补发单erp物流信息回传平台
*/
type TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest struct {
    model.Params
    // 参数
    _param0   *SendResendLogisticsMsgDto
}

// 初始化TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest对象
func NewTaobaoRdcAligeniusWarehouseResendLogisticsMsgPostRequest() *TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest{
    return &TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest) GetApiMethodName() string {
    return "taobao.rdc.aligenius.warehouse.resend.logistics.msg.post"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 参数
func (r *TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest) SetParam0(_param0 *SendResendLogisticsMsgDto) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest) GetParam0() *SendResendLogisticsMsgDto {
    return r._param0
}
