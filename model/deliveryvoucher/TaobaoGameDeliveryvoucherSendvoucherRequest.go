package deliveryvoucher

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提货券发券接口 API请求
taobao.game.deliveryvoucher.sendvoucher

提货券发券接口：同步券和订单的关联信息
*/
type TaobaoGameDeliveryvoucherSendvoucherRequest struct {
    model.Params
    // 发券参数
    _param0   *SendVoucherRequest
}

// 初始化TaobaoGameDeliveryvoucherSendvoucherRequest对象
func NewTaobaoGameDeliveryvoucherSendvoucherRequest() *TaobaoGameDeliveryvoucherSendvoucherRequest{
    return &TaobaoGameDeliveryvoucherSendvoucherRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoGameDeliveryvoucherSendvoucherRequest) GetApiMethodName() string {
    return "taobao.game.deliveryvoucher.sendvoucher"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoGameDeliveryvoucherSendvoucherRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 发券参数
func (r *TaobaoGameDeliveryvoucherSendvoucherRequest) SetParam0(_param0 *SendVoucherRequest) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TaobaoGameDeliveryvoucherSendvoucherRequest) GetParam0() *SendVoucherRequest {
    return r._param0
}
