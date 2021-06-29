package deliveryvoucher

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
回滚券 API请求
taobao.game.deliveryvoucher.rollbackvoucher

提货券发券接口：同步券和订单的关联信息
*/
type TaobaoGameDeliveryvoucherRollbackvoucherRequest struct {
    model.Params
    // 发券参数
    _param0   *RollbackVoucherRequest
}

// 初始化TaobaoGameDeliveryvoucherRollbackvoucherRequest对象
func NewTaobaoGameDeliveryvoucherRollbackvoucherRequest() *TaobaoGameDeliveryvoucherRollbackvoucherRequest{
    return &TaobaoGameDeliveryvoucherRollbackvoucherRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoGameDeliveryvoucherRollbackvoucherRequest) GetApiMethodName() string {
    return "taobao.game.deliveryvoucher.rollbackvoucher"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoGameDeliveryvoucherRollbackvoucherRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 发券参数
func (r *TaobaoGameDeliveryvoucherRollbackvoucherRequest) SetParam0(_param0 *RollbackVoucherRequest) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TaobaoGameDeliveryvoucherRollbackvoucherRequest) GetParam0() *RollbackVoucherRequest {
    return r._param0
}
