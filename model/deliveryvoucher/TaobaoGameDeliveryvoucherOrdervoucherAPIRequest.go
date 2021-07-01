package deliveryvoucher

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
预约接口 API请求
taobao.game.deliveryvoucher.ordervoucher

提货券发券接口：同步券和订单的关联信息
*/
type TaobaoGameDeliveryvoucherOrdervoucherAPIRequest struct {
    model.Params
    // 发券参数
    _param0   *OrderVoucherRequest
}

// 初始化TaobaoGameDeliveryvoucherOrdervoucherAPIRequest对象
func NewTaobaoGameDeliveryvoucherOrdervoucherRequest() *TaobaoGameDeliveryvoucherOrdervoucherAPIRequest{
    return &TaobaoGameDeliveryvoucherOrdervoucherAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoGameDeliveryvoucherOrdervoucherAPIRequest) GetApiMethodName() string {
    return "taobao.game.deliveryvoucher.ordervoucher"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoGameDeliveryvoucherOrdervoucherAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 发券参数
func (r *TaobaoGameDeliveryvoucherOrdervoucherAPIRequest) SetParam0(_param0 *OrderVoucherRequest) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TaobaoGameDeliveryvoucherOrdervoucherAPIRequest) GetParam0() *OrderVoucherRequest {
    return r._param0
}
