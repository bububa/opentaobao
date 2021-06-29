package deliveryvoucher

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
作废券 API请求
taobao.game.deliveryvoucher.cancelvoucher

提货券发券接口：同步券和订单的关联信息
*/
type TaobaoGameDeliveryvoucherCancelvoucherRequest struct {
    model.Params
    // 发券参数
    _param0   *CancelVoucherRequest
}

// 初始化TaobaoGameDeliveryvoucherCancelvoucherRequest对象
func NewTaobaoGameDeliveryvoucherCancelvoucherRequest() *TaobaoGameDeliveryvoucherCancelvoucherRequest{
    return &TaobaoGameDeliveryvoucherCancelvoucherRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoGameDeliveryvoucherCancelvoucherRequest) GetApiMethodName() string {
    return "taobao.game.deliveryvoucher.cancelvoucher"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoGameDeliveryvoucherCancelvoucherRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 发券参数
func (r *TaobaoGameDeliveryvoucherCancelvoucherRequest) SetParam0(_param0 *CancelVoucherRequest) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TaobaoGameDeliveryvoucherCancelvoucherRequest) GetParam0() *CancelVoucherRequest {
    return r._param0
}
