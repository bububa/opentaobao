package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推送及更新订单确认号 API请求
taobao.xhotel.order.update.confirmcode

商家拿到订单确认号后，异步推送给飞猪或更新给飞猪。订单确认号用于到店查无时的紧急查单依据。
*/
type TaobaoXhotelOrderUpdateConfirmcodeRequest struct {
    model.Params
    // 系统入参
    _param   *UpdateOrderConfirmCodeParam
}

// 初始化TaobaoXhotelOrderUpdateConfirmcodeRequest对象
func NewTaobaoXhotelOrderUpdateConfirmcodeRequest() *TaobaoXhotelOrderUpdateConfirmcodeRequest{
    return &TaobaoXhotelOrderUpdateConfirmcodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderUpdateConfirmcodeRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.update.confirmcode"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderUpdateConfirmcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 系统入参
func (r *TaobaoXhotelOrderUpdateConfirmcodeRequest) SetParam(_param *UpdateOrderConfirmCodeParam) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r TaobaoXhotelOrderUpdateConfirmcodeRequest) GetParam() *UpdateOrderConfirmCodeParam {
    return r._param
}
