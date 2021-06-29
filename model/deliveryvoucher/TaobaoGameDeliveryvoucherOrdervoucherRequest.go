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
type TaobaoGameDeliveryvoucherOrdervoucherRequest struct {
    model.Params
    // 发券参数
    param0   *OrderVoucherRequest
}

// 初始化TaobaoGameDeliveryvoucherOrdervoucherRequest对象
func NewTaobaoGameDeliveryvoucherOrdervoucherRequest() *TaobaoGameDeliveryvoucherOrdervoucherRequest{
    return &TaobaoGameDeliveryvoucherOrdervoucherRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoGameDeliveryvoucherOrdervoucherRequest) GetApiMethodName() string {
    return "taobao.game.deliveryvoucher.ordervoucher"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoGameDeliveryvoucherOrdervoucherRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 发券参数
func (r *TaobaoGameDeliveryvoucherOrdervoucherRequest) SetParam0(param0 *OrderVoucherRequest) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r TaobaoGameDeliveryvoucherOrdervoucherRequest) GetParam0() *OrderVoucherRequest {
    return r.param0
}
