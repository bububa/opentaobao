package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
POS收银成功后订单同步 API请求
alibaba.mj.oc.pay

此API用于在银泰商场中，消费者在收银台收银/退款时， POS系统在收银或退款成功后，调用此接口进行订单同步
*/
type AlibabaMjOcPayAPIRequest struct {
    model.Params
    // 订单数据
    _posOrder   *PosOrderDto
}

// 初始化AlibabaMjOcPayAPIRequest对象
func NewAlibabaMjOcPayRequest() *AlibabaMjOcPayAPIRequest{
    return &AlibabaMjOcPayAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMjOcPayAPIRequest) GetApiMethodName() string {
    return "alibaba.mj.oc.pay"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMjOcPayAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PosOrder Setter
// 订单数据
func (r *AlibabaMjOcPayAPIRequest) SetPosOrder(_posOrder *PosOrderDto) error {
    r._posOrder = _posOrder
    r.Set("pos_order", _posOrder)
    return nil
}

// PosOrder Getter
func (r AlibabaMjOcPayAPIRequest) GetPosOrder() *PosOrderDto {
    return r._posOrder
}
