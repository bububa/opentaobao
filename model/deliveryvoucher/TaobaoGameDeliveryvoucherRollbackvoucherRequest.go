package deliveryvoucher

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
回滚券 APIRequest
taobao.game.deliveryvoucher.rollbackvoucher

提货券发券接口：同步券和订单的关联信息
*/
type TaobaoGameDeliveryvoucherRollbackvoucherRequest struct {
    model.Params

    // 发券参数
    param0   *RollbackVoucherRequest 

}

func NewTaobaoGameDeliveryvoucherRollbackvoucherRequest() *TaobaoGameDeliveryvoucherRollbackvoucherRequest{
    return &TaobaoGameDeliveryvoucherRollbackvoucherRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoGameDeliveryvoucherRollbackvoucherRequest) GetApiMethodName() string {
    return "taobao.game.deliveryvoucher.rollbackvoucher"
}

func (r TaobaoGameDeliveryvoucherRollbackvoucherRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoGameDeliveryvoucherRollbackvoucherRequest) SetParam0(param0 *RollbackVoucherRequest) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoGameDeliveryvoucherRollbackvoucherRequest) GetParam0() *RollbackVoucherRequest {
    return r.param0
}

