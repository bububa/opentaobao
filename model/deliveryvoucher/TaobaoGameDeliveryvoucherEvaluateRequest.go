package deliveryvoucher

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卡券评价回传 API请求
taobao.game.deliveryvoucher.evaluate

卡券ISV回传商品评价
*/
type TaobaoGameDeliveryvoucherEvaluateRequest struct {
    model.Params
    // 系统自动生成
    param0   *VoucherEvaluateRequest
}

// 初始化TaobaoGameDeliveryvoucherEvaluateRequest对象
func NewTaobaoGameDeliveryvoucherEvaluateRequest() *TaobaoGameDeliveryvoucherEvaluateRequest{
    return &TaobaoGameDeliveryvoucherEvaluateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoGameDeliveryvoucherEvaluateRequest) GetApiMethodName() string {
    return "taobao.game.deliveryvoucher.evaluate"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoGameDeliveryvoucherEvaluateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 系统自动生成
func (r *TaobaoGameDeliveryvoucherEvaluateRequest) SetParam0(param0 *VoucherEvaluateRequest) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r TaobaoGameDeliveryvoucherEvaluateRequest) GetParam0() *VoucherEvaluateRequest {
    return r.param0
}
